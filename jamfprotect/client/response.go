package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// GraphQLResponse represents a GraphQL response payload, including any errors.
type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []GraphQLError  `json:"errors"`
}

// GraphQLError represents an individual error returned by the GraphQL API.
// Jamf Protect may include errorType (e.g. ArgumentValidationError), data, and errorInfo.
type GraphQLError struct {
	Message    string            `json:"message"`
	Path       []any             `json:"path,omitempty"`
	Data       json.RawMessage   `json:"data,omitempty"`
	ErrorType  string            `json:"errorType,omitempty"`
	ErrorInfo  map[string]any    `json:"errorInfo,omitempty"`
	Locations  []GraphQLLocation `json:"locations,omitempty"`
	Extensions map[string]any    `json:"extensions,omitempty"`
}

// GraphQLLocation represents the line and column of an error in a GraphQL query.
type GraphQLLocation struct {
	Line       int    `json:"line"`
	Column     int    `json:"column"`
	SourceName string `json:"sourceName,omitempty"`
}

// toInterfaceResponse converts a resty.Response to interfaces.Response
func toInterfaceResponse(resp *resty.Response) *interfaces.Response {
	if resp == nil {
		return &interfaces.Response{
			Headers: make(http.Header),
		}
	}
	return &interfaces.Response{
		StatusCode: resp.StatusCode(),
		Status:     resp.Status(),
		Headers:    resp.Header(),
		Body:       []byte(resp.String()),
		Duration:   resp.Duration(),
		ReceivedAt: resp.ReceivedAt(),
		Size:       resp.Size(),
	}
}

// validateResponse validates the HTTP response before processing.
func (t *Transport) validateResponse(resp *resty.Response, method, path string) error {
	bodyLen := len(resp.String())
	if resp.Header().Get("Content-Length") == "0" || bodyLen == 0 {
		t.logger.Debug("Empty response received",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status_code", resp.StatusCode()))
		return nil
	}
	if !IsResponseError(toInterfaceResponse(resp)) && bodyLen > 0 {
		contentType := resp.Header().Get("Content-Type")
		if contentType != "" && !strings.HasPrefix(contentType, "application/json") {
			t.logger.Warn("Unexpected Content-Type in response",
				zap.String("method", method),
				zap.String("path", path),
				zap.String("content_type", contentType),
				zap.String("expected", "application/json"))
			return fmt.Errorf("unexpected response Content-Type from %s %s: got %q, expected application/json",
				method, path, contentType)
		}
	}
	return nil
}

// Response helpers for working with interfaces.Response

// IsResponseSuccess returns true if the response status code is 2xx
func IsResponseSuccess(resp *interfaces.Response) bool {
	if resp == nil {
		return false
	}
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

// IsResponseError returns true if the response status code is 4xx or 5xx
func IsResponseError(resp *interfaces.Response) bool {
	if resp == nil {
		return false
	}
	return resp.StatusCode >= 400
}

// GetResponseHeader returns a header value from the response by key (case-insensitive)
func GetResponseHeader(resp *interfaces.Response, key string) string {
	if resp == nil || resp.Headers == nil {
		return ""
	}
	return resp.Headers.Get(key)
}

// GetResponseHeaders returns all headers from the response
func GetResponseHeaders(resp *interfaces.Response) http.Header {
	if resp == nil {
		return make(http.Header)
	}
	return resp.Headers
}

// GetRateLimitHeaders extracts rate limit–related headers from the response.
// Returns (limit, remaining, reset, retryAfter). Jamf Protect may not set these; empty strings if absent.
func GetRateLimitHeaders(resp *interfaces.Response) (limit, remaining, reset, retryAfter string) {
	if resp == nil || resp.Headers == nil {
		return "", "", "", ""
	}
	return resp.Headers.Get("X-RateLimit-Limit"),
		resp.Headers.Get("X-RateLimit-Remaining"),
		resp.Headers.Get("X-RateLimit-Reset"),
		resp.Headers.Get("Retry-After")
}
