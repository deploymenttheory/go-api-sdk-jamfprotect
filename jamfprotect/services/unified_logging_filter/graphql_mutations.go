package unifiedloggingfilter

// unifiedLoggingFilterMutationVariables returns GraphQL variables for createUnifiedLoggingFilter/updateUnifiedLoggingFilter mutations.
func unifiedLoggingFilterMutationVariables(req any) map[string]any {
	var (
		name        string
		description string
		tags        []string
		filter      string
		enabled     bool
	)

	switch r := req.(type) {
	case *CreateUnifiedLoggingFilterRequest:
		name = r.Name
		description = r.Description
		tags = r.Tags
		filter = r.Filter
		enabled = r.Enabled
	case *UpdateUnifiedLoggingFilterRequest:
		name = r.Name
		description = r.Description
		tags = r.Tags
		filter = r.Filter
		enabled = r.Enabled
	}

	return map[string]any{
		"name":        name,
		"description": description,
		"tags":        tags,
		"filter":      filter,
		"enabled":     enabled,
	}
}
