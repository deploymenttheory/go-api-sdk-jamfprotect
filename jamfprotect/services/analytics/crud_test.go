package analytics_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockGraphQLServer creates a test HTTP server that responds to GraphQL requests
func mockGraphQLServer(t *testing.T, handler func(query string, variables map[string]any) (any, error)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle token endpoint
		if r.URL.Path == "/token" {
			tokenResp := map[string]any{
				"access_token": "test-token",
				"expires_in":   3600,
				"token_type":   "Bearer",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tokenResp)
			return
		}

		// Handle GraphQL endpoint
		if r.URL.Path == "/app" {
			var req struct {
				Query     string         `json:"query"`
				Variables map[string]any `json:"variables"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			result, err := handler(req.Query, req.Variables)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := map[string]any{
				"data": result,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		http.NotFound(w, r)
	}))
}

func TestAnalyticsService_CreateAnalytic(t *testing.T) {
	mockData := map[string]any{
		"createAnalytic": map[string]any{
			"uuid":        "analytic-uuid-123",
			"name":        "Test Analytic",
			"label":       "Test Label",
			"inputType":   "unified_log",
			"filter":      "process.name = 'test'",
			"description": "Test Description",
			"level":       3,
			"severity":    "MEDIUM",
			"tags":        []string{"test", "security"},
			"categories":  []string{"malware"},
			"created":     "2024-01-01T00:00:00Z",
			"updated":     "2024-01-01T00:00:00Z",
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "createAnalytic")
		assert.Equal(t, "Test Analytic", variables["name"])
		assert.Equal(t, "unified_log", variables["inputType"])
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := analytics.NewService(transport)

	req := &analytics.CreateAnalyticRequest{
		Name:        "Test Analytic",
		InputType:   "unified_log",
		Description: "Test Description",
		Filter:      "process.name = 'test'",
		Level:       3,
		Severity:    "MEDIUM",
		Tags:        []string{"test", "security"},
		Categories:  []string{"malware"},
		AnalyticActions: []analytics.AnalyticActionInput{
			{
				Name:       "alert",
				Parameters: []string{"high"},
			},
		},
		Context:       []analytics.AnalyticContextInput{},
		SnapshotFiles: []string{},
	}

	result, _, err := service.CreateAnalytic(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "analytic-uuid-123", result.UUID)
	assert.Equal(t, "Test Analytic", result.Name)
	assert.Equal(t, "unified_log", result.InputType)
	assert.Equal(t, 3, result.Level)
	assert.Equal(t, "MEDIUM", result.Severity)
}

func TestAnalyticsService_GetAnalytic(t *testing.T) {
	mockData := map[string]any{
		"getAnalytic": map[string]any{
			"uuid":        "analytic-uuid-123",
			"name":        "Test Analytic",
			"inputType":   "unified_log",
			"filter":      "process.name = 'test'",
			"description": "Test Description",
			"level":       3,
			"severity":    "HIGH",
			"created":     "2024-01-01T00:00:00Z",
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "getAnalytic")
		assert.Equal(t, "analytic-uuid-123", variables["uuid"])
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := analytics.NewService(transport)

	result, _, err := service.GetAnalytic(context.Background(), "analytic-uuid-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "analytic-uuid-123", result.UUID)
	assert.Equal(t, "Test Analytic", result.Name)
	assert.Equal(t, "HIGH", result.Severity)
}

func TestAnalyticsService_UpdateAnalytic(t *testing.T) {
	mockData := map[string]any{
		"updateAnalytic": map[string]any{
			"uuid":        "analytic-uuid-123",
			"name":        "Updated Analytic",
			"description": "Updated Description",
			"level":       5,
			"updated":     "2024-01-02T00:00:00Z",
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "updateAnalytic")
		assert.Equal(t, "analytic-uuid-123", variables["uuid"])
		assert.Equal(t, "Updated Analytic", variables["name"])
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := analytics.NewService(transport)

	req := &analytics.UpdateAnalyticRequest{
		Name:        "Updated Analytic",
		InputType:   "unified_log",
		Description: "Updated Description",
		Filter:      "process.name = 'updated'",
		Level:       5,
		Tags:        []string{"updated"},
		Categories:  []string{"malware"},
		AnalyticActions: []analytics.AnalyticActionInput{
			{Name: "alert", Parameters: []string{"critical"}},
		},
		Context:       []analytics.AnalyticContextInput{},
		SnapshotFiles: []string{},
	}

	result, _, err := service.UpdateAnalytic(context.Background(), "analytic-uuid-123", req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "analytic-uuid-123", result.UUID)
	assert.Equal(t, "Updated Analytic", result.Name)
}

func TestAnalyticsService_DeleteAnalytic(t *testing.T) {
	mockData := map[string]any{
		"deleteAnalytic": map[string]any{
			"uuid": "analytic-uuid-123",
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "deleteAnalytic")
		assert.Equal(t, "analytic-uuid-123", variables["uuid"])
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := analytics.NewService(transport)

	_, err = service.DeleteAnalytic(context.Background(), "analytic-uuid-123")

	require.NoError(t, err)
}

func TestAnalyticsService_ListAnalytics(t *testing.T) {
	mockData := map[string]any{
		"listAnalytics": map[string]any{
			"items": []map[string]any{
				{
					"uuid":        "analytic-1",
					"name":        "Analytic 1",
					"inputType":   "es_event",
					"filter":      "event.type = 'exec'",
					"level":       3,
					"severity":    "MEDIUM",
					"description": "First analytic",
				},
				{
					"uuid":        "analytic-2",
					"name":        "Analytic 2",
					"inputType":   "unified_log",
					"filter":      "process.name = 'bash'",
					"level":       5,
					"severity":    "HIGH",
					"description": "Second analytic",
				},
			},
			"pageInfo": map[string]any{
				"next":  nil,
				"total": 2,
			},
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "listAnalytics")
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := analytics.NewService(transport)

	result, _, err := service.ListAnalytics(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "analytic-1", result[0].UUID)
	assert.Equal(t, "Analytic 1", result[0].Name)
	assert.Equal(t, "analytic-2", result[1].UUID)
	assert.Equal(t, "Analytic 2", result[1].Name)
}

func TestAnalyticsService_CreateAnalytic_ValidationErrors(t *testing.T) {
	transport, err := client.NewTransport("test-client", "test-secret")
	require.NoError(t, err)

	service := analytics.NewService(transport)

	tests := []struct {
		name    string
		req     *analytics.CreateAnalyticRequest
		wantErr string
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: "request cannot be nil",
		},
		{
			name: "missing name",
			req: &analytics.CreateAnalyticRequest{
				InputType:   "unified_log",
				Description: "test",
				Filter:      "test",
			},
			wantErr: "name is required",
		},
		{
			name: "missing inputType",
			req: &analytics.CreateAnalyticRequest{
				Name:        "test",
				Description: "test",
				Filter:      "test",
			},
			wantErr: "inputType is required",
		},
		{
			name: "missing filter",
			req: &analytics.CreateAnalyticRequest{
				Name:        "test",
				InputType:   "unified_log",
				Description: "test",
			},
			wantErr: "filter is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := service.CreateAnalytic(context.Background(), tt.req)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}
