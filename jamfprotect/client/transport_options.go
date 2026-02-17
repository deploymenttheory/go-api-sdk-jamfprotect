package client

import (
	"crypto/tls"
	"fmt"
	"maps"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// ClientOption is a function type for configuring the Client
type ClientOption func(*Transport) error

// WithBaseURL sets a custom base URL for the API client
func WithBaseURL(baseURL string) ClientOption {
	return func(t *Transport) error {
		t.baseURL = baseURL
		t.client.SetBaseURL(baseURL)
		t.logger.Info("Base URL configured", zap.String("base_url", baseURL))
		return nil
	}
}

// WithTimeout sets a custom timeout for HTTP requests
func WithTimeout(timeout time.Duration) ClientOption {
	return func(t *Transport) error {
		t.client.SetTimeout(timeout)
		t.logger.Info("HTTP timeout configured", zap.Duration("timeout", timeout))
		return nil
	}
}

// WithRetryCount sets the number of retries for failed requests
func WithRetryCount(count int) ClientOption {
	return func(t *Transport) error {
		t.client.SetRetryCount(count)
		t.logger.Info("Retry count configured", zap.Int("retry_count", count))
		return nil
	}
}

// WithRetryWaitTime sets the default wait time between retry attempts
// This is the initial/minimum wait time before the first retry
func WithRetryWaitTime(waitTime time.Duration) ClientOption {
	return func(t *Transport) error {
		t.client.SetRetryWaitTime(waitTime)
		t.logger.Info("Retry wait time configured", zap.Duration("wait_time", waitTime))
		return nil
	}
}

// WithRetryMaxWaitTime sets the maximum wait time between retry attempts
// The wait time increases exponentially with each retry up to this maximum
func WithRetryMaxWaitTime(maxWaitTime time.Duration) ClientOption {
	return func(t *Transport) error {
		t.client.SetRetryMaxWaitTime(maxWaitTime)
		t.logger.Info("Retry max wait time configured", zap.Duration("max_wait_time", maxWaitTime))
		return nil
	}
}

// WithLogger sets a custom logger for the client
func WithLogger(logger *zap.Logger) ClientOption {
	return func(t *Transport) error {
		t.logger = logger
		t.logger.Info("Custom logger configured")
		return nil
	}
}

// WithDebug enables debug mode which logs request and response details
func WithDebug() ClientOption {
	return func(t *Transport) error {
		t.client.SetDebug(true)
		t.logger.Info("Debug mode enabled")
		return nil
	}
}

// WithUserAgent sets a custom user agent string
func WithUserAgent(userAgent string) ClientOption {
	return func(t *Transport) error {
		t.client.SetHeader(HeaderUserAgent, userAgent)
		t.userAgent = userAgent
		t.logger.Info("User agent configured", zap.String("user_agent", userAgent))
		return nil
	}
}

// WithCustomAgent allows appending a custom identifier to the default user agent
// Format: "go-api-sdk-jamfprotect/1.0.0; <customAgent>; gzip"
func WithCustomAgent(customAgent string) ClientOption {
	return func(t *Transport) error {
		enhancedUA := fmt.Sprintf("%s/%s; %s; gzip", UserAgentBase, Version, customAgent)
		t.client.SetHeader(HeaderUserAgent, enhancedUA)
		t.userAgent = enhancedUA
		t.logger.Info("Custom agent configured", zap.String("user_agent", enhancedUA))
		return nil
	}
}

// WithGlobalHeader sets a global header that will be included in all requests
// Per-request headers will override global headers with the same key
func WithGlobalHeader(key, value string) ClientOption {
	return func(t *Transport) error {
		t.globalHeaders[key] = value
		t.logger.Info("Global header configured", zap.String("key", key), zap.String("value", value))
		return nil
	}
}

// WithGlobalHeaders sets multiple global headers at once
func WithGlobalHeaders(headers map[string]string) ClientOption {
	return func(t *Transport) error {
		maps.Copy(t.globalHeaders, headers)
		t.logger.Info("Multiple global headers configured", zap.Int("count", len(headers)))
		return nil
	}
}

// WithProxy sets an HTTP proxy for all requests
// Example: "http://proxy.company.com:8080" or "socks5://127.0.0.1:1080"
func WithProxy(proxyURL string) ClientOption {
	return func(t *Transport) error {
		t.client.SetProxy(proxyURL)
		t.logger.Info("Proxy configured", zap.String("proxy", proxyURL))
		return nil
	}
}

// WithTLSConfig sets a custom TLS configuration
func WithTLSConfig(tlsConfig *tls.Config) ClientOption {
	return func(t *Transport) error {
		t.client.SetTLSClientConfig(tlsConfig)
		t.logger.Info("TLS configuration set")
		return nil
	}
}

// WithTLSInsecureSkipVerify disables SSL certificate verification
// WARNING: Only use this for testing. Never in production!
func WithTLSInsecureSkipVerify() ClientOption {
	return func(t *Transport) error {
		t.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		t.logger.Warn("TLS certificate verification disabled - USE ONLY FOR TESTING")
		return nil
	}
}

// WithTransport sets a custom HTTP transport
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(t *Transport) error {
		httpClient := t.client.Client()
		if httpClient != nil {
			httpClient.Transport = transport
			t.logger.Info("Custom HTTP transport configured")
		}
		return nil
	}
}

// WithRateLimiter sets a custom rate limiter function
// The function is called before each request and can return an error to rate limit
func WithRateLimiter(limiter func() error) ClientOption {
	return func(t *Transport) error {
		// Wrap the HTTP client transport with rate limiting logic
		httpClient := t.client.Client()
		if httpClient != nil {
			originalTransport := httpClient.Transport
			if originalTransport == nil {
				originalTransport = http.DefaultTransport
			}
			httpClient.Transport = &rateLimitTransport{
				base:    originalTransport,
				limiter: limiter,
			}
		}
		t.logger.Info("Rate limiter configured")
		return nil
	}
}

// rateLimitTransport wraps an HTTP transport with rate limiting
type rateLimitTransport struct {
	base    http.RoundTripper
	limiter func() error
}

func (r *rateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := r.limiter(); err != nil {
		return nil, err
	}
	return r.base.RoundTrip(req)
}
