package jamfprotect

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/actionconfigs"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytics"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analyticsets"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exceptionsets"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plans"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/preventlists"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/telemetryv2"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/unifiedloggingfilters"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/usbcontrolsets"
)

// Client is the main entry point for the Jamf Protect API SDK.
// It aggregates all service clients and provides a unified interface.
// Users should interact with the API exclusively through the provided service methods.
type Client struct {
	// transport is the internal HTTP transport layer (not exposed to users)
	transport *client.Transport

	// Services
	ActionConfigs         *actionconfigs.Service
	Plans                 *plans.Service
	Analytics             *analytics.Service
	AnalyticSets          *analyticsets.Service
	ExceptionSets         *exceptionsets.Service
	PreventLists          *preventlists.Service
	TelemetryV2           *telemetryv2.Service
	USBControlSets        *usbcontrolsets.Service
	UnifiedLoggingFilters *unifiedloggingfilters.Service
}

// NewClient creates a new Jamf Protect API client
//
// Parameters:
//   - clientID: The Jamf Protect OAuth2 client ID
//   - clientSecret: The Jamf Protect OAuth2 client secret
//   - options: Optional client configuration options
//
// Example:
//
//	client, err := jamfprotect.NewClient(
//	    "your-client-id",
//	    "your-client-secret",
//	    jamfprotect.WithDebug(),
//	)
func NewClient(clientID, clientSecret string, options ...client.ClientOption) (*Client, error) {
	transport, err := client.NewTransport(clientID, clientSecret, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP transport: %w", err)
	}

	// Initialize service clients
	c := &Client{
		transport:             transport,
		ActionConfigs:         actionconfigs.NewService(transport),
		Plans:                 plans.NewService(transport),
		Analytics:             analytics.NewService(transport),
		AnalyticSets:          analyticsets.NewService(transport),
		ExceptionSets:         exceptionsets.NewService(transport),
		PreventLists:          preventlists.NewService(transport),
		TelemetryV2:           telemetryv2.NewService(transport),
		USBControlSets:        usbcontrolsets.NewService(transport),
		UnifiedLoggingFilters: unifiedloggingfilters.NewService(transport),
	}

	return c, nil
}

// NewClientFromEnv creates a new client using environment variables
//
// Required environment variables:
//   - JAMFPROTECT_CLIENT_ID: The OAuth2 client ID
//   - JAMFPROTECT_CLIENT_SECRET: The OAuth2 client secret
//
// Optional environment variables:
//   - JAMFPROTECT_BASE_URL: Custom base URL (defaults to https://apis.jamfprotect.cloud)
//
// Example:
//
//	client, err := jamfprotect.NewClientFromEnv()
func NewClientFromEnv(options ...client.ClientOption) (*Client, error) {
	clientID := os.Getenv("JAMFPROTECT_CLIENT_ID")
	if clientID == "" {
		return nil, fmt.Errorf("JAMFPROTECT_CLIENT_ID environment variable is required")
	}

	clientSecret := os.Getenv("JAMFPROTECT_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, fmt.Errorf("JAMFPROTECT_CLIENT_SECRET environment variable is required")
	}

	// Check for optional environment variables and append to options
	if baseURL := os.Getenv("JAMFPROTECT_BASE_URL"); baseURL != "" {
		options = append(options, client.WithBaseURL(baseURL))
	}

	return NewClient(clientID, clientSecret, options...)
}

// GetLogger returns the configured zap logger instance.
// Use this to add custom logging within your application using the same logger.
//
// Returns:
//   - *zap.Logger: The configured logger instance
func (c *Client) GetLogger() *zap.Logger {
	return c.transport.GetLogger()
}

// GetTransport returns the underlying transport layer.
// This is useful for advanced configuration like setting custom loggers at runtime.
//
// Returns:
//   - *client.Transport: The transport instance
func (c *Client) GetTransport() *client.Transport {
	return c.transport
}
