package analyticset

// ValidateAnalyticSetID is a no-op for CRUD compatibility.
func ValidateAnalyticSetID(id string) error {
	return nil
}

// ValidateCreateAnalyticSetRequest is a no-op; no allowed-value constraints in provider schema.
func ValidateCreateAnalyticSetRequest(req *CreateAnalyticSetRequest) error {
	return nil
}

// ValidateUpdateAnalyticSetRequest is a no-op; no allowed-value constraints in provider schema.
func ValidateUpdateAnalyticSetRequest(req *UpdateAnalyticSetRequest) error {
	return nil
}
