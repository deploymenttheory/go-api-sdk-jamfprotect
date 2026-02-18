package analyticset_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	analyticset "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic_set"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic_set/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBaseURL = "https://test.jamfprotect.example.com"

func setupMockClient(t *testing.T) (*analyticset.Service, string) {
	t.Helper()

	httpClient := &http.Client{}
	httpmock.ActivateNonDefault(httpClient)
	t.Cleanup(func() {
		httpmock.DeactivateAndReset()
	})

	httpmock.RegisterResponder("POST", testBaseURL+"/token",
		httpmock.NewJsonResponderOrPanic(200, map[string]any{
			"access_token": "mock-token",
			"expires_in":   3600,
			"token_type":   "Bearer",
		}),
	)

	transport, err := client.NewTransport("test-client", "test-secret",
		client.WithBaseURL(testBaseURL),
		client.WithTransport(httpClient.Transport),
	)
	require.NoError(t, err)

	return analyticset.NewService(transport), testBaseURL
}

func TestAnalyticSetService_CreateAnalyticSet(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticSetMock(baseURL)
	mockHandler.RegisterCreateAnalyticSetMock()

	req := &analyticset.CreateAnalyticSetRequest{
		Name:        "Test Analytic Set",
		Description: "A test analytic set",
		Analytics:   []string{"analytic-uuid-1"},
	}

	result, _, err := service.CreateAnalyticSet(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-uuid-1234", result.UUID)
	assert.Equal(t, "Test Analytic Set", result.Name)
}

func TestAnalyticSetService_GetAnalyticSet(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticSetMock(baseURL)
	mockHandler.RegisterGetAnalyticSetMock()

	result, _, err := service.GetAnalyticSet(context.Background(), "test-uuid-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-uuid-1234", result.UUID)
	assert.Equal(t, "Test Analytic Set", result.Name)
}

func TestAnalyticSetService_UpdateAnalyticSet(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticSetMock(baseURL)
	mockHandler.RegisterUpdateAnalyticSetMock()

	req := &analyticset.UpdateAnalyticSetRequest{
		Name:        "Updated Analytic Set",
		Description: "An updated analytic set",
		Analytics:   []string{"analytic-uuid-1"},
	}

	result, _, err := service.UpdateAnalyticSet(context.Background(), "test-uuid-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-uuid-1234", result.UUID)
	assert.Equal(t, "Updated Analytic Set", result.Name)
}

func TestAnalyticSetService_DeleteAnalyticSet(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticSetMock(baseURL)
	mockHandler.RegisterDeleteAnalyticSetMock()

	_, err := service.DeleteAnalyticSet(context.Background(), "test-uuid-1234")

	require.NoError(t, err)
}

func TestAnalyticSetService_ListAnalyticSets(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticSetMock(baseURL)
	mockHandler.RegisterListAnalyticSetsMock()

	result, _, err := service.ListAnalyticSets(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-uuid-1234", result[0].UUID)
}

func TestAnalyticSetService_ValidationErrors(t *testing.T) {
	service, _ := setupMockClient(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateAnalyticSet nil request",
			fn: func() error {
				_, _, err := service.CreateAnalyticSet(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateAnalyticSet missing name",
			fn: func() error {
				_, _, err := service.CreateAnalyticSet(context.Background(), &analyticset.CreateAnalyticSetRequest{
					Analytics: []string{"uuid-1"},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetAnalyticSet empty uuid",
			fn: func() error {
				_, _, err := service.GetAnalyticSet(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "DeleteAnalyticSet empty uuid",
			fn: func() error {
				_, err := service.DeleteAnalyticSet(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}
