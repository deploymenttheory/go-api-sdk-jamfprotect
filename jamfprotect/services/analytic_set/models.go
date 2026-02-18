package analyticset

// AnalyticSet represents a Jamf Protect analytic set
type AnalyticSet struct {
	UUID        string                `json:"uuid"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Types       []string              `json:"types"`
	Analytics   []AnalyticSetAnalytic `json:"analytics"`
	Plans       []AnalyticSetPlan     `json:"plans"`
	Created     string                `json:"created"`
	Updated     string                `json:"updated"`
	Managed     bool                  `json:"managed"`
}

// AnalyticSetAnalytic represents an analytic entry in a set
type AnalyticSetAnalytic struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Jamf bool   `json:"jamf"`
}

// AnalyticSetPlan represents a plan entry in a set
type AnalyticSetPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateAnalyticSetRequest is the request payload for creating an analytic set
type CreateAnalyticSetRequest struct {
	Name        string
	Description string
	Types       []string
	Analytics   []string
}

// UpdateAnalyticSetRequest is the request payload for updating an analytic set
type UpdateAnalyticSetRequest struct {
	Name        string
	Description string
	Types       []string
	Analytics   []string
}

// ListAnalyticSetsResponse represents the response from listing analytic sets
type ListAnalyticSetsResponse struct {
	Items    []AnalyticSet `json:"items"`
	PageInfo PageInfo      `json:"pageInfo"`
}

// PageInfo contains pagination information
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
