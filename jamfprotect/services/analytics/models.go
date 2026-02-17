package analytics

// Analytic represents a Jamf Protect analytic
type Analytic struct {
	UUID            string              `json:"uuid"`
	Name            string              `json:"name"`
	Label           string              `json:"label"`
	InputType       string              `json:"inputType"`
	Filter          string              `json:"filter"`
	Description     string              `json:"description"`
	LongDescription string              `json:"longDescription"`
	Created         string              `json:"created"`
	Updated         string              `json:"updated"`
	Actions         []string            `json:"actions"`
	AnalyticActions []AnalyticAction    `json:"analyticActions"`
	TenantActions   []AnalyticAction    `json:"tenantActions"`
	Tags            []string            `json:"tags"`
	Level           int                 `json:"level"`
	Severity        string              `json:"severity"`
	TenantSeverity  string              `json:"tenantSeverity"`
	SnapshotFiles   []string            `json:"snapshotFiles"`
	Context         []AnalyticContext   `json:"context"`
	Categories      []string            `json:"categories"`
	Jamf            bool                `json:"jamf"`
	Remediation     string              `json:"remediation"`
}

// AnalyticAction represents an action configuration for an analytic
type AnalyticAction struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
}

// AnalyticContext represents context configuration for an analytic
type AnalyticContext struct {
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	Exprs []string `json:"exprs"`
}

// CreateAnalyticRequest is the request payload for creating an analytic
type CreateAnalyticRequest struct {
	Name            string
	InputType       string
	Description     string
	Actions         []string
	AnalyticActions []AnalyticActionInput
	Tags            []string
	Categories      []string
	Filter          string
	Context         []AnalyticContextInput
	Level           int
	Severity        string
	SnapshotFiles   []string
}

// UpdateAnalyticRequest is the request payload for updating an analytic
type UpdateAnalyticRequest struct {
	Name            string
	InputType       string
	Description     string
	Actions         []string
	AnalyticActions []AnalyticActionInput
	Tags            []string
	Categories      []string
	Filter          string
	Context         []AnalyticContextInput
	Level           int
	Severity        *string
	SnapshotFiles   []string
}

// AnalyticActionInput represents an action input for create/update
type AnalyticActionInput struct {
	Name       string
	Parameters []string
}

// AnalyticContextInput represents a context input for create/update
type AnalyticContextInput struct {
	Name  string
	Type  string
	Exprs []string
}

// ListAnalyticsResponse represents the response from listing analytics
type ListAnalyticsResponse struct {
	Items    []Analytic `json:"items"`
	PageInfo PageInfo   `json:"pageInfo"`
}

// PageInfo contains pagination information
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
