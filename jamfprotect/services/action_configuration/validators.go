package actionconfiguration

// ValidateActionConfigID is a no-op for CRUD compatibility.
func ValidateActionConfigID(id string) error {
	return nil
}

// ValidateCreateActionConfigRequest is a no-op; no allowed-value constraints in provider schema.
func ValidateCreateActionConfigRequest(req *CreateActionConfigRequest) error {
	return nil
}

// ValidateUpdateActionConfigRequest is a no-op; no allowed-value constraints in provider schema.
func ValidateUpdateActionConfigRequest(req *UpdateActionConfigRequest) error {
	return nil
}
