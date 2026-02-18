package mocks

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// AnalyticSetMock provides mock responses for the AnalyticSet service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createAnalyticSet": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Test Analytic Set",
						"description": "A test analytic set",
						"managed":     false,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getAnalyticSet": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Test Analytic Set",
						"description": "A test analytic set",
						"managed":     false,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUpdateAnalyticSetMock registers a success mock for updateAnalyticSet
func (m *AnalyticSetMock) RegisterUpdateAnalyticSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("updateAnalyticSet"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updateAnalyticSet": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Updated Analytic Set",
						"description": "An updated analytic set",
						"managed":     false,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-02T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deleteAnalyticSet": map[string]any{
						"uuid": "test-uuid-1234",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listAnalyticSets": map[string]any{
						"items": []map[string]any{
							{
								"uuid":        "test-uuid-1234",
								"name":        "Test Analytic Set",
								"description": "A test analytic set",
								"managed":     false,
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

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *AnalyticSetMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getAnalyticSet"),
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
func (m *AnalyticSetMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getAnalyticSet"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getAnalyticSet": nil,
				},
				"errors": []map[string]any{
					{"message": "Analytic set not found"},
				},
			})
			return resp, nil
		},
	)
}
