package mocks

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// USBControlSetMock provides mock responses for the RemovableStorageControlSet service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
type USBControlSetMock struct {
	baseURL string
}

// NewUSBControlSetMock creates a new USBControlSetMock instance
func NewUSBControlSetMock(baseURL string) *USBControlSetMock {
	return &USBControlSetMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for USB control set operations
func (m *USBControlSetMock) RegisterMocks() {
	m.RegisterCreateUSBControlSetMock()
	m.RegisterGetUSBControlSetMock()
	m.RegisterUpdateUSBControlSetMock()
	m.RegisterDeleteUSBControlSetMock()
	m.RegisterListUSBControlSetsMock()
	m.RegisterListUSBControlSetNamesMock()
}

// RegisterErrorMocks registers error response mocks
func (m *USBControlSetMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreateUSBControlSetMock registers a success mock for createUSBControlSet
func (m *USBControlSetMock) RegisterCreateUSBControlSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("createUSBControlSet"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createUSBControlSet": map[string]any{
						"id":          "test-id-1234",
						"name":        "Test USB Control Set",
						"description": "A test USB control set",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterGetUSBControlSetMock registers a success mock for getUSBControlSet
func (m *USBControlSetMock) RegisterGetUSBControlSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getUSBControlSet"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getUSBControlSet": map[string]any{
						"id":          "test-id-1234",
						"name":        "Test USB Control Set",
						"description": "A test USB control set",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUpdateUSBControlSetMock registers a success mock for updateUSBControlSet
func (m *USBControlSetMock) RegisterUpdateUSBControlSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("updateUSBControlSet"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updateUSBControlSet": map[string]any{
						"id":          "test-id-1234",
						"name":        "Updated USB Control Set",
						"description": "An updated USB control set",
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-02T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterDeleteUSBControlSetMock registers a success mock for deleteUSBControlSet
func (m *USBControlSetMock) RegisterDeleteUSBControlSetMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("deleteUSBControlSet"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deleteUSBControlSet": map[string]any{
						"id": "test-id-1234",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListUSBControlSetsMock registers a success mock for listUSBControlSets
func (m *USBControlSetMock) RegisterListUSBControlSetsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listUSBControlSets"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listUSBControlSets": map[string]any{
						"items": []map[string]any{
							{
								"id":          "test-id-1234",
								"name":        "Test USB Control Set",
								"description": "A test USB control set",
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

// RegisterListUSBControlSetNamesMock registers a success mock for listUsbControlNames
func (m *USBControlSetMock) RegisterListUSBControlSetNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listUsbControlNames"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listUsbControlNames": map[string]any{
						"items": []map[string]any{
							{"name": "Test USB Control Set"},
						},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *USBControlSetMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getUSBControlSet"),
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
func (m *USBControlSetMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getUSBControlSet"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getUSBControlSet": nil,
				},
				"errors": []map[string]any{
					{"message": "USB control set not found"},
				},
			})
			return resp, nil
		},
	)
}
