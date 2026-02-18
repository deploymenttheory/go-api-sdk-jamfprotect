package unifiedloggingfilter

// ValidateUnifiedLoggingFilterUUID is a no-op for CRUD compatibility.
func ValidateUnifiedLoggingFilterUUID(uuid string) error {
	return nil
}

// ValidateCreateUnifiedLoggingFilterRequest is a no-op; no allowed-value constraints in provider schema.
func ValidateCreateUnifiedLoggingFilterRequest(req *CreateUnifiedLoggingFilterRequest) error {
	return nil
}

// ValidateUpdateUnifiedLoggingFilterRequest is a no-op; no allowed-value constraints in provider schema.
func ValidateUpdateUnifiedLoggingFilterRequest(req *UpdateUnifiedLoggingFilterRequest) error {
	return nil
}
