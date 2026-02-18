package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
	"go.uber.org/zap"
	"resty.dev/v3"
)

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

	clientResp := toInterfaceResponse(resp)

	if err != nil {
		t.logger.Error("Request failed",
			zap.String("method", method),
			zap.String("path", path),
			zap.Error(err))
		return clientResp, fmt.Errorf("request failed: %w", err)
	}

	if err := t.validateResponse(resp, method, path); err != nil {
		return clientResp, err
	}

	if IsResponseError(clientResp) {
		return clientResp, ParseErrorResponse(
			[]byte(resp.String()),
			resp.StatusCode(),
			resp.Status(),
			method,
			path,
			t.logger,
		)
	}

	t.logger.Debug("Request completed successfully",
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status_code", clientResp.StatusCode))

	return clientResp, nil
}
