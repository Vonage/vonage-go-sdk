// This package offers a simple API wrapper and helper functions to get
// users started with the Nexmo APIs
// Pull requests, issues and comments are all welcome and gratefully received

package nexmo

// All the various Auth types support a common interface
type Auth interface {
	getCreds() []string
}

// Auth to represent the API key and API secret combination
type KeySecretAuth struct {
	apiKey    string
	apiSecret string
}

func (auth *KeySecretAuth) getCreds() []string {
	creds := []string{auth.apiKey, auth.apiSecret}
	return creds
}

// Create an Auth struct by supplying an API key and API secret
func CreateAuthFromKeySecret(apiKey string, apiSecret string) *KeySecretAuth {
	auth := new(KeySecretAuth)
	auth.apiKey = apiKey
	auth.apiSecret = apiSecret
	return auth
}
