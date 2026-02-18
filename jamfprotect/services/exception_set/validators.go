package exceptionset

import (
	"fmt"
	"regexp"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/validate"
)

// uuidRegex matches a canonical UUID string (8-4-4-4-12 hex digits).
var uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// ValidateExceptionSetUUID checks that uuid is non-empty and matches UUID format.
func ValidateExceptionSetUUID(uuid string) error {
	if uuid == "" {
		return fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}
	if !uuidRegex.MatchString(uuid) {
		return fmt.Errorf("%w: uuid must be a valid UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)", client.ErrInvalidInput)
	}
	return nil
}

// Allowed values from provider schema (exception block: type, ignore_activity).
const (
	ExceptionTypeUser           = "User"
	ExceptionTypeAppSigningInfo = "AppSigningInfo"
	ExceptionTypeTeamId         = "TeamId"
	ExceptionTypeExecutable     = "Executable"
	ExceptionTypePlatformBinary = "PlatformBinary"
	ExceptionTypePath           = "Path"

	IgnoreActivityAnalytics        = "Analytics"
	IgnoreActivityThreatPrevention = "ThreatPrevention"
	IgnoreActivityTelemetryV2      = "TelemetryV2"
	IgnoreActivityTelemetry        = "Telemetry"
)

// ValidateExceptionType validates exception type is an allowed enum value.
func ValidateExceptionType(typ string) error {
	return validate.OneOf("exception.type", typ,
		ExceptionTypeUser, ExceptionTypeAppSigningInfo, ExceptionTypeTeamId,
		ExceptionTypeExecutable, ExceptionTypePlatformBinary, ExceptionTypePath)
}

// ValidateIgnoreActivity validates ignore_activity is an allowed enum value.
func ValidateIgnoreActivity(activity string) error {
	return validate.OneOf("exception.ignore_activity", activity,
		IgnoreActivityAnalytics, IgnoreActivityThreatPrevention, IgnoreActivityTelemetryV2, IgnoreActivityTelemetry)
}

// ValidateExceptionInput validates allowed-value constraints on a single exception input.
func ValidateExceptionInput(ex ExceptionInput) error {
	if err := ValidateExceptionType(ex.Type); err != nil {
		return err
	}
	return ValidateIgnoreActivity(ex.IgnoreActivity)
}

// ValidateEsExceptionInput validates allowed-value constraints on a single ES exception input.
func ValidateEsExceptionInput(ex EsExceptionInput) error {
	if err := ValidateExceptionType(ex.Type); err != nil {
		return err
	}
	return ValidateIgnoreActivity(ex.IgnoreActivity)
}

// ValidateCreateExceptionSetRequest validates allowed-value constraints on create request.
func ValidateCreateExceptionSetRequest(req *CreateExceptionSetRequest) error {
	if req == nil {
		return nil
	}
	for _, ex := range req.Exceptions {
		if err := ValidateExceptionInput(ex); err != nil {
			return err
		}
	}
	for _, ex := range req.EsExceptions {
		if err := ValidateEsExceptionInput(ex); err != nil {
			return err
		}
	}
	return nil
}

// ValidateUpdateExceptionSetRequest validates allowed-value constraints on update request.
func ValidateUpdateExceptionSetRequest(req *UpdateExceptionSetRequest) error {
	if req == nil {
		return nil
	}
	for _, ex := range req.Exceptions {
		if err := ValidateExceptionInput(ex); err != nil {
			return err
		}
	}
	for _, ex := range req.EsExceptions {
		if err := ValidateEsExceptionInput(ex); err != nil {
			return err
		}
	}
	return nil
}

// ValidateExceptionSetID is a no-op for CRUD compatibility.
func ValidateExceptionSetID(id string) error {
	return nil
}
