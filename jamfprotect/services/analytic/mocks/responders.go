package mocks

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// AnalyticMock provides mock responses for the Analytic service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createAnalytic": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Test Analytic",
						"label":       "test_analytic",
						"inputType":   "GPFSEvent",
						"filter":      "",
						"description": "A test analytic",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterGetAnalyticMock registers a success mock for getAnalytic
func (m *AnalyticMock) RegisterGetAnalyticMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getAnalytic": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Test Analytic",
						"label":       "test_analytic",
						"inputType":   "GPFSEvent",
						"filter":      "",
						"description": "A test analytic",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updateAnalytic": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Updated Analytic",
						"label":       "updated_analytic",
						"inputType":   "GPFSEvent",
						"filter":      "",
						"description": "An updated test analytic",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-02T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deleteAnalytic": map[string]any{
						"uuid": "test-uuid-1234",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListAnalyticsMock registers a success mock for listAnalytics
func (m *AnalyticMock) RegisterListAnalyticsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listAnalytics"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listAnalytics": map[string]any{
						"items": []map[string]any{
							{
								"uuid":        "test-uuid-1234",
								"name":        "Test Analytic",
								"label":       "test_analytic",
								"inputType":   "GPFSEvent",
								"description": "A test analytic",
							},
						},
						"pageInfo": map[string]any{
							"next":  nil,
							"total": 1,
						},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListAnalyticsLiteMock registers a success mock for listAnalyticsLite
func (m *AnalyticMock) RegisterListAnalyticsLiteMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listAnalyticsLite"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listAnalytics": map[string]any{
						"items": []map[string]any{
							{
								"uuid":        "test-uuid-1234",
								"name":        "Test Analytic",
								"label":       "test_analytic",
								"inputType":   "GPFSEvent",
								"description": "A test analytic",
								"tags":        []string{"security"},
							},
						},
						"pageInfo": map[string]any{
							"next":  nil,
							"total": 1,
						},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListAnalyticsNamesMock registers a success mock for listAnalyticsNames
func (m *AnalyticMock) RegisterListAnalyticsNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listAnalyticsNames"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listAnalyticsNames": map[string]any{
						"items": []map[string]any{
							{"name": "Test Analytic"},
						},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListAnalyticsCategoriesMock registers a success mock for listAnalyticsCategories
func (m *AnalyticMock) RegisterListAnalyticsCategoriesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listAnalyticsCategories"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listAnalyticsCategories": []map[string]any{
						{"value": "Security", "count": 5},
						{"value": "Network", "count": 3},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListAnalyticsTagsMock registers a success mock for listAnalyticsTags
func (m *AnalyticMock) RegisterListAnalyticsTagsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listAnalyticsTags"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listAnalyticsTags": []map[string]any{
						{"value": "endpoint", "count": 10},
						{"value": "network", "count": 7},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListAnalyticsFilterOptionsMock registers a success mock for listAnalyticsFilterOptions
func (m *AnalyticMock) RegisterListAnalyticsFilterOptionsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listAnalyticsFilterOptions"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listAnalyticsTags": []map[string]any{
						{"value": "endpoint", "count": 10},
					},
					"listAnalyticsCategories": []map[string]any{
						{"value": "Security", "count": 5},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *AnalyticMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(401, map[string]any{
				"errors": []map[string]any{
					{"message": "Unauthorized"},
				},
			})
			return resp, nil
		},
	)
}

// RegisterNotFoundErrorMock registers a not-found error mock
func (m *AnalyticMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getAnalytic"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getAnalytic": nil,
				},
				"errors": []map[string]any{
					{"message": "Analytic not found"},
				},
			})
			return resp, nil
		},
	)
}
