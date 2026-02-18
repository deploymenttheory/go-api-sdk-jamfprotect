package telemetry_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/telemetry"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/telemetry/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBaseURL = "https://test.jamfprotect.example.com"

func setupMockClient(t *testing.T) (*telemetry.Service, string) {
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

	return telemetry.NewService(transport), testBaseURL
}

func TestTelemetryService_CreateTelemetryV2(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewTelemetryMock(baseURL)
	mockHandler.RegisterCreateTelemetryV2Mock()

	req := &telemetry.CreateTelemetryV2Request{
		Name:     "Test Telemetry V2",
		Description: "A test telemetry v2",
		LogFiles: []string{},
	}

	result, _, err := service.CreateTelemetryV2(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Telemetry V2", result.Name)
}

func TestTelemetryService_GetTelemetryV2(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewTelemetryMock(baseURL)
	mockHandler.RegisterGetTelemetryV2Mock()

	result, _, err := service.GetTelemetryV2(context.Background(), "test-id-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Telemetry V2", result.Name)
}

func TestTelemetryService_UpdateTelemetryV2(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewTelemetryMock(baseURL)
	mockHandler.RegisterUpdateTelemetryV2Mock()

	req := &telemetry.UpdateTelemetryV2Request{
		Name:     "Updated Telemetry V2",
		Description: "An updated telemetry v2",
		LogFiles: []string{},
	}

	result, _, err := service.UpdateTelemetryV2(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Updated Telemetry V2", result.Name)
}

func TestTelemetryService_DeleteTelemetryV2(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewTelemetryMock(baseURL)
	mockHandler.RegisterDeleteTelemetryV2Mock()

	_, err := service.DeleteTelemetryV2(context.Background(), "test-id-1234")

	require.NoError(t, err)
}

func TestTelemetryService_ListTelemetriesV2(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewTelemetryMock(baseURL)
	mockHandler.RegisterListTelemetriesV2Mock()

	result, _, err := service.ListTelemetriesV2(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id-1234", result[0].ID)
	assert.Equal(t, "Test Telemetry V2", result[0].Name)
}

func TestTelemetryService_ListTelemetriesCombined(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewTelemetryMock(baseURL)
	mockHandler.RegisterListTelemetriesCombinedMock()

	result, _, err := service.ListTelemetriesCombined(context.Background(), false)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result.Telemetries, 1)
	assert.Equal(t, "tel-id-1", result.Telemetries[0].ID)
	assert.Equal(t, "Test Telemetry V1", result.Telemetries[0].Name)
	assert.Len(t, result.TelemetriesV2, 1)
	assert.Equal(t, "telv2-id-1", result.TelemetriesV2[0].ID)
	assert.Equal(t, "Test Telemetry V2", result.TelemetriesV2[0].Name)
}

func TestTelemetryService_ValidationErrors(t *testing.T) {
	service, _ := setupMockClient(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateTelemetryV2 nil request",
			fn: func() error {
				_, _, err := service.CreateTelemetryV2(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateTelemetryV2 missing name",
			fn: func() error {
				_, _, err := service.CreateTelemetryV2(context.Background(), &telemetry.CreateTelemetryV2Request{
					LogFiles: []string{},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateTelemetryV2 nil logFiles",
			fn: func() error {
				_, _, err := service.CreateTelemetryV2(context.Background(), &telemetry.CreateTelemetryV2Request{
					Name: "test",
				})
				return err
			},
			wantErr: "logFiles is required",
		},
		{
			name: "GetTelemetryV2 empty id",
			fn: func() error {
				_, _, err := service.GetTelemetryV2(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateTelemetryV2 empty id",
			fn: func() error {
				_, _, err := service.UpdateTelemetryV2(context.Background(), "", &telemetry.UpdateTelemetryV2Request{
					Name:     "test",
					LogFiles: []string{},
				})
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateTelemetryV2 nil request",
			fn: func() error {
				_, _, err := service.UpdateTelemetryV2(context.Background(), "test-id", nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "DeleteTelemetryV2 empty id",
			fn: func() error {
				_, err := service.DeleteTelemetryV2(context.Background(), "")
				return err
			},
			wantErr: "id is required",
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
