package exceptionset

// ExceptionSet represents a Jamf Protect exception set
type ExceptionSet struct {
	UUID         string        `json:"uuid"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Exceptions   []Exception   `json:"exceptions"`
	EsExceptions []EsException `json:"esExceptions"`
	Created      string        `json:"created"`
	Updated      string        `json:"updated"`
	Managed      bool          `json:"managed"`
}

// Exception represents an exception entry
type Exception struct {
	Type           string          `json:"type"`
	Value          string          `json:"value"`
	AppSigningInfo *AppSigningInfo `json:"appSigningInfo"`
	IgnoreActivity string          `json:"ignoreActivity"`
	AnalyticTypes  []string        `json:"analyticTypes"`
	AnalyticUuid   string          `json:"analyticUuid"`
	Analytic       *AnalyticRef    `json:"analytic"`
}

// AnalyticRef represents an analytic reference on an exception
type AnalyticRef struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// EsException represents an ES exception entry
type EsException struct {
	Type              string          `json:"type"`
	Value             string          `json:"value"`
	AppSigningInfo    *AppSigningInfo `json:"appSigningInfo"`
	IgnoreActivity    string          `json:"ignoreActivity"`
	IgnoreListType    string          `json:"ignoreListType"`
	IgnoreListSubType string          `json:"ignoreListSubType"`
	EventType         string          `json:"eventType"`
}

// AppSigningInfo represents app signing info in responses
type AppSigningInfo struct {
	AppId  string `json:"appId"`
	TeamId string `json:"teamId"`
}

// CreateExceptionSetRequest is the request payload for creating an exception set
type CreateExceptionSetRequest struct {
	Name         string
	Description  string
	Exceptions   []ExceptionInput
	EsExceptions []EsExceptionInput
}

// UpdateExceptionSetRequest is the request payload for updating an exception set
type UpdateExceptionSetRequest struct {
	Name         string
	Description  string
	Exceptions   []ExceptionInput
	EsExceptions []EsExceptionInput
}

// ExceptionInput represents an exception entry input
type ExceptionInput struct {
	Type           string
	Value          string
	AppSigningInfo *AppSigningInfoInput
	IgnoreActivity string
	AnalyticTypes  []string
	AnalyticUuid   string
}

// AppSigningInfoInput represents app signing info in input
type AppSigningInfoInput struct {
	AppId  string
	TeamId string
}

// EsExceptionInput represents an ES exception entry input
type EsExceptionInput struct {
	Type              string
	Value             string
	AppSigningInfo    *AppSigningInfoInput
	IgnoreActivity    string
	IgnoreListType    string
	IgnoreListSubType string
	EventType         string
}

// ListExceptionSetsResponse represents the response from listing exception sets
type ListExceptionSetsResponse struct {
	Items    []ExceptionSetListItem `json:"items"`
	PageInfo PageInfo               `json:"pageInfo"`
}

// ExceptionSetListItem represents a list item for exception sets
type ExceptionSetListItem struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Managed bool   `json:"managed"`
}

// PageInfo contains pagination information
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
