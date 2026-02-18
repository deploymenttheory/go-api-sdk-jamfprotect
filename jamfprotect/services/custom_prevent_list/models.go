package custompreventlist

// PreventList represents a Jamf Protect prevent list
type PreventList struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	List        []string `json:"list"`
	Count       int64    `json:"count"`
	Created     string   `json:"created"`
}

// CreatePreventListRequest is the request payload for creating a prevent list
type CreatePreventListRequest struct {
	Name        string
	Description string
	Type        string
	Tags        []string
	List        []string
}

// UpdatePreventListRequest is the request payload for updating a prevent list
type UpdatePreventListRequest struct {
	Name        string
	Description string
	Type        string
	Tags        []string
	List        []string
}

// ListPreventListsResponse represents the response from listing prevent lists
type ListPreventListsResponse struct {
	Items    []PreventList `json:"items"`
	PageInfo PageInfo      `json:"pageInfo"`
}

// PageInfo contains pagination information
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}

// PreventListName is a lightweight prevent list containing only the name
type PreventListName struct {
	Name string `json:"name"`
}

// ListPreventListNamesResponse is the response wrapper for listing prevent list names
type ListPreventListNamesResponse struct {
	Items []PreventListName `json:"items"`
}
