package custompreventlist_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/custom_prevent_list"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/custom_prevent_list/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBaseURL = "https://test.jamfprotect.example.com"

func setupMockClient(t *testing.T) (*custompreventlist.Service, string) {
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

	return custompreventlist.NewService(transport), testBaseURL
}

func TestPreventListService_CreatePreventList(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewPreventListMock(baseURL)
	mockHandler.RegisterCreatePreventListMock()

	req := &custompreventlist.CreatePreventListRequest{
		Name:        "Test Prevent List",
		Description: "A test prevent list",
		Type:        "FILEHASH",
	}

	result, _, err := service.CreatePreventList(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Prevent List", result.Name)
}

func TestPreventListService_GetPreventList(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewPreventListMock(baseURL)
	mockHandler.RegisterGetPreventListMock()

	result, _, err := service.GetPreventList(context.Background(), "test-id-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Prevent List", result.Name)
}

func TestPreventListService_UpdatePreventList(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewPreventListMock(baseURL)
	mockHandler.RegisterUpdatePreventListMock()

	req := &custompreventlist.UpdatePreventListRequest{
		Name:        "Updated Prevent List",
		Description: "An updated prevent list",
		Type:        "FILEHASH",
	}

	result, _, err := service.UpdatePreventList(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Updated Prevent List", result.Name)
}

func TestPreventListService_DeletePreventList(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewPreventListMock(baseURL)
	mockHandler.RegisterDeletePreventListMock()

	_, err := service.DeletePreventList(context.Background(), "test-id-1234")

	require.NoError(t, err)
}

func TestPreventListService_ListPreventLists(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewPreventListMock(baseURL)
	mockHandler.RegisterListPreventListsMock()

	result, _, err := service.ListPreventLists(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id-1234", result[0].ID)
	assert.Equal(t, "Test Prevent List", result[0].Name)
}

func TestPreventListService_ListPreventListNames(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewPreventListMock(baseURL)
	mockHandler.RegisterListPreventListNamesMock()

	result, _, err := service.ListPreventListNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Prevent List", result[0])
}

func TestPreventListService_ValidationErrors(t *testing.T) {
	service, _ := setupMockClient(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreatePreventList nil request",
			fn: func() error {
				_, _, err := service.CreatePreventList(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreatePreventList missing name",
			fn: func() error {
				_, _, err := service.CreatePreventList(context.Background(), &custompreventlist.CreatePreventListRequest{
					Type: "FILEHASH",
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetPreventList empty id",
			fn: func() error {
				_, _, err := service.GetPreventList(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "DeletePreventList empty id",
			fn: func() error {
				_, err := service.DeletePreventList(context.Background(), "")
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
