package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// Ensure Transport implements interfaces.GraphQLClient
var _ interfaces.GraphQLClient = (*Transport)(nil)

// Transport is the HTTP transport layer for the Jamf Protect GraphQL API
type Transport struct {
	client        *resty.Client
	baseURL       string
	userAgent     string
	logger        *zap.Logger
	authManager   *authManager
	globalHeaders map[string]string
}

// NewTransport creates a new Jamf Protect GraphQL transport client with default settings
func NewTransport(clientID, clientSecret string, options ...ClientOption) (*Transport, error) {
	return NewTransportWithVersion(clientID, clientSecret, Version, options...)
}

// NewTransportWithVersion creates a new Jamf Protect GraphQL transport client with a custom version string
func NewTransportWithVersion(clientID, clientSecret, version string, options ...ClientOption) (*Transport, error) {
	if clientID == "" {
		return nil, fmt.Errorf("%w: clientID is required", ErrInvalidInput)
	}
	if clientSecret == "" {
		return nil, fmt.Errorf("%w: clientSecret is required", ErrInvalidInput)
	}

	// Set up default logger first
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create default logger: %w", err)
	}

	// Format user agent with version
	userAgent := fmt.Sprintf("%s/%s", UserAgentBase, version)

	// Create resty client with defaults
	restyClient := resty.New()
	restyClient.SetTimeout(time.Duration(DefaultTimeout) * time.Second)
	restyClient.SetRetryCount(MaxRetries)
	restyClient.SetRetryWaitTime(time.Duration(RetryWaitTime) * time.Second)
	restyClient.SetRetryMaxWaitTime(time.Duration(RetryMaxWaitTime) * time.Second)
	restyClient.SetHeader(HeaderUserAgent, userAgent)
	restyClient.SetHeader(HeaderContentType, ContentTypeJSON)
	restyClient.SetHeader("Accept", AcceptJSON)

	// Create default transport with global headers
	globalHeaders := map[string]string{
		HeaderContentType: ContentTypeJSON,
		"Accept":          AcceptJSON,
	}
	
	transport := &Transport{
		client:        restyClient,
		baseURL:       DefaultBaseURL,
		userAgent:     userAgent,
		logger:        logger,
		globalHeaders: globalHeaders,
	}

	// Set base URL for resty client
	restyClient.SetBaseURL(transport.baseURL)

	// Apply functional options (can override logger and base URL)
	for _, opt := range options {
		if err := opt(transport); err != nil {
			return nil, fmt.Errorf("applying client option: %w", err)
		}
	}

	// Initialize auth manager
	authConfig := AuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     strings.TrimRight(transport.baseURL, "/") + EndpointToken,
	}
	transport.authManager = newAuthManager(authConfig, restyClient.Client(), transport.logger)

	transport.logger.Info("Jamf Protect API client created",
		zap.String("base_url", transport.baseURL),
		zap.String("client_id", clientID),
	)

	return transport, nil
}

// DoGraphQL executes a GraphQL query/mutation against a specified endpoint path
// Optional headers parameter allows per-request header customization (overrides global headers)
// Returns the result, the raw HTTP response, and any error
func (t *Transport) DoGraphQL(ctx context.Context, endpoint, query string, variables map[string]any, target any, headers ...map[string]string) (*interfaces.Response, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("%w: endpoint path is required", ErrInvalidInput)
	}
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	// Get valid access token
	token, err := t.authManager.GetToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrAuthentication, err)
	}

	// Prepare GraphQL request
	payload := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	// Create resty request
	var gqlResp GraphQLResponse
	req := t.client.R().
		SetContext(ctx).
		SetBody(payload).
		SetResult(&gqlResp)

	// Apply headers (global first, then per-request)
	// Authorization is always required for GraphQL
	requestHeaders := map[string]string{
		HeaderAuthorization: token.AccessToken,
	}
	
	// Merge any additional headers from the call
	if len(headers) > 0 && headers[0] != nil {
		for k, v := range headers[0] {
			requestHeaders[k] = v
		}
	}
	
	t.applyHeaders(req, requestHeaders)

	resp, err := req.Post(endpoint)
	apiResp := toInterfaceResponse(resp)
	
	if err != nil {
		return apiResp, fmt.Errorf("executing graphql request: %w", err)
	}

	// Validate response
	if err := t.validateResponse(resp, "POST", endpoint); err != nil {
		return apiResp, err
	}

	// Check HTTP status
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return apiResp, fmt.Errorf("graphql request returned %d: %s", resp.StatusCode(), resp.String())
	}

	// Check for GraphQL errors
	if err := MapGraphQLErrors(gqlResp.Errors); err != nil {
		return apiResp, err
	}

	// Unmarshal data into target
	if target == nil || len(gqlResp.Data) == 0 {
		return apiResp, nil
	}
	if err := json.Unmarshal(gqlResp.Data, target); err != nil {
		return apiResp, fmt.Errorf("decoding graphql data: %w", err)
	}
	return apiResp, nil
}

// GetHTTPClient returns the underlying resty client
func (t *Transport) GetHTTPClient() *resty.Client {
	return t.client
}

// GetLogger returns the configured zap logger
func (t *Transport) GetLogger() *zap.Logger {
	return t.logger
}

// SetLogger updates the logger at runtime
func (t *Transport) SetLogger(logger *zap.Logger) {
	if logger != nil {
		t.logger = logger
		t.authManager.logger = logger
	}
}

// AccessToken retrieves a valid access token, refreshing if necessary
func (t *Transport) AccessToken(ctx context.Context) (string, error) {
	token, err := t.authManager.GetToken(ctx)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrAuthentication, err)
	}
	return token.AccessToken, nil
}

