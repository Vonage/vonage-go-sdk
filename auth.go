// This package offers a simple API wrapper and helper functions to get
// users started with the Nexmo APIs
// Pull requests, issues and comments are all welcome and gratefully received

package nexmo

import (
	"github.com/nexmo-community/nexmo-go/jwt"
)

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

// JWTAuth is an Auth type to represent a JWT token
type JWTAuth struct {
	JWT string
}

func (auth *JWTAuth) GetCreds() []string {
	creds := []string{auth.JWT}
	return creds
}

func CreateAuthFromAppPrivateKey(appID string, privateKey []byte) (*JWTAuth, error) {
	jwtGen := jwt.NewGenerator(appID, privateKey)
	token, tokenErr := jwtGen.GenerateToken()
	if tokenErr != nil {
		return &JWTAuth{}, tokenErr
	}

	auth := new(JWTAuth)
	auth.JWT = token
	return auth, nil
}
