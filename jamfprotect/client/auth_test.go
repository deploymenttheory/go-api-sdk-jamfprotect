package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewAuthManager(t *testing.T) {
	config := AuthConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		TokenURL:     "https://example.com/token",
	}

	httpClient := &http.Client{}
	logger, _ := newTestLogger()

	manager := newAuthManager(config, httpClient, logger)

	assert.NotNil(t, manager)
	assert.Equal(t, config.ClientID, manager.config.ClientID)
	assert.Equal(t, config.ClientSecret, manager.config.ClientSecret)
	assert.NotNil(t, manager.httpClient)
	assert.NotNil(t, manager.logger)
}

// Helper function to create a test logger
func newTestLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func TestAuthManager_GetToken_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/token", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		tokenResp := TokenResponse{
			AccessToken: "test-access-token",
			TokenType:   "Bearer",
			ExpiresIn:   3600,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResp)
	}))
	defer server.Close()

	config := AuthConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		TokenURL:     server.URL + "/token",
	}

	logger, _ := newTestLogger()
	manager := newAuthManager(config, http.DefaultClient, logger)

	token, err := manager.GetToken(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, "test-access-token", token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)
	assert.False(t, token.Expiry.IsZero())
}

func TestAuthManager_GetToken_CachesToken(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		tokenResp := TokenResponse{
			AccessToken: "cached-token",
			TokenType:   "Bearer",
			ExpiresIn:   3600,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResp)
	}))
	defer server.Close()

	config := AuthConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		TokenURL:     server.URL + "/token",
	}

	logger, _ := newTestLogger()
	manager := newAuthManager(config, http.DefaultClient, logger)

	// First call should hit the server
	token1, err := manager.GetToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, callCount)

	// Second call should use cached token
	token2, err := manager.GetToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, callCount, "Expected token to be cached")
	assert.Equal(t, token1.AccessToken, token2.AccessToken)
}

func TestAuthManager_GetToken_RefreshesExpiredToken(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		tokenResp := TokenResponse{
			AccessToken: "refreshed-token",
			TokenType:   "Bearer",
			ExpiresIn:   1, // Short expiry for testing
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResp)
	}))
	defer server.Close()

	config := AuthConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		TokenURL:     server.URL + "/token",
	}

	logger, _ := newTestLogger()
	manager := newAuthManager(config, http.DefaultClient, logger)

	// Get initial token
	token1, err := manager.GetToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, callCount)

	// Wait for token to expire
	time.Sleep(2 * time.Second)

	// Should refresh token
	token2, err := manager.GetToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 2, callCount, "Expected token to be refreshed")
	assert.NotEqual(t, token1, token2)
}

func TestAuthManager_GetToken_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"invalid_client"}`))
	}))
	defer server.Close()

	config := AuthConfig{
		ClientID:     "invalid-client",
		ClientSecret: "invalid-secret",
		TokenURL:     server.URL + "/token",
	}

	logger, _ := newTestLogger()
	manager := newAuthManager(config, http.DefaultClient, logger)

	token, err := manager.GetToken(context.Background())

	assert.Error(t, err)
	assert.Nil(t, token)
	assert.Contains(t, err.Error(), "401")
}

func TestAuthManager_CurrentToken_Empty(t *testing.T) {
	config := AuthConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		TokenURL:     "https://example.com/token",
	}

	logger, _ := newTestLogger()
	manager := newAuthManager(config, http.DefaultClient, logger)

	token := manager.currentToken()
	assert.Nil(t, token)
}

func TestAuthManager_ConcurrentGetToken(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		time.Sleep(100 * time.Millisecond) // Simulate slow token fetch

		tokenResp := TokenResponse{
			AccessToken: "concurrent-token",
			TokenType:   "Bearer",
			ExpiresIn:   3600,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResp)
	}))
	defer server.Close()

	config := AuthConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		TokenURL:     server.URL + "/token",
	}

	logger, _ := newTestLogger()
	manager := newAuthManager(config, http.DefaultClient, logger)

	// Launch multiple concurrent requests
	const numConcurrent = 10
	tokens := make([]*http.Cookie, numConcurrent)
	errors := make([]error, numConcurrent)
	done := make(chan struct{}, numConcurrent)

	for i := 0; i < numConcurrent; i++ {
		go func(idx int) {
			defer func() { done <- struct{}{} }()
			token, err := manager.GetToken(context.Background())
			errors[idx] = err
			if token != nil {
				// Store something to verify token was returned
				tokens[idx] = &http.Cookie{Name: token.AccessToken}
			}
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < numConcurrent; i++ {
		<-done
	}

	// Verify all requests succeeded
	for i := 0; i < numConcurrent; i++ {
		assert.NoError(t, errors[i], "Request %d failed", i)
		assert.NotNil(t, tokens[i], "Request %d didn't get a token", i)
	}

	// singleflight should ensure only one actual HTTP call was made
	assert.Equal(t, 1, callCount, "Expected only one token fetch due to singleflight")
}

func TestRedactTokenRequestBody(t *testing.T) {
	body := redactTokenRequestBody("test-client-id")

	var parsed map[string]string
	err := json.Unmarshal(body, &parsed)
	require.NoError(t, err)

	assert.Equal(t, "test-client-id", parsed["client_id"])
	assert.Equal(t, "[REDACTED]", parsed["password"])
}

func TestRedactTokenResponseBody(t *testing.T) {
	originalBody := []byte(`{"access_token":"secret-token","token_type":"Bearer","expires_in":3600}`)

	redacted := redactTokenResponseBody(originalBody)

	var parsed map[string]any
	err := json.Unmarshal(redacted, &parsed)
	require.NoError(t, err)

	assert.Equal(t, "[REDACTED]", parsed["access_token"])
	assert.Equal(t, "Bearer", parsed["token_type"])
	assert.Equal(t, float64(3600), parsed["expires_in"])
}

func TestRedactTokenResponseBody_Invalid(t *testing.T) {
	// Test with invalid JSON
	invalidJSON := []byte(`{invalid json`)
	result := redactTokenResponseBody(invalidJSON)
	assert.Equal(t, invalidJSON, result, "Should return original on error")
}
