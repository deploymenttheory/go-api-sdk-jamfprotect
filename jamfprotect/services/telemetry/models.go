package telemetry

// TelemetryV2 represents a telemetry v2 configuration
type TelemetryV2 struct {
	ID                 string            `json:"id"`
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	Created            string            `json:"created"`
	Updated            string            `json:"updated"`
	LogFiles           []string          `json:"logFiles"`
	LogFileCollection  bool              `json:"logFileCollection"`
	PerformanceMetrics bool              `json:"performanceMetrics"`
	Plans              []TelemetryV2Plan `json:"plans"`
	Events             []string          `json:"events"`
	FileHashing        bool              `json:"fileHashing"`
}

// TelemetryV2Plan represents a plan entry on telemetry v2
type TelemetryV2Plan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateTelemetryV2Request is the request payload for creating telemetry v2
type CreateTelemetryV2Request struct {
	Name               string
	Description        string
	LogFiles           []string
	LogFileCollection  bool
	PerformanceMetrics bool
	Events             []string
	FileHashing        bool
}

// UpdateTelemetryV2Request is the request payload for updating telemetry v2
type UpdateTelemetryV2Request struct {
	Name               string
	Description        string
	LogFiles           []string
	LogFileCollection  bool
	PerformanceMetrics bool
	Events             []string
	FileHashing        bool
}

// ListTelemetriesV2Response represents the response from listing telemetries v2
type ListTelemetriesV2Response struct {
	Items    []TelemetryV2 `json:"items"`
	PageInfo PageInfo      `json:"pageInfo"`
}

// PageInfo contains pagination information
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
