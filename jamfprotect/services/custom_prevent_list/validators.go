package custompreventlist

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/validate"
)

// Allowed values from provider schema (prevent_type).
const (
	PreventTypeTEAMID    = "TEAMID"
	PreventTypeFILEHASH  = "FILEHASH"
	PreventTypeCDHASH    = "CDHASH"
	PreventTypeSIGNINGID = "SIGNINGID"
)

// ValidatePreventListType validates prevent list type is an allowed enum value.
func ValidatePreventListType(typ string) error {
	return validate.OneOf("type", typ, PreventTypeTEAMID, PreventTypeFILEHASH, PreventTypeCDHASH, PreventTypeSIGNINGID)
}

// ValidateCreatePreventListRequest validates allowed-value constraints on create prevent list request.
func ValidateCreatePreventListRequest(req *CreatePreventListRequest) error {
	if req == nil {
		return nil
	}
	return ValidatePreventListType(req.Type)
}

// ValidateUpdatePreventListRequest validates allowed-value constraints on update prevent list request.
func ValidateUpdatePreventListRequest(req *UpdatePreventListRequest) error {
	if req == nil {
		return nil
	}
	return ValidatePreventListType(req.Type)
}

// ValidatePreventListID is a no-op for CRUD compatibility.
func ValidatePreventListID(id string) error {
	return nil
}
