package mocks

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jarcoal/httpmock"
)

// PreventListMock provides mock responses for the CustomPreventList service GraphQL operations.
// All operations POST to the /graphql endpoint and are distinguished by operation name
// in the request body.
type PreventListMock struct {
	baseURL string
}

// NewPreventListMock creates a new PreventListMock instance
func NewPreventListMock(baseURL string) *PreventListMock {
	return &PreventListMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for prevent list operations
func (m *PreventListMock) RegisterMocks() {
	m.RegisterCreatePreventListMock()
	m.RegisterGetPreventListMock()
	m.RegisterUpdatePreventListMock()
	m.RegisterDeletePreventListMock()
	m.RegisterListPreventListsMock()
	m.RegisterListPreventListNamesMock()
}

// RegisterErrorMocks registers error response mocks
func (m *PreventListMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreatePreventListMock registers a success mock for createPreventList
func (m *PreventListMock) RegisterCreatePreventListMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("createPreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("create_prevent_list_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterGetPreventListMock registers a success mock for getPreventList
func (m *PreventListMock) RegisterGetPreventListMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getPreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("get_prevent_list_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUpdatePreventListMock registers a success mock for updatePreventList
func (m *PreventListMock) RegisterUpdatePreventListMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("updatePreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("update_prevent_list_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterDeletePreventListMock registers a success mock for deletePreventList
func (m *PreventListMock) RegisterDeletePreventListMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("deletePreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("delete_prevent_list_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListPreventListsMock registers a success mock for listPreventLists
func (m *PreventListMock) RegisterListPreventListsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listPreventLists"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_prevent_lists_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListPreventListNamesMock registers a success mock for listPreventListNames
func (m *PreventListMock) RegisterListPreventListNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listPreventListNames"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_prevent_list_names_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *PreventListMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getPreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(401, m.loadMockData("error_unauthorized.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterNotFoundErrorMock registers a not-found error mock
func (m *PreventListMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getPreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("error_not_found.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// loadMockData loads mock JSON data from a file relative to this source file
func (m *PreventListMock) loadMockData(filename string) []byte {
	_, currentFile, _, _ := runtime.Caller(0)
	mockDir := filepath.Dir(currentFile)
	mockFile := filepath.Join(mockDir, filename)

	data, err := os.ReadFile(mockFile)
	if err != nil {
		panic("Failed to load mock data: " + err.Error())
	}

	return data
}
