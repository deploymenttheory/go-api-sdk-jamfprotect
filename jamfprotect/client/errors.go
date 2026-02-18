package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

// Sentinel errors
var (
	ErrAuthentication  = errors.New("authentication failed")
	ErrGraphQL         = errors.New("graphql operation failed")
	ErrNotFound        = errors.New("resource not found")
	ErrInvalidInput    = errors.New("invalid input")
	ErrRateLimited     = errors.New("rate limit exceeded")
	ErrInvalidResponse = errors.New("invalid response format")
)

// HTTP status codes (aligned with typical REST/GraphQL API usage)
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusConflict            = 409
	StatusUnprocessableEntity = 422
	StatusTooManyRequests     = 429
	StatusInternalServerError = 500
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
	StatusGatewayTimeout      = 504
)

// API error codes (for parsed JSON error responses)
const (
	ErrorCodeBadRequest       = "BadRequest"
	ErrorCodeUnauthorized     = "Unauthorized"
	ErrorCodeForbidden        = "Forbidden"
	ErrorCodeNotFound         = "NotFound"
	ErrorCodeConflict         = "Conflict"
	ErrorCodeValidation       = "ValidationError"
	ErrorCodeRateLimited      = "RateLimitExceeded"
	ErrorCodeInternal         = "InternalError"
	ErrorCodeTransient        = "TransientError"
	ErrorCodeDeadlineExceeded = "DeadlineExceeded"
	ErrorCodeGraphQL          = "GraphQL"
)

// APIError represents an error response from the Jamf Protect API (HTTP or GraphQL layer).
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`

	StatusCode int
	Status     string
	Endpoint   string
	Method     string
}

// formatGraphQLPath converts a GraphQL error path into a readable string format.
func formatGraphQLPath(path []any) string {
	parts := make([]string, 0, len(path))
	for _, p := range path {
		switch v := p.(type) {
		case string:
			parts = append(parts, v)
		case float64:
			parts = append(parts, fmt.Sprintf("%d", int64(v)))
		default:
			parts = append(parts, fmt.Sprintf("%v", v))
		}
	}
	return strings.Join(parts, ".")
}

// formatGraphQLExtensions converts the extensions map of a GraphQL error into a JSON string.
func formatGraphQLExtensions(ext map[string]any) string {
	if len(ext) == 0 {
		return ""
	}
	data, err := json.Marshal(ext)
	if err != nil {
		return ""
	}
	return string(data)
}

// formatGraphQLLocations converts a slice of GraphQLLocation into a readable string format.
func formatGraphQLLocations(locations []GraphQLLocation) string {
	parts := make([]string, 0, len(locations))
	for _, loc := range locations {
		parts = append(parts, fmt.Sprintf("%d:%d", loc.Line, loc.Column))
	}
	return strings.Join(parts, ", ")
}

// MapGraphQLErrors converts a slice of GraphQLError into a single error.
// It returns an *APIError so that IsGraphQL, IsNotFound, and GetErrorCode in errors.go work correctly.
func MapGraphQLErrors(errs []GraphQLError) error {
	if len(errs) == 0 {
		return nil
	}

	messages := make([]string, 0, len(errs))
	isNotFound := false

	for _, e := range errs {
		if e.Message == "" {
			continue
		}

		lower := strings.ToLower(e.Message)
		if strings.Contains(lower, "not found") || strings.Contains(lower, "not_found") {
			isNotFound = true
		}

		msg := e.Message
		if e.ErrorType != "" {
			msg = fmt.Sprintf("%s: %s", e.ErrorType, msg)
		}
		if len(e.Path) > 0 {
			msg = fmt.Sprintf("%s (path: %s)", msg, formatGraphQLPath(e.Path))
		}
		if len(e.Locations) > 0 {
			msg = fmt.Sprintf("%s (locations: %s)", msg, formatGraphQLLocations(e.Locations))
		}
		if ext := formatGraphQLExtensions(e.Extensions); ext != "" {
			msg = fmt.Sprintf("%s (extensions: %s)", msg, ext)
		}
		messages = append(messages, msg)
	}

	if len(messages) == 0 {
		return NewAPIErrorFromGraphQL("graphql operation failed", false)
	}

	errMsg := strings.Join(messages, "; ")
	return NewAPIErrorFromGraphQL(errMsg, isNotFound)
}

// jamfErrorResponse is a common wrapper for API error JSON (e.g. {"error": {"code": "...", "message": "..."}}).
type jamfErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("Jamf Protect API error (%d %s) [%s] at %s %s: %s",
			e.StatusCode, e.Status, e.Code, e.Method, e.Endpoint, e.Message)
	}
	return fmt.Sprintf("Jamf Protect API error (%d %s) at %s %s: %s",
		e.StatusCode, e.Status, e.Method, e.Endpoint, e.Message)
}

// NewAPIErrorFromGraphQL builds an APIError from GraphQL errors so that IsNotFound, GetErrorCode, and IsGraphQL work.
// When isNotFound is true, StatusCode is set to StatusNotFound so IsNotFound(err) returns true.
func NewAPIErrorFromGraphQL(messages string, isNotFound bool) *APIError {
	code := ErrorCodeGraphQL
	statusCode := StatusBadRequest
	if isNotFound {
		statusCode = StatusNotFound
	}
	return &APIError{
		Code:       code,
		Message:    messages,
		StatusCode: statusCode,
		Status:     "",
		Endpoint:   "",
		Method:     "POST",
	}
}

// ParseErrorResponse parses an error response body and returns an APIError.
// If the body is not a recognized error format, Message is set from the raw body or a default for statusCode.
func ParseErrorResponse(body []byte, statusCode int, status, method, endpoint string, logger *zap.Logger) error {
	apiError := &APIError{
		StatusCode: statusCode,
		Status:     status,
		Endpoint:   endpoint,
		Method:     method,
	}

	var jerr jamfErrorResponse
	if err := json.Unmarshal(body, &jerr); err == nil && (jerr.Error.Code != "" || jerr.Error.Message != "") {
		apiError.Code = jerr.Error.Code
		apiError.Message = jerr.Error.Message
		logger.Error("API error response",
			zap.Int("status_code", statusCode),
			zap.String("status", status),
			zap.String("method", method),
			zap.String("endpoint", endpoint),
			zap.String("error_code", apiError.Code),
			zap.String("message", apiError.Message))
	} else {
		apiError.Message = string(body)
		if apiError.Message == "" {
			apiError.Message = getDefaultErrorMessage(statusCode)
		}
		logger.Error("API error response",
			zap.Int("status_code", statusCode),
			zap.String("status", status),
			zap.String("method", method),
			zap.String("endpoint", endpoint),
			zap.String("message", apiError.Message))
	}
	return apiError
}

func getDefaultErrorMessage(statusCode int) string {
	switch statusCode {
	case StatusBadRequest:
		return "The request is invalid or malformed."
	case StatusUnauthorized:
		return "Authentication required or token invalid."
	case StatusForbidden:
		return "You are not allowed to perform the requested operation."
	case StatusNotFound:
		return "The requested resource was not found."
	case StatusConflict:
		return "The resource already exists or conflicts with current state."
	case StatusUnprocessableEntity:
		return "Validation error."
	case StatusTooManyRequests:
		return "Rate limit exceeded. Retry after the indicated period."
	case StatusInternalServerError:
		return "Internal server error."
	case StatusBadGateway:
		return "Bad gateway."
	case StatusServiceUnavailable:
		return "Service temporarily unavailable. Retry later."
	case StatusGatewayTimeout:
		return "The operation took too long to complete."
	default:
		return "Unknown error."
	}
}

// Is* helpers for APIError

// IsBadRequest returns true if the error is a 400 Bad Request
func IsBadRequest(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == StatusBadRequest
}

// IsUnauthorized returns true if the error is 401 Unauthorized
func IsUnauthorized(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == StatusUnauthorized
}

// IsForbidden returns true if the error is 403 Forbidden
func IsForbidden(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == StatusForbidden
}

// IsNotFound returns true if the error is 404 Not Found
func IsNotFound(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == StatusNotFound
}

// IsConflict returns true if the error is 409 Conflict
func IsConflict(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == StatusConflict
}

// IsValidationError returns true if the error is 422 Unprocessable Entity
func IsValidationError(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == StatusUnprocessableEntity
}

// IsRateLimited returns true if the error is 429 Too Many Requests
func IsRateLimited(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == StatusTooManyRequests
}

// IsServerError returns true if the error is 5xx
func IsServerError(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode >= 500 && e.StatusCode < 600
}

// IsTransient returns true if the error is typically retryable (503, 504, or TransientError code)
func IsTransient(err error) bool {
	var e *APIError
	if !errors.As(err, &e) {
		return false
	}
	return e.Code == ErrorCodeTransient ||
		e.StatusCode == StatusServiceUnavailable ||
		e.StatusCode == StatusGatewayTimeout
}

// GetErrorCode returns the API error code from an APIError, or empty string
func GetErrorCode(err error) string {
	var e *APIError
	if errors.As(err, &e) {
		return e.Code
	}
	return ""
}

// IsGraphQL returns true if the error is from GraphQL errors (Code == ErrorCodeGraphQL)
func IsGraphQL(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.Code == ErrorCodeGraphQL
}
