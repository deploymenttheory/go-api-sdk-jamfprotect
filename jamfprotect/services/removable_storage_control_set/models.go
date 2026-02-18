package removablestoragecontrolset

// USBControlSet represents a USB control set in API responses.
type USBControlSet struct {
	ID                   string              `json:"id"`
	Name                 string              `json:"name"`
	Description          string              `json:"description"`
	DefaultMountAction   string              `json:"defaultMountAction"`
	DefaultMessageAction string              `json:"defaultMessageAction"`
	Rules                []USBControlRule    `json:"rules"`
	Plans                []USBControlSetPlan `json:"plans"`
	Created              string              `json:"created"`
	Updated              string              `json:"updated"`
}

// USBControlSetPlan represents a plan entry in a USB control set.
type USBControlSetPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// USBControlRule represents a USB control rule in API responses.
type USBControlRule struct {
	Type          string                  `json:"type"`
	MountAction   string                  `json:"mountAction"`
	MessageAction string                  `json:"messageAction"`
	ApplyTo       string                  `json:"applyTo"`
	Vendors       []string                `json:"vendors"`
	Serials       []string                `json:"serials"`
	Products      []USBControlProductPair `json:"products"`
}

// USBControlProductPair represents a vendor+product pair.
type USBControlProductPair struct {
	Vendor  string `json:"vendor"`
	Product string `json:"product"`
}

// CreateUSBControlSetRequest is the request payload for creating a USB control set.
type CreateUSBControlSetRequest struct {
	Name                 string
	Description          string
	DefaultMountAction   string
	DefaultMessageAction string
	Rules                []USBControlRuleInput
}

// UpdateUSBControlSetRequest is the request payload for updating a USB control set.
type UpdateUSBControlSetRequest struct {
	Name                 string
	Description          string
	DefaultMountAction   string
	DefaultMessageAction string
	Rules                []USBControlRuleInput
}

// USBControlRuleInput represents a USB control rule input variant.
type USBControlRuleInput struct {
	Type           string                        `json:"type"`
	VendorRule     *USBControlRuleDetails        `json:"vendorRule,omitempty"`
	SerialRule     *USBControlRuleDetails        `json:"serialRule,omitempty"`
	ProductRule    *USBControlProductRuleDetails `json:"productRule,omitempty"`
	EncryptionRule *USBControlRuleDetails        `json:"encryptionRule,omitempty"`
}

// USBControlRuleDetails represents shared rule fields.
type USBControlRuleDetails struct {
	MountAction   string   `json:"mountAction"`
	MessageAction *string  `json:"messageAction,omitempty"`
	ApplyTo       *string  `json:"applyTo,omitempty"`
	Vendors       []string `json:"vendors,omitempty"`
	Serials       []string `json:"serials,omitempty"`
}

// USBControlProductRuleDetails represents product rule details.
type USBControlProductRuleDetails struct {
	MountAction   string                  `json:"mountAction"`
	MessageAction *string                 `json:"messageAction,omitempty"`
	ApplyTo       *string                 `json:"applyTo,omitempty"`
	Products      []USBControlProductPair `json:"products,omitempty"`
}

// ListUSBControlSetsResponse represents the response from listing USB control sets.
type ListUSBControlSetsResponse struct {
	Items    []USBControlSet `json:"items"`
	PageInfo PageInfo        `json:"pageInfo"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}

// USBControlSetName is a lightweight USB control set containing only the name
type USBControlSetName struct {
	Name string `json:"name"`
}

// ListUSBControlSetNamesResponse is the response wrapper for listing USB control set names
type ListUSBControlSetNamesResponse struct {
	Items []USBControlSetName `json:"items"`
}
