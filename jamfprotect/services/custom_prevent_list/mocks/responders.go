package mocks

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// PreventListMock provides mock responses for the CustomPreventList service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
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
		m.baseURL+"/app",
		httpmock.BodyContainsString("createPreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createPreventList": map[string]any{
						"id":          "test-id-1234",
						"name":        "Test Prevent List",
						"description": "A test prevent list",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterGetPreventListMock registers a success mock for getPreventList
func (m *PreventListMock) RegisterGetPreventListMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getPreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getPreventList": map[string]any{
						"id":          "test-id-1234",
						"name":        "Test Prevent List",
						"description": "A test prevent list",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUpdatePreventListMock registers a success mock for updatePreventList
func (m *PreventListMock) RegisterUpdatePreventListMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("updatePreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updatePreventList": map[string]any{
						"id":          "test-id-1234",
						"name":        "Updated Prevent List",
						"description": "An updated prevent list",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-02T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterDeletePreventListMock registers a success mock for deletePreventList
func (m *PreventListMock) RegisterDeletePreventListMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("deletePreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deletePreventList": map[string]any{
						"id": "test-id-1234",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListPreventListsMock registers a success mock for listPreventLists
func (m *PreventListMock) RegisterListPreventListsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listPreventLists"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listPreventLists": map[string]any{
						"items": []map[string]any{
							{
								"id":          "test-id-1234",
								"name":        "Test Prevent List",
								"description": "A test prevent list",
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

// RegisterListPreventListNamesMock registers a success mock for listPreventListNames
func (m *PreventListMock) RegisterListPreventListNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listPreventListNames"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listPreventListNames": map[string]any{
						"items": []map[string]any{
							{"name": "Test Prevent List"},
						},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *PreventListMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getPreventList"),
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
func (m *PreventListMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getPreventList"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getPreventList": nil,
				},
				"errors": []map[string]any{
					{"message": "Prevent list not found"},
				},
			})
			return resp, nil
		},
	)
}
