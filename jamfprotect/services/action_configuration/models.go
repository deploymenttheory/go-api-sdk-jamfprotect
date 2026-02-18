package actionconfiguration

// ActionConfig represents a Jamf Protect action configuration.
type ActionConfig struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Hash        string         `json:"hash"`
	Created     string         `json:"created"`
	Updated     string         `json:"updated"`
	AlertConfig *AlertConfig   `json:"alertConfig"`
	Clients     []ReportClient `json:"clients"`
}

// ActionConfigListItem is the list view for action configurations.
type ActionConfigListItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}

// CreateActionConfigRequest is the request payload for creating an action configuration.
type CreateActionConfigRequest struct {
	Name        string
	Description string
	AlertConfig map[string]any
	Clients     []map[string]any
}

// UpdateActionConfigRequest is the request payload for updating an action configuration.
type UpdateActionConfigRequest struct {
	Name        string
	Description string
	AlertConfig map[string]any
	Clients     []map[string]any
}

// AlertConfig maps alert configuration data for action configs.
type AlertConfig struct {
	Data *AlertData `json:"data"`
}

// AlertData contains event-type alert enrichment configuration.
type AlertData struct {
	Binary              *AlertEventType `json:"binary"`
	ClickEvent          *AlertEventType `json:"clickEvent"`
	DownloadEvent       *AlertEventType `json:"downloadEvent"`
	File                *AlertEventType `json:"file"`
	FsEvent             *AlertEventType `json:"fsEvent"`
	Group               *AlertEventType `json:"group"`
	ProcEvent           *AlertEventType `json:"procEvent"`
	Process             *AlertEventType `json:"process"`
	ScreenshotEvent     *AlertEventType `json:"screenshotEvent"`
	UsbEvent            *AlertEventType `json:"usbEvent"`
	User                *AlertEventType `json:"user"`
	GkEvent             *AlertEventType `json:"gkEvent"`
	KeylogRegisterEvent *AlertEventType `json:"keylogRegisterEvent"`
	MrtEvent            *AlertEventType `json:"mrtEvent"`
}

// AlertEventType describes an event type's included attributes and related objects.
type AlertEventType struct {
	Attrs   []string `json:"attrs"`
	Related []string `json:"related"`
}

// ReportClient represents a reporting client configuration.
type ReportClient struct {
	ID               string              `json:"id"`
	Type             string              `json:"type"`
	SupportedReports []string            `json:"supportedReports"`
	BatchConfig      *BatchConfig        `json:"batchConfig"`
	Params           *ReportClientParams `json:"params"`
}

// BatchConfig represents batching configuration for a report client.
type BatchConfig struct {
	Delimiter       string `json:"delimiter"`
	SizeIndex       int64  `json:"sizeIndex"`
	WindowInSeconds int64  `json:"windowInSeconds"`
	SizeInBytes     int64  `json:"sizeInBytes"`
}

// ReportClientParams represents endpoint-specific parameters for a report client.
// Different client types populate different fields (JamfCloud, HTTP, Kafka, Syslog, LogFile).
type ReportClientParams struct {
	DestinationFilter string               `json:"destinationFilter"`
	Headers           []ReportClientHeader `json:"headers"`
	Method            string               `json:"method"`
	URL               string               `json:"url"`
	Host              string               `json:"host"`
	Port              int64                `json:"port"`
	Topic             string               `json:"topic"`
	ClientCN          string               `json:"clientCN"`
	ServerCN          string               `json:"serverCN"`
	Scheme            string               `json:"scheme"`
	Path              string               `json:"path"`
	Permissions       string               `json:"permissions"`
	MaxSizeMB         int64                `json:"maxSizeMB"`
	Ownership         string               `json:"ownership"`
	Backups           int64                `json:"backups"`
}

// ReportClientHeader represents a header entry for HTTP-based clients.
type ReportClientHeader struct {
	Header string `json:"header"`
	Value  string `json:"value"`
}

// ListActionConfigsResponse represents the response from listing action configurations.
type ListActionConfigsResponse struct {
	Items    []ActionConfigListItem `json:"items"`
	PageInfo PageInfo               `json:"pageInfo"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}

// ActionConfigName is a lightweight action configuration containing only the name
type ActionConfigName struct {
	Name string `json:"name"`
}

// ListActionConfigNamesResponse is the response wrapper for listing action configuration names
type ListActionConfigNamesResponse struct {
	Items []ActionConfigName `json:"items"`
}
