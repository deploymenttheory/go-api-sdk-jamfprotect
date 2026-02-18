package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// Ensure Transport implements interfaces.GraphQLClient
var _ interfaces.GraphQLClient = (*Transport)(nil)

// Transport is the HTTP transport layer for the Jamf Protect GraphQL API.
// It provides authentication, retry, and logging. Request execution is in request.go.
type Transport struct {
	client        *resty.Client
	baseURL       string
	userAgent     string
	logger        *zap.Logger
	authConfig    *AuthConfig
	tokenManager  *TokenManager
	globalHeaders map[string]string
}

// NewTransport creates a new Jamf Protect GraphQL transport.
func NewTransport(clientID, clientSecret string, options ...ClientOption) (*Transport, error) {

	if err := ValidateTransportConfig(clientID, clientSecret); err != nil {
		return nil, fmt.Errorf("invalid transport configuration: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	userAgent := fmt.Sprintf("%s/%s", UserAgentBase, Version)

	restyClient := resty.New()
	restyClient.SetTimeout(time.Duration(DefaultTimeout) * time.Second)
	restyClient.SetRetryCount(MaxRetries)
	restyClient.SetRetryWaitTime(time.Duration(RetryWaitTime) * time.Second)
	restyClient.SetRetryMaxWaitTime(time.Duration(RetryMaxWaitTime) * time.Second)
	restyClient.SetHeader(HeaderUserAgent, userAgent)
	restyClient.SetHeader(HeaderContentType, ContentTypeJSON)
	restyClient.SetHeader("Accept", AcceptJSON)

	transport := &Transport{
		client:        restyClient,
		logger:        logger,
		baseURL:       DefaultBaseURL,
		globalHeaders: make(map[string]string),
		userAgent:     userAgent,
	}

	// Apply options before auth setup so that WithBaseURL is respected in the token URL
	for _, opt := range options {
		if err := opt(transport); err != nil {
			return nil, fmt.Errorf("applying client option: %w", err)
		}
	}

	authConfig := &AuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     strings.TrimRight(transport.baseURL, "/") + EndpointToken,
	}
	transport.authConfig = authConfig

	// Setup OAuth2 authentication
	tokenManager, err := SetupAuthentication(restyClient, authConfig, transport.logger)
	if err != nil {
		return nil, fmt.Errorf("failed to setup authentication: %w", err)
	}
	transport.tokenManager = tokenManager

	restyClient.SetBaseURL(transport.baseURL)

	transport.logger.Info("Jamf Protect API client created",
		zap.String("base_url", transport.baseURL),
		zap.String("client_id", clientID))

	return transport, nil
}

// GetHTTPClient returns the underlying resty client
func (t *Transport) GetHTTPClient() *resty.Client {
	return t.client
}

// GetLogger returns the configured zap logger
func (t *Transport) GetLogger() *zap.Logger {
	return t.logger
}

// GetTokenManager returns the token manager
func (t *Transport) GetTokenManager() *TokenManager {
	return t.tokenManager
}

// SetLogger updates the logger at runtime
func (t *Transport) SetLogger(logger *zap.Logger) {
	if logger != nil {
		t.logger = logger
		t.tokenManager.logger = logger
	}
}

// AccessToken retrieves a valid access token, refreshing if necessary
func (t *Transport) AccessToken(ctx context.Context) (string, error) {
	return t.tokenManager.GetToken(ctx)
}

// RefreshToken manually refreshes the OAuth2 access token
func (t *Transport) RefreshToken(ctx context.Context) error {
	_, err := t.tokenManager.RefreshToken(ctx)
	return err
}

// InvalidateToken invalidates the current token, forcing a refresh on next use
func (t *Transport) InvalidateToken() {
	t.tokenManager.InvalidateToken()
}
