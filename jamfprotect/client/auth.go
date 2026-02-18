package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
	"resty.dev/v3"
)

// TokenRequest is the payload sent to the /token endpoint
type TokenRequest struct {
	ClientID string `json:"client_id"`
	Password string `json:"password"`
}

// TokenResponse is the response from the /token endpoint
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}

// AuthConfig holds OAuth2 client credentials configuration
type AuthConfig struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
}

// Validate checks if the auth configuration is valid
func (a *AuthConfig) Validate() error {
	if a.ClientID == "" {
		return fmt.Errorf("client ID is required")
	}
	if a.ClientSecret == "" {
		return fmt.Errorf("client secret is required")
	}
	if a.TokenURL == "" {
		return fmt.Errorf("token URL is required")
	}
	return nil
}

// TokenManager handles OAuth2 token lifecycle
type TokenManager struct {
	authConfig    *AuthConfig
	httpClient    *http.Client
	logger        *zap.Logger
	currentToken  *TokenResponse
	tokenExpiry   time.Time
	mu            sync.RWMutex
	refreshBuffer time.Duration
}

// NewTokenManager creates a new token manager
func NewTokenManager(authConfig *AuthConfig, httpClient *http.Client, logger *zap.Logger) *TokenManager {
	return &TokenManager{
		authConfig:    authConfig,
		httpClient:    httpClient,
		logger:        logger,
		refreshBuffer: time.Duration(TokenExpirySkew) * time.Second,
	}
}

// GetToken returns a valid access token, refreshing if necessary
func (tm *TokenManager) GetToken(ctx context.Context) (string, error) {
	tm.mu.RLock()
	if tm.currentToken != nil && time.Now().Add(tm.refreshBuffer).Before(tm.tokenExpiry) {
		token := tm.currentToken.AccessToken
		tm.mu.RUnlock()
		return token, nil
	}
	tm.mu.RUnlock()

	return tm.RefreshToken(ctx)
}

// RefreshToken requests a new access token from the OAuth2 endpoint (thread-safe)
func (tm *TokenManager) RefreshToken(ctx context.Context) (string, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// Double-check in case another goroutine just refreshed
	if tm.currentToken != nil && time.Now().Add(tm.refreshBuffer).Before(tm.tokenExpiry) {
		return tm.currentToken.AccessToken, nil
	}

	if tm.logger != nil {
		tm.logger.Info("Requesting new OAuth2 access token",
			zap.String("token_url", tm.authConfig.TokenURL))
	}

	tokenResp, err := tm.fetchToken(ctx)
	if err != nil {
		return "", err
	}

	tm.currentToken = tokenResp
	tm.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	if tm.logger != nil {
		tm.logger.Info("Successfully obtained access token",
			zap.String("token_type", tokenResp.TokenType),
			zap.Int64("expires_in", tokenResp.ExpiresIn),
			zap.Time("expires_at", tm.tokenExpiry))
	}

	return tokenResp.AccessToken, nil
}

// InvalidateToken clears the current token, forcing a refresh on next use
func (tm *TokenManager) InvalidateToken() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.currentToken = nil
	tm.tokenExpiry = time.Time{}

	if tm.logger != nil {
		tm.logger.Info("Access token invalidated")
	}
}

// fetchToken performs the actual HTTP request to obtain a new access token
func (tm *TokenManager) fetchToken(ctx context.Context) (*TokenResponse, error) {
	body, err := json.Marshal(TokenRequest{
		ClientID: tm.authConfig.ClientID,
		Password: tm.authConfig.ClientSecret,
	})
	if err != nil {
		return nil, fmt.Errorf("marshalling token request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		tm.authConfig.TokenURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("creating token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", DefaultUserAgent)

	if tm.logger != nil {
		tm.logger.Debug("OAuth2 token request",
			zap.String("method", http.MethodPost),
			zap.String("url", req.URL.String()),
			zap.ByteString("body", redactTokenRequestBody(tm.authConfig.ClientID)))
	}

	resp, err := tm.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting token: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading token response: %w", err)
	}

	if tm.logger != nil {
		tm.logger.Debug("OAuth2 token response",
			zap.Int("status_code", resp.StatusCode),
			zap.ByteString("body", redactTokenResponseBody(respBody)))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token request returned %d: %s", resp.StatusCode, string(respBody))
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(respBody, &tokenResp); err != nil {
		return nil, fmt.Errorf("decoding token response: %w", err)
	}

	if tokenResp.AccessToken == "" {
		return nil, fmt.Errorf("%w: token response missing access_token", ErrAuthentication)
	}

	if tokenResp.ExpiresIn <= 0 {
		return nil, fmt.Errorf("%w: token response missing expires_in", ErrAuthentication)
	}

	return &tokenResp, nil
}

// SetupAuthentication configures the resty client with OAuth2 bearer token authentication
func SetupAuthentication(client *resty.Client, authConfig *AuthConfig, logger *zap.Logger) (*TokenManager, error) {
	if err := authConfig.Validate(); err != nil {
		if logger != nil {
			logger.Error("Authentication validation failed", zap.Error(err))
		}
		return nil, fmt.Errorf("authentication validation failed: %w", err)
	}

	// Use the underlying *http.Client for token requests so they bypass resty middleware
	tokenManager := NewTokenManager(authConfig, client.Client(), logger)

	// Fetch initial token
	token, err := tokenManager.RefreshToken(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to obtain initial access token: %w", err)
	}

	client.SetAuthToken(token)

	// Add request middleware to ensure token is valid before each request
	client.AddRequestMiddleware(func(c *resty.Client, req *resty.Request) error {
		token, err := tokenManager.GetToken(req.Context())
		if err != nil {
			if logger != nil {
				logger.Error("Failed to get valid token for request", zap.Error(err))
			}
			return fmt.Errorf("%w: %w", ErrAuthentication, err)
		}
		req.SetAuthToken(token)
		return nil
	})

	if logger != nil {
		logger.Info("OAuth2 authentication configured successfully",
			zap.String("token_url", authConfig.TokenURL))
	}

	return tokenManager, nil
}

// redactTokenRequestBody creates a redacted version of the token request body for logging
func redactTokenRequestBody(clientID string) []byte {
	data, err := json.Marshal(map[string]string{
		"client_id": clientID,
		"password":  "[REDACTED]",
	})
	if err != nil {
		return []byte(`{"password":"[REDACTED]"}`)
	}
	return data
}

// redactTokenResponseBody creates a redacted version of the token response body for logging
func redactTokenResponseBody(body []byte) []byte {
	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		return body
	}
	if _, ok := payload["access_token"]; ok {
		payload["access_token"] = "[REDACTED]"
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return body
	}
	return data
}
