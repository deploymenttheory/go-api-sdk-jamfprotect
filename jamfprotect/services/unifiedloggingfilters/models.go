package unifiedloggingfilters

// UnifiedLoggingFilter represents a Jamf Protect unified logging filter
type UnifiedLoggingFilter struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Created     string   `json:"created"`
	Updated     string   `json:"updated"`
	Filter      string   `json:"filter"`
	Tags        []string `json:"tags"`
	Enabled     bool     `json:"enabled"`
}

// CreateUnifiedLoggingFilterRequest is the request payload for creating a unified logging filter
type CreateUnifiedLoggingFilterRequest struct {
	Name        string
	Description string
	Tags        []string
	Filter      string
	Enabled     bool
}

// UpdateUnifiedLoggingFilterRequest is the request payload for updating a unified logging filter
type UpdateUnifiedLoggingFilterRequest struct {
	Name        string
	Description string
	Tags        []string
	Filter      string
	Enabled     bool
}

// ListUnifiedLoggingFiltersResponse represents the response from listing unified logging filters
type ListUnifiedLoggingFiltersResponse struct {
	Items    []UnifiedLoggingFilter `json:"items"`
	PageInfo PageInfo               `json:"pageInfo"`
}

// PageInfo contains pagination information
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
