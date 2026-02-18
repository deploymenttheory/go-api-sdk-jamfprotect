package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// GraphQLRequest represents a GraphQL request payload.
type GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

// GraphQLPost sends a GraphQL query or mutation via HTTP POST.
// Path is supplied by the caller (e.g. service CRUD). It uses Post and adds GraphQL response handling (errors in body, unmarshal data into target).
// Headers are applied if provided (nil allowed). Returns the HTTP response and any error; response is non-nil on error.
func (t *Transport) GraphQLPost(ctx context.Context, path string, query string, variables map[string]any, target any, headers map[string]string) (*interfaces.Response, error) {
	if path == "" {
		return nil, fmt.Errorf("%w: path is required", ErrInvalidInput)
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	payload := GraphQLRequest{Query: query, Variables: variables}
	var gqlResp GraphQLResponse

	ifaceResp, err := t.Post(ctx, path, payload, headers, &gqlResp)
	if err != nil {
		return ifaceResp, err
	}

	if err := MapGraphQLErrors(gqlResp.Errors); err != nil {
		return ifaceResp, err
	}

	if target == nil || len(gqlResp.Data) == 0 {
		return ifaceResp, nil
	}

	if err := json.Unmarshal(gqlResp.Data, target); err != nil {
		return ifaceResp, fmt.Errorf("decoding graphql response: %w", err)
	}

	return ifaceResp, nil
}

// Post executes a POST request with JSON body. Auth is applied automatically via middleware.
func (t *Transport) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	req := t.client.R().
		SetContext(ctx).
		SetResult(result)

	if body != nil {
		req.SetBody(body)
	}

	t.applyHeaders(req, headers)

	return t.executeRequest(req, "POST", path)
}

// executeRequest is a centralized request executor that handles error processing.
// Returns response metadata and error. Response is always non-nil for accessing headers.
func (t *Transport) executeRequest(req *resty.Request, method, path string) (*interfaces.Response, error) {
	t.logger.Debug("Executing API request",
		zap.String("method", method),
		zap.String("path", path))

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = req.Get(path)
	case "POST":
		resp, err = req.Post(path)
	default:
		return toInterfaceResponse(nil), fmt.Errorf("unsupported HTTP method: %s", method)
	}

	ifaceResp := toInterfaceResponse(resp)

	if err != nil {
		t.logger.Error("Request failed",
			zap.String("method", method),
			zap.String("path", path),
			zap.Error(err))
		return ifaceResp, fmt.Errorf("request failed: %w", err)
	}

	if err := t.validateResponse(resp, method, path); err != nil {
		return ifaceResp, err
	}

	if !IsResponseSuccess(ifaceResp) {
		return ifaceResp, ParseErrorResponse(
			ifaceResp.Body,
			ifaceResp.StatusCode,
			ifaceResp.Status,
			method,
			path,
			t.logger,
		)
	}

	t.logger.Debug("Request completed successfully",
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status_code", ifaceResp.StatusCode))

	return ifaceResp, nil
}
