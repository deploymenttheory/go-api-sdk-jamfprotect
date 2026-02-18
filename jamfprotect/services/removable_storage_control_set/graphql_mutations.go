package removablestoragecontrolset

// usbControlSetMutationVariables returns GraphQL variables for createUSBControlSet/updateUSBControlSet mutations.
func usbControlSetMutationVariables(req any, id string) map[string]any {
	var (
		name                 string
		description          string
		defaultMountAction   string
		defaultMessageAction string
		rules                []USBControlRuleInput
	)

	switch r := req.(type) {
	case *CreateUSBControlSetRequest:
		name = r.Name
		description = r.Description
		defaultMountAction = r.DefaultMountAction
		defaultMessageAction = r.DefaultMessageAction
		rules = r.Rules
	case *UpdateUSBControlSetRequest:
		name = r.Name
		description = r.Description
		defaultMountAction = r.DefaultMountAction
		defaultMessageAction = r.DefaultMessageAction
		rules = r.Rules
	}

	vars := map[string]any{
		"name":                 name,
		"description":          description,
		"defaultMountAction":   defaultMountAction,
		"defaultMessageAction": defaultMessageAction,
		"rules":                rules,
	}

	if id != "" {
		vars["id"] = id
	}

	return vars
}
