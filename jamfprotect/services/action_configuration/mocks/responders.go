package mocks

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jarcoal/httpmock"
)

// ActionConfigMock provides mock responses for the ActionConfiguration service GraphQL operations.
// All operations POST to the /app endpoint and are distinguished by operation name in the request body.
type ActionConfigMock struct {
	baseURL string
}

// NewActionConfigMock creates a new ActionConfigMock instance
func NewActionConfigMock(baseURL string) *ActionConfigMock {
	return &ActionConfigMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for action config operations
func (m *ActionConfigMock) RegisterMocks() {
	m.RegisterCreateActionConfigMock()
	m.RegisterGetActionConfigMock()
	m.RegisterUpdateActionConfigMock()
	m.RegisterDeleteActionConfigMock()
	m.RegisterListActionConfigsMock()
	m.RegisterListActionConfigNamesMock()
}

// RegisterErrorMocks registers error response mocks
func (m *ActionConfigMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreateActionConfigMock registers a success mock for createActionConfigs
func (m *ActionConfigMock) RegisterCreateActionConfigMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("createActionConfigs"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("create_action_config_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterGetActionConfigMock registers a success mock for getActionConfigs
func (m *ActionConfigMock) RegisterGetActionConfigMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getActionConfigs"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("get_action_config_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUpdateActionConfigMock registers a success mock for updateActionConfigs
func (m *ActionConfigMock) RegisterUpdateActionConfigMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("updateActionConfigs"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("update_action_config_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterDeleteActionConfigMock registers a success mock for deleteActionConfigs
func (m *ActionConfigMock) RegisterDeleteActionConfigMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("deleteActionConfigs"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("delete_action_config_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListActionConfigsMock registers a success mock for listActionConfigs
func (m *ActionConfigMock) RegisterListActionConfigsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listActionConfigs"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_action_configs_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListActionConfigNamesMock registers a success mock for listActionConfigNames
func (m *ActionConfigMock) RegisterListActionConfigNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listActionConfigNames"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_action_config_names_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *ActionConfigMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getActionConfigs"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(401, m.loadMockData("error_unauthorized.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterNotFoundErrorMock registers a not-found error mock
func (m *ActionConfigMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getActionConfigs"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("error_not_found.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// loadMockData loads mock JSON data from a file relative to this source file
func (m *ActionConfigMock) loadMockData(filename string) []byte {
	_, currentFile, _, _ := runtime.Caller(0)
	mockDir := filepath.Dir(currentFile)
	mockFile := filepath.Join(mockDir, filename)

	data, err := os.ReadFile(mockFile)
	if err != nil {
		panic("Failed to load mock data: " + err.Error())
	}

	return data
}
