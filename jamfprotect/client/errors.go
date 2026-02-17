package client

import "errors"

// Common error types
var (
	// ErrAuthentication indicates an authentication failure
	ErrAuthentication = errors.New("authentication failed")
	
	// ErrGraphQL indicates a GraphQL operation error
	ErrGraphQL = errors.New("graphql operation failed")
	
	// ErrNotFound indicates a resource was not found
	ErrNotFound = errors.New("resource not found")
	
	// ErrInvalidInput indicates invalid input parameters
	ErrInvalidInput = errors.New("invalid input")
	
	// ErrRateLimited indicates the rate limit has been exceeded
	ErrRateLimited = errors.New("rate limit exceeded")
	
	// ErrInvalidResponse indicates an invalid response from the API
	ErrInvalidResponse = errors.New("invalid response format")
)
