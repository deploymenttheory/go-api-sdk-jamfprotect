package mocks

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// PlanMock provides mock responses for the Plan service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
type PlanMock struct {
	baseURL string
}

// NewPlanMock creates a new PlanMock instance
func NewPlanMock(baseURL string) *PlanMock {
	return &PlanMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for plan operations
func (m *PlanMock) RegisterMocks() {
	m.RegisterCreatePlanMock()
	m.RegisterGetPlanMock()
	m.RegisterUpdatePlanMock()
	m.RegisterDeletePlanMock()
	m.RegisterListPlansMock()
	m.RegisterListPlanNamesMock()
	m.RegisterGetPlanConfigurationAndSetOptionsMock()
}

// RegisterErrorMocks registers error response mocks
func (m *PlanMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreatePlanMock registers a success mock for createPlan
func (m *PlanMock) RegisterCreatePlanMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("createPlan"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createPlan": map[string]any{
						"id":          "test-id-1234",
						"name":        "Test Plan",
						"description": "A test plan",
						"autoUpdate":  false,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterGetPlanMock registers a success mock for getPlan
func (m *PlanMock) RegisterGetPlanMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getPlan"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getPlan": map[string]any{
						"id":          "test-id-1234",
						"name":        "Test Plan",
						"description": "A test plan",
						"autoUpdate":  false,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUpdatePlanMock registers a success mock for updatePlan
func (m *PlanMock) RegisterUpdatePlanMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("updatePlan"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updatePlan": map[string]any{
						"id":          "test-id-1234",
						"name":        "Updated Plan",
						"description": "An updated plan",
						"autoUpdate":  false,
						"created":     "2024-01-01T00:00:00Z",
						"updated":     "2024-01-02T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterDeletePlanMock registers a success mock for deletePlan
func (m *PlanMock) RegisterDeletePlanMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("deletePlan"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deletePlan": map[string]any{
						"id": "test-id-1234",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListPlansMock registers a success mock for listPlans
func (m *PlanMock) RegisterListPlansMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listPlans"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listPlans": map[string]any{
						"items": []map[string]any{
							{
								"id":          "test-id-1234",
								"name":        "Test Plan",
								"description": "A test plan",
								"autoUpdate":  false,
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

// RegisterListPlanNamesMock registers a success mock for listPlanNames
func (m *PlanMock) RegisterListPlanNamesMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listPlanNames"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listPlanNames": map[string]any{
						"items": []map[string]any{
							{"name": "Test Plan"},
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

// RegisterGetPlanConfigurationAndSetOptionsMock registers a success mock for getPlanConfigurationAndSetOptions
func (m *PlanMock) RegisterGetPlanConfigurationAndSetOptionsMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getPlanConfigurationAndSetOptions"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"actionConfigs": map[string]any{
						"items": []map[string]any{
							{"id": "ac-id-1", "name": "Action Config 1"},
						},
					},
					"telemetries": map[string]any{
						"items": []map[string]any{
							{"id": "tel-id-1", "name": "Telemetry 1"},
						},
					},
					"telemetriesV2": map[string]any{
						"items": []map[string]any{
							{"id": "telv2-id-1", "name": "Telemetry V2 1"},
						},
					},
					"usbControlSets": map[string]any{
						"items": []map[string]any{
							{"id": "usb-id-1", "name": "USB Control Set 1"},
						},
					},
					"exceptionSets": map[string]any{
						"items": []map[string]any{
							{"uuid": "exc-uuid-1", "name": "Exception Set 1", "managed": false},
						},
					},
					"analyticSets": map[string]any{
						"items": []map[string]any{
							{"uuid": "as-uuid-1", "name": "Analytic Set 1", "managed": false, "types": []string{}},
						},
					},
					"managedAnalyticSets": map[string]any{
						"items": []map[string]any{
							{"uuid": "mas-uuid-1", "name": "Managed Analytic Set 1", "managed": true, "types": []string{}},
						},
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock
func (m *PlanMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getPlan"),
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
func (m *PlanMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getPlan"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getPlan": nil,
				},
				"errors": []map[string]any{
					{"message": "Plan not found"},
				},
			})
			return resp, nil
		},
	)
}
