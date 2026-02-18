package plan

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/validate"
)

// Allowed values from provider schema / API enums.

const (
	LogLevelDISABLED = "DISABLED"
	LogLevelERROR    = "ERROR"
	LogLevelWARNING  = "WARNING"
	LogLevelINFO     = "INFO"
	LogLevelDEBUG    = "DEBUG"
)

const (
	ProtocolMQTT    = "mqtt"
	ProtocolWSSMQTT = "wss/mqtt"
)

const (
	SignaturesModeBlocking   = "blocking"
	SignaturesModeReportOnly = "reportOnly"
	SignaturesModeDisabled   = "disabled"
)

// ValidateLogLevel validates plan log level is an allowed enum value.
func ValidateLogLevel(logLevel *string) error {
	if logLevel == nil || *logLevel == "" {
		return nil
	}
	return validate.OneOf("logLevel", *logLevel, LogLevelDISABLED, LogLevelERROR, LogLevelWARNING, LogLevelINFO, LogLevelDEBUG)
}

// ValidateCommsProtocol validates communications protocol is an allowed value.
func ValidateCommsProtocol(protocol string) error {
	return validate.OneOf("commsConfig.protocol", protocol, ProtocolMQTT, ProtocolWSSMQTT)
}

// ValidateSignaturesFeedMode validates signatures feed mode is an allowed value.
func ValidateSignaturesFeedMode(mode string) error {
	if mode == "" {
		return nil
	}
	return validate.OneOf("signaturesFeedConfig.mode", mode, SignaturesModeBlocking, SignaturesModeReportOnly, SignaturesModeDisabled)
}

// ValidateCreatePlanRequest validates allowed-value constraints on create plan request.
func ValidateCreatePlanRequest(req *CreatePlanRequest) error {
	if req == nil {
		return nil
	}
	if err := ValidateLogLevel(req.LogLevel); err != nil {
		return err
	}
	if err := ValidateCommsProtocol(req.CommsConfig.Protocol); err != nil {
		return err
	}
	if err := ValidateSignaturesFeedMode(req.SignaturesFeedConfig.Mode); err != nil {
		return err
	}
	return nil
}

// ValidateUpdatePlanRequest validates allowed-value constraints on update plan request.
func ValidateUpdatePlanRequest(req *UpdatePlanRequest) error {
	if req == nil {
		return nil
	}
	if err := ValidateLogLevel(req.LogLevel); err != nil {
		return err
	}
	if err := ValidateCommsProtocol(req.CommsConfig.Protocol); err != nil {
		return err
	}
	if err := ValidateSignaturesFeedMode(req.SignaturesFeedConfig.Mode); err != nil {
		return err
	}
	return nil
}
