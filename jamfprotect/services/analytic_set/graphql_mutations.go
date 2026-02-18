package analyticset

// analyticSetMutationVariables returns GraphQL variables for createAnalyticSet/updateAnalyticSet mutations.
func analyticSetMutationVariables(req any, uuid string) map[string]any {
	var (
		name        string
		description string
		types       []string
		analytics   []string
	)

	switch r := req.(type) {
	case *CreateAnalyticSetRequest:
		name = r.Name
		description = r.Description
		types = r.Types
		analytics = r.Analytics
	case *UpdateAnalyticSetRequest:
		name = r.Name
		description = r.Description
		types = r.Types
		analytics = r.Analytics
	}

	vars := map[string]any{
		"name":             name,
		"description":      description,
		"types":            types,
		"analytics":        analytics,
		"RBAC_Plan":        true,
		"excludeAnalytics": false,
	}

	if uuid != "" {
		vars["uuid"] = uuid
	}

	return vars
}
