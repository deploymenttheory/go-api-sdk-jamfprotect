package mocks

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"

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
			resp := httpmock.NewBytesResponse(200, m.loadMockData("create_telemetry_v2_success.json"))
			resp.Header.Set("Content-Type", "application/json")
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
			resp := httpmock.NewBytesResponse(200, m.loadMockData("get_telemetry_v2_success.json"))
			resp.Header.Set("Content-Type", "application/json")
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
			resp := httpmock.NewBytesResponse(200, m.loadMockData("update_telemetry_v2_success.json"))
			resp.Header.Set("Content-Type", "application/json")
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
			resp := httpmock.NewBytesResponse(200, m.loadMockData("delete_telemetry_v2_success.json"))
			resp.Header.Set("Content-Type", "application/json")
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
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_telemetries_v2_success.json"))
			resp.Header.Set("Content-Type", "application/json")
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
			resp := httpmock.NewBytesResponse(200, m.loadMockData("list_telemetries_combined_success.json"))
			resp.Header.Set("Content-Type", "application/json")
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
			resp := httpmock.NewBytesResponse(401, m.loadMockData("error_unauthorized.json"))
			resp.Header.Set("Content-Type", "application/json")
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
			resp := httpmock.NewBytesResponse(200, m.loadMockData("error_not_found.json"))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)
}

// loadMockData loads mock JSON data from a file relative to this source file
func (m *TelemetryMock) loadMockData(filename string) []byte {
	_, currentFile, _, _ := runtime.Caller(0)
	mockDir := filepath.Dir(currentFile)
	mockFile := filepath.Join(mockDir, filename)

	data, err := os.ReadFile(mockFile)
	if err != nil {
		panic("Failed to load mock data: " + err.Error())
	}

	return data
}
