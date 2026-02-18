package unifiedloggingfilter_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/unified_logging_filter"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/unified_logging_filter/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBaseURL = "https://test.jamfprotect.example.com"

const testUUID = "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee"

func setupMockClient(t *testing.T) (*unifiedloggingfilter.Service, string) {
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

	return unifiedloggingfilter.NewService(transport), testBaseURL
}

func TestUnifiedLoggingFilterService_CreateUnifiedLoggingFilter(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewUnifiedLoggingFilterMock(baseURL)
	mockHandler.RegisterCreateUnifiedLoggingFilterMock()

	req := &unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{
		Name:        "Test Unified Logging Filter",
		Description: "A test unified logging filter",
		Filter:      "process.name == \"test\"",
		Enabled:     true,
	}

	result, _, err := service.CreateUnifiedLoggingFilter(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Unified Logging Filter", result.Name)
}

func TestUnifiedLoggingFilterService_GetUnifiedLoggingFilter(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewUnifiedLoggingFilterMock(baseURL)
	mockHandler.RegisterGetUnifiedLoggingFilterMock()

	result, _, err := service.GetUnifiedLoggingFilter(context.Background(), testUUID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Unified Logging Filter", result.Name)
}

func TestUnifiedLoggingFilterService_UpdateUnifiedLoggingFilter(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewUnifiedLoggingFilterMock(baseURL)
	mockHandler.RegisterUpdateUnifiedLoggingFilterMock()

	req := &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
		Name:        "Updated Unified Logging Filter",
		Description: "An updated unified logging filter",
		Filter:      "process.name == \"updated\"",
		Enabled:     true,
	}

	result, _, err := service.UpdateUnifiedLoggingFilter(context.Background(), testUUID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Updated Unified Logging Filter", result.Name)
}

func TestUnifiedLoggingFilterService_DeleteUnifiedLoggingFilter(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewUnifiedLoggingFilterMock(baseURL)
	mockHandler.RegisterDeleteUnifiedLoggingFilterMock()

	_, err := service.DeleteUnifiedLoggingFilter(context.Background(), testUUID)

	require.NoError(t, err)
}

func TestUnifiedLoggingFilterService_ListUnifiedLoggingFilters(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewUnifiedLoggingFilterMock(baseURL)
	mockHandler.RegisterListUnifiedLoggingFiltersMock()

	result, _, err := service.ListUnifiedLoggingFilters(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testUUID, result[0].UUID)
	assert.Equal(t, "Test Unified Logging Filter", result[0].Name)
}

func TestUnifiedLoggingFilterService_ListUnifiedLoggingFilterNames(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewUnifiedLoggingFilterMock(baseURL)
	mockHandler.RegisterListUnifiedLoggingFilterNamesMock()

	result, _, err := service.ListUnifiedLoggingFilterNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Unified Logging Filter", result[0])
}

func TestUnifiedLoggingFilterService_ValidationErrors(t *testing.T) {
	service, _ := setupMockClient(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateUnifiedLoggingFilter nil request",
			fn: func() error {
				_, _, err := service.CreateUnifiedLoggingFilter(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateUnifiedLoggingFilter missing name",
			fn: func() error {
				_, _, err := service.CreateUnifiedLoggingFilter(context.Background(), &unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{
					Filter: "test",
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateUnifiedLoggingFilter missing filter",
			fn: func() error {
				_, _, err := service.CreateUnifiedLoggingFilter(context.Background(), &unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "filter is required",
		},
		{
			name: "GetUnifiedLoggingFilter empty uuid",
			fn: func() error {
				_, _, err := service.GetUnifiedLoggingFilter(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "UpdateUnifiedLoggingFilter empty uuid",
			fn: func() error {
				_, _, err := service.UpdateUnifiedLoggingFilter(context.Background(), "", &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
					Name:   "test",
					Filter: "test",
				})
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "UpdateUnifiedLoggingFilter nil request",
			fn: func() error {
				_, _, err := service.UpdateUnifiedLoggingFilter(context.Background(), testUUID, nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "DeleteUnifiedLoggingFilter empty uuid",
			fn: func() error {
				_, err := service.DeleteUnifiedLoggingFilter(context.Background(), "")
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
