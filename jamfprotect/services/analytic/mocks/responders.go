package mocks

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jarcoal/httpmock"
)

// AnalyticMock provides mock responses for the Analytic service GraphQL operations.
// Mutations (create/update/delete) POST to the /app endpoint; queries POST to the /graphql endpoint.
// Operations are distinguished by operation name in the request body.
type AnalyticMock struct {
	baseURL string
}

// NewAnalyticMock creates a new AnalyticMock instance
func NewAnalyticMock(baseURL string) *AnalyticMock {
	return &AnalyticMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for analytic operations
func (m *AnalyticMock) RegisterMocks() {
	m.RegisterCreateAnalyticMock()
	m.RegisterGetAnalyticMock()
	m.RegisterUpdateAnalyticMock()
	m.RegisterDeleteAnalyticMock()
	m.RegisterListAnalyticsMock()
	m.RegisterListAnalyticsLiteMock()
	m.RegisterListAnalyticsNamesMock()
	m.RegisterListAnalyticsCategoriesMock()
	m.RegisterListAnalyticsTagsMock()
	m.RegisterListAnalyticsFilterOptionsMock()
}

// RegisterErrorMocks registers error response mocks
func (m *AnalyticMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreateAnalyticMock registers a success mock for createAnalytic
func (m *AnalyticMock) RegisterCreateAnalyticMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("createAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("create_analytic_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterGetAnalyticMock registers a success mock for getAnalytic
func (m *AnalyticMock) RegisterGetAnalyticMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("get_analytic_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUpdateAnalyticMock registers a success mock for updateAnalytic
func (m *AnalyticMock) RegisterUpdateAnalyticMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("updateAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("update_analytic_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterDeleteAnalyticMock registers a success mock for deleteAnalytic
func (m *AnalyticMock) RegisterDeleteAnalyticMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("deleteAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("delete_analytic_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListAnalyticsMock registers a success mock for listAnalytics
func (m *AnalyticMock) RegisterListAnalyticsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listAnalytics"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_analytics_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListAnalyticsLiteMock registers a success mock for listAnalytics (lite query)
func (m *AnalyticMock) RegisterListAnalyticsLiteMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listAnalyticsLite"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_analytics_lite_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListAnalyticsNamesMock registers a success mock for listAnalyticsNames
func (m *AnalyticMock) RegisterListAnalyticsNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listAnalyticsNames"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_analytics_names_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListAnalyticsCategoriesMock registers a success mock for listAnalyticsCategories
func (m *AnalyticMock) RegisterListAnalyticsCategoriesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listAnalyticsCategories"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_analytics_categories_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListAnalyticsTagsMock registers a success mock for listAnalyticsTags
func (m *AnalyticMock) RegisterListAnalyticsTagsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listAnalyticsTags"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_analytics_tags_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterListAnalyticsFilterOptionsMock registers a success mock for listAnalyticsFilterOptions
func (m *AnalyticMock) RegisterListAnalyticsFilterOptionsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listAnalyticsFilterOptions"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_analytics_filter_options_success.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *AnalyticMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(401, m.loadMockData("error_unauthorized.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// RegisterNotFoundErrorMock registers a not-found error mock
func (m *AnalyticMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(200, m.loadMockData("error_not_found.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// loadMockData loads mock JSON data from a file relative to this source file
func (m *AnalyticMock) loadMockData(filename string) []byte {
	_, currentFile, _, _ := runtime.Caller(0)
	mockDir := filepath.Dir(currentFile)
	mockFile := filepath.Join(mockDir, filename)

	data, err := os.ReadFile(mockFile)
	if err != nil {
		panic("Failed to load mock data: " + err.Error())
	}

	return data
}
