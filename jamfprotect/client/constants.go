package client

// Base URL
const (
	// DefaultBaseURL is the default base URL for the Jamf Protect API
	DefaultBaseURL = "https://apis.jamfprotect.cloud"
)

// API Endpoints
const (
	// EndpointApp is the main GraphQL endpoint for full API access
	EndpointApp = "/app"
	
	// EndpointGraphQL is the limited schema GraphQL endpoint
	EndpointGraphQL = "/graphql"
	
	// EndpointToken is the OAuth2 token endpoint
	EndpointToken = "/token"
)

// User Agent
const (
	// UserAgentBase is the base name for the user agent
	UserAgentBase = "go-api-sdk-jamfprotect"
	
	// Version is the SDK version
	Version = "1.0.0"
	
	// DefaultUserAgent is the default user agent string
	DefaultUserAgent = UserAgentBase + "/" + Version
)

// Timeouts and Retries
const (
	// DefaultTimeout is the default HTTP client timeout in seconds
	DefaultTimeout = 60
	
	// MaxRetries is the maximum number of retries for failed requests
	MaxRetries = 3
	
	// RetryWaitTime is the wait time between retries in seconds
	RetryWaitTime = 2
	
	// RetryMaxWaitTime is the maximum wait time between retries in seconds
	RetryMaxWaitTime = 10
)

// Authentication
const (
	// TokenExpirySkew is the time buffer before token expiry to trigger refresh
	TokenExpirySkew = 60 // seconds
)

// HTTP Headers
const (
	// ContentTypeJSON is the JSON content type header value
	ContentTypeJSON = "application/json"
	
	// AcceptJSON is the JSON accept header value
	AcceptJSON = "application/json"
	
	// HeaderAuthorization is the Authorization header name
	HeaderAuthorization = "Authorization"
	
	// HeaderContentType is the Content-Type header name
	HeaderContentType = "Content-Type"
	
	// HeaderUserAgent is the User-Agent header name
	HeaderUserAgent = "User-Agent"
)
