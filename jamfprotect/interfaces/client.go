package interfaces

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Response represents HTTP response metadata that can be returned alongside errors.
// This allows callers to access response headers (rate limits, retry-after, etc.) even on errors.
type Response struct {
	StatusCode int         // HTTP status code
	Status     string      // HTTP status text
	Headers    http.Header // Response headers
	Body       []byte      // Raw response body
	Duration   time.Duration // Time taken for the request
	ReceivedAt time.Time     // When the response was received
	Size       int64         // Response body size in bytes
}

// GraphQLClient interface that services will use.
// This breaks import cycles by providing a contract without implementation.
type GraphQLClient interface {
	// DoGraphQL executes a GraphQL query/mutation against a specified endpoint path.
	// Optional headers parameter allows per-request header customization (overrides global headers).
	// Returns the raw HTTP response metadata and any error.
	// Response is non-nil even on error for accessing headers.
	DoGraphQL(
		ctx context.Context, // request context
		endpoint string, // GraphQL endpoint path
		query string, // GraphQL query or mutation string
		variables map[string]any, // GraphQL variables
		target any, // pointer to unmarshal response into
		headers ...map[string]string, // optional per-request headers
	) (*Response, error)

	// GetLogger returns the configured zap logger instance.
	GetLogger() *zap.Logger
}
