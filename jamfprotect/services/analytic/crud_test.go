package analytic_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBaseURL = "https://test.jamfprotect.example.com"

func setupMockClient(t *testing.T) (*analytic.Service, string) {
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

	return analytic.NewService(transport), testBaseURL
}

func TestAnalyticService_CreateAnalytic(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterCreateAnalyticMock()

	req := &analytic.CreateAnalyticRequest{
		Name:        "Test Analytic",
		InputType:   "GPFSEvent",
		Filter:      "process.name = 'test'",
		Description: "A test analytic",
		Level:       3,
		Severity:    "Medium",
	}

	result, _, err := service.CreateAnalytic(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-uuid-1234", result.UUID)
	assert.Equal(t, "Test Analytic", result.Name)
}

func TestAnalyticService_GetAnalytic(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterGetAnalyticMock()

	result, _, err := service.GetAnalytic(context.Background(), "test-uuid-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-uuid-1234", result.UUID)
	assert.Equal(t, "Test Analytic", result.Name)
}

func TestAnalyticService_UpdateAnalytic(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterUpdateAnalyticMock()

	req := &analytic.UpdateAnalyticRequest{
		Name:        "Updated Analytic",
		InputType:   "GPFSEvent",
		Filter:      "process.name = 'updated'",
		Description: "An updated test analytic",
		Level:       5,
	}

	result, _, err := service.UpdateAnalytic(context.Background(), "test-uuid-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-uuid-1234", result.UUID)
	assert.Equal(t, "Updated Analytic", result.Name)
}

func TestAnalyticService_DeleteAnalytic(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterDeleteAnalyticMock()

	_, err := service.DeleteAnalytic(context.Background(), "test-uuid-1234")

	require.NoError(t, err)
}

func TestAnalyticService_ListAnalytics(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterListAnalyticsMock()

	result, _, err := service.ListAnalytics(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-uuid-1234", result[0].UUID)
	assert.Equal(t, "Test Analytic", result[0].Name)
}

func TestAnalyticService_ListAnalyticsLite(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterListAnalyticsLiteMock()

	result, _, err := service.ListAnalyticsLite(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-uuid-1234", result[0].UUID)
	assert.Equal(t, "Test Analytic", result[0].Name)
}

func TestAnalyticService_ListAnalyticsNames(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterListAnalyticsNamesMock()

	result, _, err := service.ListAnalyticsNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Analytic", result[0])
}

func TestAnalyticService_ListAnalyticsCategories(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterListAnalyticsCategoriesMock()

	result, _, err := service.ListAnalyticsCategories(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Security", result[0].Value)
	assert.Equal(t, 5, result[0].Count)
}

func TestAnalyticService_ListAnalyticsTags(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterListAnalyticsTagsMock()

	result, _, err := service.ListAnalyticsTags(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "endpoint", result[0].Value)
	assert.Equal(t, 10, result[0].Count)
}

func TestAnalyticService_ListAnalyticsFilterOptions(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewAnalyticMock(baseURL)
	mockHandler.RegisterListAnalyticsFilterOptionsMock()

	result, _, err := service.ListAnalyticsFilterOptions(context.Background())

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result.Tags, 1)
	assert.Len(t, result.Categories, 1)
	assert.Equal(t, "endpoint", result.Tags[0].Value)
	assert.Equal(t, "Security", result.Categories[0].Value)
}

func TestAnalyticService_ValidationErrors(t *testing.T) {
	service, _ := setupMockClient(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateAnalytic nil request",
			fn: func() error {
				_, _, err := service.CreateAnalytic(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateAnalytic missing name",
			fn: func() error {
				_, _, err := service.CreateAnalytic(context.Background(), &analytic.CreateAnalyticRequest{
					InputType: "GPFSEvent",
					Filter:    "test",
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateAnalytic missing inputType",
			fn: func() error {
				_, _, err := service.CreateAnalytic(context.Background(), &analytic.CreateAnalyticRequest{
					Name:   "test",
					Filter: "test",
				})
				return err
			},
			wantErr: "inputType is required",
		},
		{
			name: "CreateAnalytic missing filter",
			fn: func() error {
				_, _, err := service.CreateAnalytic(context.Background(), &analytic.CreateAnalyticRequest{
					Name:      "test",
					InputType: "GPFSEvent",
				})
				return err
			},
			wantErr: "filter is required",
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
