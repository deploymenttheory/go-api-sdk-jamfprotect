package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
)

// GraphQLRequest represents a GraphQL request payload.
type GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

// GraphQLPost sends a GraphQL query or mutation via HTTP POST.
// Path is supplied by the caller (e.g. service CRUD). Headers are applied if provided (nil allowed).
// Returns the HTTP response and any error; response is non-nil on error.
func (t *Transport) GraphQLPost(ctx context.Context, path string, query string, variables map[string]any, target any, headers map[string]string) (*interfaces.Response, error) {
	if path == "" {
		return nil, fmt.Errorf("%w: path is required", ErrInvalidInput)
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	payload := GraphQLRequest{Query: query, Variables: variables}
	var gqlResp GraphQLResponse

	clientResp, err := t.Post(ctx, path, payload, headers, &gqlResp)
	if err != nil {
		return clientResp, err
	}

	if err := MapGraphQLErrors(gqlResp.Errors); err != nil {
		return clientResp, err
	}

	if target == nil || len(gqlResp.Data) == 0 {
		return clientResp, nil
	}

	if err := json.Unmarshal(gqlResp.Data, target); err != nil {
		return clientResp, fmt.Errorf("decoding graphql response: %w", err)
	}

	return clientResp, nil
}
