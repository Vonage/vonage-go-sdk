// This package offers a simple API wrapper and helper functions to get
// users started with the Nexmo APIs
// Pull requests, issues and comments are all welcome and gratefully received

package nexmo

// Auth types are various but support a common interface
type Auth interface {
	GetCreds() []string
}

// KeySecretAuth is an Auth type to represent the API key and API secret combination
type KeySecretAuth struct {
	apiKey    string
	apiSecret string
}

// GetCreds gives an array of credential strings
func (auth *KeySecretAuth) GetCreds() []string {
	creds := []string{auth.apiKey, auth.apiSecret}
	return creds
}

// CreateAuthFromKeySecret returns an Auth type given an API key and API secret
func CreateAuthFromKeySecret(apiKey string, apiSecret string) *KeySecretAuth {
	auth := new(KeySecretAuth)
	auth.apiKey = apiKey
	auth.apiSecret = apiSecret
	return auth
}
