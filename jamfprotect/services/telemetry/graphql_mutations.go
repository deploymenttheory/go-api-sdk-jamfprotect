package telemetry

// telemetryMutationVariables returns GraphQL variables for createTelemetryV2/updateTelemetryV2 mutations.
func telemetryMutationVariables(req any) map[string]any {
	var (
		name               string
		description        string
		logFiles           []string
		logFileCollection  bool
		performanceMetrics bool
		events             []string
		fileHashing        bool
	)

	switch r := req.(type) {
	case *CreateTelemetryV2Request:
		name = r.Name
		description = r.Description
		logFiles = r.LogFiles
		logFileCollection = r.LogFileCollection
		performanceMetrics = r.PerformanceMetrics
		events = r.Events
		fileHashing = r.FileHashing
	case *UpdateTelemetryV2Request:
		name = r.Name
		description = r.Description
		logFiles = r.LogFiles
		logFileCollection = r.LogFileCollection
		performanceMetrics = r.PerformanceMetrics
		events = r.Events
		fileHashing = r.FileHashing
	}

	return map[string]any{
		"name":               name,
		"description":        description,
		"logFiles":           logFiles,
		"logFileCollection":  logFileCollection,
		"performanceMetrics": performanceMetrics,
		"events":             events,
		"fileHashing":        fileHashing,
	}
}
