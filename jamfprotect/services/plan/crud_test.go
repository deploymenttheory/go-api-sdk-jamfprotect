package plan_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plan"
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

func TestPlansService_CreatePlan(t *testing.T) {
	mockData := map[string]any{
		"createPlan": map[string]any{
			"id":          "plan-123",
			"name":        "Test Plan",
			"description": "Test Description",
			"logLevel":    "INFO",
			"autoUpdate":  true,
			"created":     "2024-01-01T00:00:00Z",
			"updated":     "2024-01-01T00:00:00Z",
			"commsConfig": map[string]any{
				"fqdn":     "test.example.com",
				"protocol": "HTTPS",
			},
			"infoSync": map[string]any{
				"attrs":                []string{"hostname"},
				"insightsSyncInterval": 3600,
			},
			"signaturesFeedConfig": map[string]any{
				"mode": "AUTO",
			},
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		// Verify request
		assert.Contains(t, query, "createPlan")
		assert.Equal(t, "Test Plan", variables["name"])
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := plan.NewService(transport)

	logLevel := "INFO"
	req := &plan.CreatePlanRequest{
		Name:          "Test Plan",
		Description:   "Test Description",
		LogLevel:      &logLevel,
		ActionConfigs: "action-123",
		AutoUpdate:    true,
		CommsConfig: plan.CommsConfigInput{
			FQDN:     "test.example.com",
			Protocol: "HTTPS",
		},
		InfoSync: plan.InfoSyncInput{
			Attrs:                []string{"hostname"},
			InsightsSyncInterval: 3600,
		},
		SignaturesFeedConfig: plan.SignaturesFeedConfigInput{
			Mode: "AUTO",
		},
	}

	result, _, err := service.CreatePlan(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "plan-123", result.ID)
	assert.Equal(t, "Test Plan", result.Name)
	assert.Equal(t, "Test Description", result.Description)
	assert.Equal(t, "INFO", result.LogLevel)
	assert.True(t, result.AutoUpdate)
}

func TestPlansService_GetPlan(t *testing.T) {
	mockData := map[string]any{
		"getPlan": map[string]any{
			"id":          "plan-123",
			"name":        "Test Plan",
			"description": "Test Description",
			"logLevel":    "DEBUG",
			"created":     "2024-01-01T00:00:00Z",
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "getPlan")
		assert.Equal(t, "plan-123", variables["id"])
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := plan.NewService(transport)

	result, _, err := service.GetPlan(context.Background(), "plan-123")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "plan-123", result.ID)
	assert.Equal(t, "Test Plan", result.Name)
}

func TestPlansService_UpdatePlan(t *testing.T) {
	mockData := map[string]any{
		"updatePlan": map[string]any{
			"id":          "plan-123",
			"name":        "Updated Plan",
			"description": "Updated Description",
			"logLevel":    "WARN",
			"updated":     "2024-01-02T00:00:00Z",
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "updatePlan")
		assert.Equal(t, "plan-123", variables["id"])
		assert.Equal(t, "Updated Plan", variables["name"])
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := plan.NewService(transport)

	logLevel := "WARN"
	req := &plan.UpdatePlanRequest{
		Name:          "Updated Plan",
		Description:   "Updated Description",
		LogLevel:      &logLevel,
		ActionConfigs: "action-123",
		AutoUpdate:    true,
		CommsConfig: plan.CommsConfigInput{
			FQDN:     "test.example.com",
			Protocol: "HTTPS",
		},
		InfoSync: plan.InfoSyncInput{
			Attrs:                []string{"hostname"},
			InsightsSyncInterval: 3600,
		},
		SignaturesFeedConfig: plan.SignaturesFeedConfigInput{
			Mode: "AUTO",
		},
	}

	result, _, err := service.UpdatePlan(context.Background(), "plan-123", req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "plan-123", result.ID)
	assert.Equal(t, "Updated Plan", result.Name)
}

func TestPlansService_DeletePlan(t *testing.T) {
	mockData := map[string]any{
		"deletePlan": map[string]any{
			"id": "plan-123",
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "deletePlan")
		assert.Equal(t, "plan-123", variables["id"])
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := plan.NewService(transport)

	_, err = service.DeletePlan(context.Background(), "plan-123")

	require.NoError(t, err)
}

func TestPlansService_ListPlans(t *testing.T) {
	mockData := map[string]any{
		"listPlans": map[string]any{
			"items": []map[string]any{
				{
					"id":          "plan-1",
					"name":        "Plan 1",
					"description": "First plan",
					"logLevel":    "INFO",
				},
				{
					"id":          "plan-2",
					"name":        "Plan 2",
					"description": "Second plan",
					"logLevel":    "DEBUG",
				},
			},
			"pageInfo": map[string]any{
				"next":  nil,
				"total": 2,
			},
		},
	}

	server := mockGraphQLServer(t, func(query string, variables map[string]any) (any, error) {
		assert.Contains(t, query, "listPlans")
		return mockData, nil
	})
	defer server.Close()

	transport, err := client.NewTransport("test-client", "test-secret", client.WithBaseURL(server.URL))
	require.NoError(t, err)

	service := plan.NewService(transport)

	result, _, err := service.ListPlans(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "plan-1", result[0].ID)
	assert.Equal(t, "Plan 1", result[0].Name)
	assert.Equal(t, "plan-2", result[1].ID)
	assert.Equal(t, "Plan 2", result[1].Name)
}

func TestPlansService_CreatePlan_ValidationErrors(t *testing.T) {
	transport, err := client.NewTransport("test-client", "test-secret")
	require.NoError(t, err)

	service := plan.NewService(transport)

	tests := []struct {
		name    string
		req     *plan.CreatePlanRequest
		wantErr string
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: "request cannot be nil",
		},
		{
			name: "missing name",
			req: &plan.CreatePlanRequest{
				Description:   "test",
				ActionConfigs: "action-123",
			},
			wantErr: "name is required",
		},
		{
			name: "missing action configs",
			req: &plan.CreatePlanRequest{
				Name:        "test",
				Description: "test",
			},
			wantErr: "actionConfigs is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := service.CreatePlan(context.Background(), tt.req)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}
