package mocks

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// ActionConfigMock provides mock responses for the ActionConfiguration service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
type ActionConfigMock struct {
	baseURL string
}

// NewActionConfigMock creates a new ActionConfigMock instance
func NewActionConfigMock(baseURL string) *ActionConfigMock {
	return &ActionConfigMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for action configuration operations
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createActionConfigs": map[string]any{
						"id":          "test-id-1234",
						"name":        "Test Action Config",
						"description": "A test action configuration",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getActionConfigs": map[string]any{
						"id":          "test-id-1234",
						"name":        "Test Action Config",
						"description": "A test action configuration",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updateActionConfigs": map[string]any{
						"id":          "test-id-1234",
						"name":        "Updated Action Config",
						"description": "An updated action configuration",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-02T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deleteActionConfigs": map[string]any{
						"id": "test-id-1234",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listActionConfigs": map[string]any{
						"items": []map[string]any{
							{
								"id":          "test-id-1234",
								"name":        "Test Action Config",
								"description": "A test action configuration",
								"created":     "2024-01-01T00:00:00Z",
								"updated":     "2024-01-01T00:00:00Z",
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

// RegisterListActionConfigNamesMock registers a success mock for listActionConfigNames
func (m *ActionConfigMock) RegisterListActionConfigNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listActionConfigNames"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listActionConfigNames": map[string]any{
						"items": []map[string]any{
							{"name": "Test Action Config"},
						},
					},
				},
			})
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
func (m *ActionConfigMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getActionConfigs"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getActionConfigs": nil,
				},
				"errors": []map[string]any{
					{"message": "Action configuration not found"},
				},
			})
			return resp, nil
		},
	)
}
