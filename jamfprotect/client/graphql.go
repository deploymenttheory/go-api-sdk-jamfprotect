package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GraphQLRequest represents a GraphQL request payload
type GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

// GraphQLResponse represents a GraphQL response payload, including any errors
type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []GraphQLError  `json:"errors"`
}

// GraphQLError represents an individual error returned by the GraphQL API
type GraphQLError struct {
	Message    string            `json:"message"`
	Locations  []GraphQLLocation `json:"locations,omitempty"`
	Path       []any             `json:"path,omitempty"`
	Extensions map[string]any    `json:"extensions,omitempty"`
}

// GraphQLLocation represents the line and column of an error in a GraphQL query
type GraphQLLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// MapGraphQLErrors converts a slice of GraphQLError into a single error
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
		return ErrGraphQL
	}
	
	errMsg := strings.Join(messages, "; ")
	if isNotFound {
		return fmt.Errorf("%w: %w: %s", ErrNotFound, ErrGraphQL, errMsg)
	}
	return fmt.Errorf("%w: %s", ErrGraphQL, errMsg)
}

// formatGraphQLPath converts a GraphQL error path into a readable string format
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

// formatGraphQLExtensions converts the extensions map of a GraphQL error into a JSON string
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

// formatGraphQLLocations converts a slice of GraphQLLocation into a readable string format
func formatGraphQLLocations(locations []GraphQLLocation) string {
	parts := make([]string, 0, len(locations))
	for _, loc := range locations {
		parts = append(parts, fmt.Sprintf("%d:%d", loc.Line, loc.Column))
	}
	return strings.Join(parts, ", ")
}
