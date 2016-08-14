package nexmo

import (
	"time"
)

// DefaultBaseURL is the service url used unless overridden.
const DefaultBaseURL = "https://api.nexmo.com"

// DefaultConnectionTimeout of 5000ms used by this client unless specifically overridden onb the constructor
const DefaultConnectionTimeout = 5000 * time.Millisecond

// DefaultSoTimeout is the read timeout of 30000ms used by this client unless specifically overridden onb the constructor
const DefaultSoTimeout = 30000 * time.Millisecond

// PathVerify is the endpoint path for submitting verification requests
const pathVerify = "/verify/json"

// PathVerifyCheck is the endpoint path for submitting verification check requests
const pathVerifyCheck = "/verify/check/json"

// PathVerifySearch is the endpoint path for submitting verification search requests
const pathVerifySearch = "/verify/search/json"

// Client represents a connection to the Nexmo API.
type Client interface {
	SetBaseURL(baseURL string)
	SetConnectionTimeout(timeout time.Duration)
	SetSoTimeout(timeout time.Duration)
	Verify(VerifyRequest) (*VerifyResponse, error)
	Check(requestID, code string) (*CheckResponse, error)
}

// NewClient creates a Client for sending requests to Nexmo.
func NewClient(apiKey, apiSecret string) Client {
	return &nexmoClient{
		apiKey:            apiKey,
		apiSecret:         apiSecret,
		baseURL:           DefaultBaseURL,
		connectionTimeout: DefaultConnectionTimeout,
		soTimeout:         DefaultSoTimeout,
	}
}
