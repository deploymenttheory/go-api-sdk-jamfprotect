package mocks

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// TelemetryMock provides mock responses for the Telemetry service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
type TelemetryMock struct {
	baseURL string
}

// NewTelemetryMock creates a new TelemetryMock instance
func NewTelemetryMock(baseURL string) *TelemetryMock {
	return &TelemetryMock{baseURL: baseURL}
}

// RegisterMocks registers all successful response mocks for telemetry operations
func (m *TelemetryMock) RegisterMocks() {
	m.RegisterCreateTelemetryV2Mock()
	m.RegisterGetTelemetryV2Mock()
	m.RegisterUpdateTelemetryV2Mock()
	m.RegisterDeleteTelemetryV2Mock()
	m.RegisterListTelemetriesV2Mock()
	m.RegisterListTelemetriesCombinedMock()
}

// RegisterErrorMocks registers error response mocks
func (m *TelemetryMock) RegisterErrorMocks() {
	m.RegisterUnauthorizedErrorMock()
	m.RegisterNotFoundErrorMock()
}

// RegisterCreateTelemetryV2Mock registers a success mock for createTelemetryV2
func (m *TelemetryMock) RegisterCreateTelemetryV2Mock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("createTelemetryV2"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"createTelemetryV2": map[string]any{
						"id":                 "test-id-1234",
						"name":               "Test Telemetry V2",
						"description":        "A test telemetry v2",
						"logFiles":           []string{},
						"logFileCollection":  false,
						"performanceMetrics": false,
						"events":             []string{},
						"fileHashing":        false,
						"created":            "2024-01-01T00:00:00Z",
						"updated":            "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterGetTelemetryV2Mock registers a success mock for getTelemetryV2
func (m *TelemetryMock) RegisterGetTelemetryV2Mock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getTelemetryV2"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getTelemetryV2": map[string]any{
						"id":                 "test-id-1234",
						"name":               "Test Telemetry V2",
						"description":        "A test telemetry v2",
						"logFiles":           []string{},
						"logFileCollection":  false,
						"performanceMetrics": false,
						"events":             []string{},
						"fileHashing":        false,
						"created":            "2024-01-01T00:00:00Z",
						"updated":            "2024-01-01T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterUpdateTelemetryV2Mock registers a success mock for updateTelemetryV2
func (m *TelemetryMock) RegisterUpdateTelemetryV2Mock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("updateTelemetryV2"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"updateTelemetryV2": map[string]any{
						"id":                 "test-id-1234",
						"name":               "Updated Telemetry V2",
						"description":        "An updated telemetry v2",
						"logFiles":           []string{},
						"logFileCollection":  false,
						"performanceMetrics": false,
						"events":             []string{},
						"fileHashing":        false,
						"created":            "2024-01-01T00:00:00Z",
						"updated":            "2024-01-02T00:00:00Z",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterDeleteTelemetryV2Mock registers a success mock for deleteTelemetryV2
func (m *TelemetryMock) RegisterDeleteTelemetryV2Mock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("deleteTelemetryV2"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"deleteTelemetryV2": map[string]any{
						"id": "test-id-1234",
					},
				},
			})
			return resp, nil
		},
	)
}

// RegisterListTelemetriesV2Mock registers a success mock for listTelemetriesV2
func (m *TelemetryMock) RegisterListTelemetriesV2Mock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listTelemetriesV2"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listTelemetriesV2": map[string]any{
						"items": []map[string]any{
							{
								"id":          "test-id-1234",
								"name":        "Test Telemetry V2",
								"description": "A test telemetry v2",
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

// RegisterListTelemetriesCombinedMock registers a success mock for listTelemetriesCombined
func (m *TelemetryMock) RegisterListTelemetriesCombinedMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("listTelemetriesCombined"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"listTelemetries": map[string]any{
						"items": []map[string]any{
							{
								"id":          "tel-id-1",
								"name":        "Test Telemetry V1",
								"description": "A test v1 telemetry",
								"verbose":     false,
								"level":       0,
							},
						},
						"pageInfo": map[string]any{
							"next":  nil,
							"total": 1,
						},
					},
					"listTelemetriesV2": map[string]any{
						"items": []map[string]any{
							{
								"id":          "telv2-id-1",
								"name":        "Test Telemetry V2",
								"description": "A test v2 telemetry",
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
func (m *TelemetryMock) RegisterUnauthorizedErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getTelemetryV2"),
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
func (m *TelemetryMock) RegisterNotFoundErrorMock() {
	httpmock.RegisterMatcherResponder(
		"POST",
		m.baseURL+"/app",
		httpmock.BodyContainsString("getTelemetryV2"),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, map[string]any{
				"data": map[string]any{
					"getTelemetryV2": nil,
				},
				"errors": []map[string]any{
					{"message": "Telemetry not found"},
				},
			})
			return resp, nil
		},
	)
}
