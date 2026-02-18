package actionconfiguration_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	actionconfiguration "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/action_configuration"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/action_configuration/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBaseURL = "https://test.jamfprotect.example.com"

func setupMockClient(t *testing.T) (*actionconfiguration.Service, string) {
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

	return actionconfiguration.NewService(transport), testBaseURL
}

func TestActionConfigService_CreateActionConfig(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewActionConfigMock(baseURL)
	mockHandler.RegisterCreateActionConfigMock()

	req := &actionconfiguration.CreateActionConfigRequest{
		Name:        "Test Action Config",
		Description: "A test action configuration",
		AlertConfig: map[string]any{
			"type": "alert",
		},
	}

	result, _, err := service.CreateActionConfig(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Action Config", result.Name)
}

func TestActionConfigService_GetActionConfig(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewActionConfigMock(baseURL)
	mockHandler.RegisterGetActionConfigMock()

	result, _, err := service.GetActionConfig(context.Background(), "test-id-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Action Config", result.Name)
}

func TestActionConfigService_UpdateActionConfig(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewActionConfigMock(baseURL)
	mockHandler.RegisterUpdateActionConfigMock()

	req := &actionconfiguration.UpdateActionConfigRequest{
		Name:        "Updated Action Config",
		Description: "An updated action configuration",
		AlertConfig: map[string]any{
			"type": "alert",
		},
	}

	result, _, err := service.UpdateActionConfig(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Updated Action Config", result.Name)
}

func TestActionConfigService_DeleteActionConfig(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewActionConfigMock(baseURL)
	mockHandler.RegisterDeleteActionConfigMock()

	_, err := service.DeleteActionConfig(context.Background(), "test-id-1234")

	require.NoError(t, err)
}

func TestActionConfigService_ListActionConfigs(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewActionConfigMock(baseURL)
	mockHandler.RegisterListActionConfigsMock()

	result, _, err := service.ListActionConfigs(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id-1234", result[0].ID)
	assert.Equal(t, "Test Action Config", result[0].Name)
}

func TestActionConfigService_ListActionConfigNames(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewActionConfigMock(baseURL)
	mockHandler.RegisterListActionConfigNamesMock()

	result, _, err := service.ListActionConfigNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Action Config", result[0])
}

func TestActionConfigService_ValidationErrors(t *testing.T) {
	service, _ := setupMockClient(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateActionConfig nil request",
			fn: func() error {
				_, _, err := service.CreateActionConfig(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateActionConfig missing name",
			fn: func() error {
				_, _, err := service.CreateActionConfig(context.Background(), &actionconfiguration.CreateActionConfigRequest{
					AlertConfig: map[string]any{"type": "alert"},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateActionConfig missing alertConfig",
			fn: func() error {
				_, _, err := service.CreateActionConfig(context.Background(), &actionconfiguration.CreateActionConfigRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "alertConfig is required",
		},
		{
			name: "GetActionConfig empty id",
			fn: func() error {
				_, _, err := service.GetActionConfig(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "DeleteActionConfig empty id",
			fn: func() error {
				_, err := service.DeleteActionConfig(context.Background(), "")
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
