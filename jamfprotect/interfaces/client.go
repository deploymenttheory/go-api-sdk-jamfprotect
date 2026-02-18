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
	StatusCode int           // HTTP status code
	Status     string        // HTTP status text
	Headers    http.Header   // Response headers
	Body       []byte        // Raw response body
	Duration   time.Duration // Time taken for the request
	ReceivedAt time.Time     // When the response was received
	Size       int64         // Response body size in bytes
}

// GraphQLClient interface that services will use.
// This breaks import cycles by providing a contract without implementation.
type GraphQLClient interface {
	// GraphQLPost sends a GraphQL query or mutation via HTTP POST.
	// Path is supplied by the caller (service CRUD), e.g. client.EndpointApp.
	// Headers are applied if provided (nil is allowed). Returns the raw HTTP response metadata and any error.
	// Response is non-nil even on error for accessing headers.
	GraphQLPost(
		ctx context.Context, // request context
		path string, // API path (caller-owned, e.g. from service constants)
		query string, // GraphQL query or mutation string
		variables map[string]any, // GraphQL variables
		target any, // pointer to unmarshal response into
		headers map[string]string, // HTTP headers (nil allowed)
	) (*Response, error)

	// GetLogger returns the configured zap logger instance.
	GetLogger() *zap.Logger
}
