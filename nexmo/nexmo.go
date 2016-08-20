package nexmo

import (
	"net/http"
	"time"
)

// DefaultBaseURL is the service url used unless overridden.
const DefaultBaseURL = "https://api.nexmo.com"

// defaultTimeout of 5000ms used by this client unless specifically overridden on the http.Client.
const defaultTimeout = 5000 * time.Millisecond

// Client represents a connection to the Nexmo API.
type Client interface {
	SetBaseURL(baseURL string)
	Verify(VerifyRequest) (*VerifyResponse, error)
	Check(requestID, code string) (*CheckResponse, error)
}

// NewClient creates a Client for sending requests to Nexmo.
func NewClient(apiKey, apiSecret string, client *http.Client) Client {
	if client == nil {
		client = &http.Client{
			Timeout: defaultTimeout,
		}
	}

	return &nexmoClient{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		baseURL:    DefaultBaseURL,
		httpClient: client,
	}
}
