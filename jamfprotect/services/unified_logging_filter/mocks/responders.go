package mocks

import (
	"net/http"

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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createUnifiedLoggingFilter": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Test Unified Logging Filter",
						"description": "A test unified logging filter",
						"enabled":     true,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getUnifiedLoggingFilter": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Test Unified Logging Filter",
						"description": "A test unified logging filter",
						"enabled":     true,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updateUnifiedLoggingFilter": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Updated Unified Logging Filter",
						"description": "An updated unified logging filter",
						"enabled":     true,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-02T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deleteUnifiedLoggingFilter": map[string]any{
						"uuid": "test-uuid-1234",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listUnifiedLoggingFilters": map[string]any{
						"items": []map[string]any{
							{
								"uuid":        "test-uuid-1234",
								"name":        "Test Unified Logging Filter",
								"description": "A test unified logging filter",
								"enabled":     true,
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

// RegisterListUnifiedLoggingFilterNamesMock registers a success mock for listUnifiedLoggingFilterNames
func (m *UnifiedLoggingFilterMock) RegisterListUnifiedLoggingFilterNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("listUnifiedLoggingFilterNames"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listUnifiedLoggingFilterNames": map[string]any{
						"items": []map[string]any{
							{"name": "Test Unified Logging Filter"},
						},
					},
				},
			})
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
func (m *UnifiedLoggingFilterMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/graphql",
		httpmock.BodyContainsString("getUnifiedLoggingFilter"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getUnifiedLoggingFilter": nil,
				},
				"errors": []map[string]any{
					{"message": "Unified logging filter not found"},
				},
			})
			return resp, nil
		},
	)
}
