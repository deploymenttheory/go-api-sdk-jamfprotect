package mocks

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jarcoal/httpmock"
)

// AnalyticSetMock provides mock responses for the AnalyticSet service GraphQL operations.
// Most operations POST to /app; updateAnalyticSet posts to /graphql.
type AnalyticSetMock struct {
	baseURL string
}

// NewAnalyticSetMock creates a new AnalyticSetMock instance
func NewAnalyticSetMock(baseURL string) *AnalyticSetMock {
	return &AnalyticSetMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for analytic set operations
func (m *AnalyticSetMock) RegisterMocks() {
	m.RegisterCreateAnalyticSetMock()
	m.RegisterGetAnalyticSetMock()
	m.RegisterUpdateAnalyticSetMock()
	m.RegisterDeleteAnalyticSetMock()
	m.RegisterListAnalyticSetsMock()
}

// RegisterErrorMocks registers error response mocks
func (m *AnalyticSetMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreateAnalyticSetMock registers a success mock for createAnalyticSet
func (m *AnalyticSetMock) RegisterCreateAnalyticSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("createAnalyticSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("create_analytic_set_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterGetAnalyticSetMock registers a success mock for getAnalyticSet
func (m *AnalyticSetMock) RegisterGetAnalyticSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getAnalyticSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("get_analytic_set_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUpdateAnalyticSetMock registers a success mock for updateAnalyticSet
func (m *AnalyticSetMock) RegisterUpdateAnalyticSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("updateAnalyticSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("update_analytic_set_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterDeleteAnalyticSetMock registers a success mock for deleteAnalyticSet
func (m *AnalyticSetMock) RegisterDeleteAnalyticSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("deleteAnalyticSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("delete_analytic_set_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListAnalyticSetsMock registers a success mock for listAnalyticSets
func (m *AnalyticSetMock) RegisterListAnalyticSetsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listAnalyticSets"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_analytic_sets_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *AnalyticSetMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getAnalyticSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(401, m.loadMockData("error_unauthorized.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterNotFoundErrorMock registers a not-found error mock
func (m *AnalyticSetMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getAnalyticSet"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("error_not_found.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// loadMockData loads mock JSON data from a file relative to this source file
func (m *AnalyticSetMock) loadMockData(filename string) []byte {
	_, currentFile, _, _ := runtime.Caller(0)
	mockDir := filepath.Dir(currentFile)
	mockFile := filepath.Join(mockDir, filename)

	data, err := os.ReadFile(mockFile)
	if err != nil {
		panic("Failed to load mock data: " + err.Error())
	}

	return data
}
