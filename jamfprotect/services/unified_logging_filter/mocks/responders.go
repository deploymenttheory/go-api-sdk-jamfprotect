package mocks

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jarcoal/httpmock"
)

// UnifiedLoggingFilterMock provides mock responses for the UnifiedLoggingFilter service GraphQL operations.
// All operations POST to the /graphql endpoint and are distinguished by operation name
// in the request body.
type UnifiedLoggingFilterMock struct {
	baseURL string
}

// NewUnifiedLoggingFilterMock creates a new UnifiedLoggingFilterMock instance
func NewUnifiedLoggingFilterMock(baseURL string) *UnifiedLoggingFilterMock {
	return &UnifiedLoggingFilterMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for unified logging filter operations
func (m *UnifiedLoggingFilterMock) RegisterMocks() {
	m.RegisterCreateUnifiedLoggingFilterMock()
	m.RegisterGetUnifiedLoggingFilterMock()
	m.RegisterUpdateUnifiedLoggingFilterMock()
	m.RegisterDeleteUnifiedLoggingFilterMock()
	m.RegisterListUnifiedLoggingFiltersMock()
	m.RegisterListUnifiedLoggingFilterNamesMock()
}

// RegisterErrorMocks registers error response mocks
func (m *UnifiedLoggingFilterMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreateUnifiedLoggingFilterMock registers a success mock for createUnifiedLoggingFilter
func (m *UnifiedLoggingFilterMock) RegisterCreateUnifiedLoggingFilterMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("createUnifiedLoggingFilter"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("create_unified_logging_filter_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterGetUnifiedLoggingFilterMock registers a success mock for getUnifiedLoggingFilter
func (m *UnifiedLoggingFilterMock) RegisterGetUnifiedLoggingFilterMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getUnifiedLoggingFilter"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("get_unified_logging_filter_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUpdateUnifiedLoggingFilterMock registers a success mock for updateUnifiedLoggingFilter
func (m *UnifiedLoggingFilterMock) RegisterUpdateUnifiedLoggingFilterMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("updateUnifiedLoggingFilter"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("update_unified_logging_filter_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterDeleteUnifiedLoggingFilterMock registers a success mock for deleteUnifiedLoggingFilter
func (m *UnifiedLoggingFilterMock) RegisterDeleteUnifiedLoggingFilterMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("deleteUnifiedLoggingFilter"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("delete_unified_logging_filter_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListUnifiedLoggingFiltersMock registers a success mock for listUnifiedLoggingFilters
func (m *UnifiedLoggingFilterMock) RegisterListUnifiedLoggingFiltersMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listUnifiedLoggingFilters"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_unified_logging_filters_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListUnifiedLoggingFilterNamesMock registers a success mock for listUnifiedLoggingFilterNames
func (m *UnifiedLoggingFilterMock) RegisterListUnifiedLoggingFilterNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listUnifiedLoggingFilterNames"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_unified_logging_filter_names_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *UnifiedLoggingFilterMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getUnifiedLoggingFilter"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(401, m.loadMockData("error_unauthorized.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterNotFoundErrorMock registers a not-found error mock
func (m *UnifiedLoggingFilterMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getUnifiedLoggingFilter"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("error_not_found.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// loadMockData loads mock JSON data from a file relative to this source file
func (m *UnifiedLoggingFilterMock) loadMockData(filename string) []byte {
	_, currentFile, _, _ := runtime.Caller(0)
	mockDir := filepath.Dir(currentFile)
	mockFile := filepath.Join(mockDir, filename)

	data, err := os.ReadFile(mockFile)
	if err != nil {
		panic("Failed to load mock data: " + err.Error())
	}

	return data
}
