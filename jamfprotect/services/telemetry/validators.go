package telemetry

// ValidateTelemetryV2ID is a no-op for CRUD compatibility.
func ValidateTelemetryV2ID(id string) error {
	return nil
}

// ValidateCreateTelemetryV2Request is a no-op; no allowed-value constraints in provider schema.
func ValidateCreateTelemetryV2Request(req *CreateTelemetryV2Request) error {
	return nil
}

// ValidateUpdateTelemetryV2Request is a no-op; no allowed-value constraints in provider schema.
func ValidateUpdateTelemetryV2Request(req *UpdateTelemetryV2Request) error {
	return nil
}
