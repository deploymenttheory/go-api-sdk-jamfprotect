package removablestoragecontrolset

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/validate"
)

// Allowed values for default mount action / permission (provider: default_permission and override permission).
const (
	MountActionReadOnly  = "ReadOnly"
	MountActionReadWrite = "ReadWrite"
	MountActionPrevented = "Prevented"
)

// ValidateDefaultMountAction validates default mount action is an allowed enum value.
func ValidateDefaultMountAction(action string) error {
	return validate.OneOf("defaultMountAction", action, MountActionReadOnly, MountActionReadWrite, MountActionPrevented)
}

// ValidateRuleMountAction validates a rule's mount action is an allowed value.
func ValidateRuleMountAction(action string) error {
	return validate.OneOf("rules[].mountAction", action, MountActionReadOnly, MountActionReadWrite, MountActionPrevented)
}

// ValidateCreateUSBControlSetRequest validates allowed-value constraints on create request.
func ValidateCreateUSBControlSetRequest(req *CreateUSBControlSetRequest) error {
	if req == nil {
		return nil
	}
	if err := ValidateDefaultMountAction(req.DefaultMountAction); err != nil {
		return err
	}
	for i, r := range req.Rules {
		ma := mountActionFromRule(r)
		if ma != "" {
			if err := ValidateRuleMountAction(ma); err != nil {
				return err
			}
		}
		_ = i
	}
	return nil
}

// ValidateUpdateUSBControlSetRequest validates allowed-value constraints on update request.
func ValidateUpdateUSBControlSetRequest(req *UpdateUSBControlSetRequest) error {
	if req == nil {
		return nil
	}
	if err := ValidateDefaultMountAction(req.DefaultMountAction); err != nil {
		return err
	}
	for _, r := range req.Rules {
		ma := mountActionFromRule(r)
		if ma != "" {
			if err := ValidateRuleMountAction(ma); err != nil {
				return err
			}
		}
	}
	return nil
}

func mountActionFromRule(r USBControlRuleInput) string {
	if r.VendorRule != nil && r.VendorRule.MountAction != "" {
		return r.VendorRule.MountAction
	}
	if r.SerialRule != nil && r.SerialRule.MountAction != "" {
		return r.SerialRule.MountAction
	}
	if r.ProductRule != nil && r.ProductRule.MountAction != "" {
		return r.ProductRule.MountAction
	}
	if r.EncryptionRule != nil && r.EncryptionRule.MountAction != "" {
		return r.EncryptionRule.MountAction
	}
	return ""
}

// ValidateUSBControlSetID is a no-op for CRUD compatibility.
func ValidateUSBControlSetID(id string) error {
	return nil
}
