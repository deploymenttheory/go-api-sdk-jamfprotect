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
	"golang.org/x/oauth2"
	"golang.org/x/sync/singleflight"
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

// authManager handles token acquisition and refresh
type authManager struct {
	config     AuthConfig
	httpClient *http.Client
	logger     *zap.Logger
	mu         sync.RWMutex
	token      *oauth2.Token
	tokenGroup singleflight.Group
}

// newAuthManager creates a new authentication manager
func newAuthManager(config AuthConfig, httpClient *http.Client, logger *zap.Logger) *authManager {
	return &authManager{
		config:     config,
		httpClient: httpClient,
		logger:     logger,
	}
}

// GetToken ensures a valid token is available and returns it
func (am *authManager) GetToken(ctx context.Context) (*oauth2.Token, error) {
	token, err := am.authenticate(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrAuthentication, err)
	}
	return token, nil
}

// authenticate obtains (or refreshes) an access token (thread-safe)
func (am *authManager) authenticate(ctx context.Context) (*oauth2.Token, error) {
	if token := am.currentToken(); token != nil {
		return token, nil
	}

	value, err, _ := am.tokenGroup.Do("token", func() (any, error) {
		if token := am.currentToken(); token != nil {
			return token, nil
		}
		token, err := am.fetchToken(ctx)
		if err != nil {
			return nil, err
		}
		am.mu.Lock()
		am.token = token
		am.mu.Unlock()
		return token, nil
	})
	if err != nil {
		return nil, err
	}
	token, ok := value.(*oauth2.Token)
	if !ok {
		return nil, fmt.Errorf("unexpected token type %T", value)
	}
	return token, nil
}

// currentToken returns the current token if it's valid, or nil if it's missing or expired
func (am *authManager) currentToken() *oauth2.Token {
	am.mu.RLock()
	defer am.mu.RUnlock()
	if am.token != nil && am.token.Valid() {
		return am.token
	}
	return nil
}

// fetchToken performs the actual HTTP request to obtain a new access token
func (am *authManager) fetchToken(ctx context.Context) (*oauth2.Token, error) {
	body, err := json.Marshal(TokenRequest{
		ClientID: am.config.ClientID,
		Password: am.config.ClientSecret,
	})
	if err != nil {
		return nil, fmt.Errorf("marshalling token request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		am.config.TokenURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("creating token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", DefaultUserAgent)

	if am.logger != nil {
		am.logger.Debug("OAuth2 token request",
			zap.String("method", http.MethodPost),
			zap.String("url", req.URL.String()),
			zap.ByteString("body", redactTokenRequestBody(am.config.ClientID)))
	}

	resp, err := am.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting token: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading token response: %w", err)
	}
	if am.logger != nil {
		am.logger.Debug("OAuth2 token response",
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

	expiry := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	if time.Duration(tokenResp.ExpiresIn)*time.Second > time.Duration(TokenExpirySkew)*time.Second {
		expiry = expiry.Add(-time.Duration(TokenExpirySkew) * time.Second)
	}
	return &oauth2.Token{
		AccessToken: tokenResp.AccessToken,
		TokenType:   tokenResp.TokenType,
		Expiry:      expiry,
	}, nil
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
