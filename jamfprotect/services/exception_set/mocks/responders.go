package mocks

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jarcoal/httpmock"
)

// ExceptionSetMock provides mock responses for the ExceptionSet service GraphQL operations.
// All operations POST to the /app endpoint and are distinguished by operation name in the request body.
type ExceptionSetMock struct {
	baseURL string
}

// NewExceptionSetMock creates a new ExceptionSetMock instance
func NewExceptionSetMock(baseURL string) *ExceptionSetMock {
	return &ExceptionSetMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for exception set operations
func (m *ExceptionSetMock) RegisterMocks() {
	m.RegisterCreateExceptionSetMock()
	m.RegisterGetExceptionSetMock()
	m.RegisterUpdateExceptionSetMock()
	m.RegisterDeleteExceptionSetMock()
	m.RegisterListExceptionSetsMock()
	m.RegisterListExceptionSetNamesMock()
}

// RegisterErrorMocks registers error response mocks
func (m *ExceptionSetMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreateExceptionSetMock registers a success mock for createExceptionSet
func (m *ExceptionSetMock) RegisterCreateExceptionSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("createExceptionSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("create_exception_set_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterGetExceptionSetMock registers a success mock for getExceptionSet
func (m *ExceptionSetMock) RegisterGetExceptionSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getExceptionSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("get_exception_set_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUpdateExceptionSetMock registers a success mock for updateExceptionSet
func (m *ExceptionSetMock) RegisterUpdateExceptionSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("updateExceptionSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("update_exception_set_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterDeleteExceptionSetMock registers a success mock for deleteExceptionSet
func (m *ExceptionSetMock) RegisterDeleteExceptionSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("deleteExceptionSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("delete_exception_set_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListExceptionSetsMock registers a success mock for listExceptionSets
func (m *ExceptionSetMock) RegisterListExceptionSetsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listExceptionSets"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_exception_sets_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListExceptionSetNamesMock registers a success mock for listExceptionSetNames
func (m *ExceptionSetMock) RegisterListExceptionSetNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listExceptionSetNames"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_exception_set_names_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *ExceptionSetMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getExceptionSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(401, m.loadMockData("error_unauthorized.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterNotFoundErrorMock registers a not-found error mock
func (m *ExceptionSetMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getExceptionSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("error_not_found.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// loadMockData loads mock JSON data from a file relative to this source file
func (m *ExceptionSetMock) loadMockData(filename string) []byte {
	_, currentFile, _, _ := runtime.Caller(0)
	mockDir := filepath.Dir(currentFile)
	mockFile := filepath.Join(mockDir, filename)

	data, err := os.ReadFile(mockFile)
	if err != nil {
		panic("Failed to load mock data: " + err.Error())
	}

	return data
}
