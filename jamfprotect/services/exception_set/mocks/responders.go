package mocks

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// ExceptionSetMock provides mock responses for the ExceptionSet service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createExceptionSet": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Test Exception Set",
						"description": "A test exception set",
						"managed":     false,
						"exceptions":  []any{},
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getExceptionSet": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Test Exception Set",
						"description": "A test exception set",
						"managed":     false,
						"exceptions":  []any{},
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updateExceptionSet": map[string]any{
						"uuid":        "test-uuid-1234",
						"name":        "Updated Exception Set",
						"description": "An updated exception set",
						"managed":     false,
						"exceptions":  []any{},
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-02T00:00:00Z",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deleteExceptionSet": map[string]any{
						"uuid": "test-uuid-1234",
					},
				},
			})
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
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listExceptionSets": map[string]any{
						"items": []map[string]any{
							{
								"uuid":    "test-uuid-1234",
								"name":    "Test Exception Set",
								"managed": false,
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

// RegisterListExceptionSetNamesMock registers a success mock for listExceptionSetNames
func (m *ExceptionSetMock) RegisterListExceptionSetNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listExceptionSetNames"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listExceptionSetNames": map[string]any{
						"items": []map[string]any{
							{"name": "Test Exception Set"},
						},
					},
				},
			})
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
func (m *ExceptionSetMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getExceptionSet"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getExceptionSet": nil,
				},
				"errors": []map[string]any{
					{"message": "Exception set not found"},
				},
			})
			return resp, nil
		},
	)
}
