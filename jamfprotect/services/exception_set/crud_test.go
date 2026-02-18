package exceptionset_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	exceptionset "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exception_set"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exception_set/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBaseURL = "https://test.jamfprotect.example.com"

const testUUID = "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee"

func setupMockClient(t *testing.T) (*exceptionset.Service, string) {
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

	return exceptionset.NewService(transport), testBaseURL
}

func TestExceptionSetService_CreateExceptionSet(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewExceptionSetMock(baseURL)
	mockHandler.RegisterCreateExceptionSetMock()

	req := &exceptionset.CreateExceptionSetRequest{
		Name:        "Test Exception Set",
		Description: "A test exception set",
	}

	result, _, err := service.CreateExceptionSet(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Exception Set", result.Name)
}

func TestExceptionSetService_GetExceptionSet(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewExceptionSetMock(baseURL)
	mockHandler.RegisterGetExceptionSetMock()

	result, _, err := service.GetExceptionSet(context.Background(), testUUID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Exception Set", result.Name)
}

func TestExceptionSetService_UpdateExceptionSet(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewExceptionSetMock(baseURL)
	mockHandler.RegisterUpdateExceptionSetMock()

	req := &exceptionset.UpdateExceptionSetRequest{
		Name:        "Updated Exception Set",
		Description: "An updated exception set",
	}

	result, _, err := service.UpdateExceptionSet(context.Background(), testUUID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Updated Exception Set", result.Name)
}

func TestExceptionSetService_DeleteExceptionSet(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewExceptionSetMock(baseURL)
	mockHandler.RegisterDeleteExceptionSetMock()

	_, err := service.DeleteExceptionSet(context.Background(), testUUID)

	require.NoError(t, err)
}

func TestExceptionSetService_ListExceptionSets(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewExceptionSetMock(baseURL)
	mockHandler.RegisterListExceptionSetsMock()

	result, _, err := service.ListExceptionSets(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testUUID, result[0].UUID)
}

func TestExceptionSetService_ListExceptionSetNames(t *testing.T) {
	service, baseURL := setupMockClient(t)
	mockHandler := mocks.NewExceptionSetMock(baseURL)
	mockHandler.RegisterListExceptionSetNamesMock()

	result, _, err := service.ListExceptionSetNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Exception Set", result[0])
}

func TestExceptionSetService_ValidationErrors(t *testing.T) {
	service, _ := setupMockClient(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateExceptionSet nil request",
			fn: func() error {
				_, _, err := service.CreateExceptionSet(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateExceptionSet missing name",
			fn: func() error {
				_, _, err := service.CreateExceptionSet(context.Background(), &exceptionset.CreateExceptionSetRequest{})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetExceptionSet empty uuid",
			fn: func() error {
				_, _, err := service.GetExceptionSet(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "DeleteExceptionSet empty uuid",
			fn: func() error {
				_, err := service.DeleteExceptionSet(context.Background(), "")
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
