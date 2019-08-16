package nexmo

import (
	"crypto/rsa"
	"math/rand"
	"strconv"
	"time"

	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/nexmo-community/nexmo-go/sling"
)

type AuthType uint8

type RandomProvider interface {
	Int31() int32
}

const (
	ApiSecretAuth AuthType = iota + 1
	ApiSecretPathAuth
	JwtAuth
)

// API credentials to access the Nexmo APIs
type AuthSet struct {
	apiSecret *apiSecretAuth
	appAuth   *applicationAuth
}

func NewAuthSet() *AuthSet {
	return new(AuthSet)
}

func (a *AuthSet) ApplyAPICredentials(request apiSecretRequest) {
	a.apiSecret.apply(request)
}

func (a *AuthSet) ApplyJWT(sling *sling.Sling) error {
	token, err := a.GenerateToken()
	if err != nil {
		return err
	}
	sling.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return nil
}

func (a *AuthSet) GenerateToken() (string, error) {
	if a.appAuth == nil {
		return "", fmt.Errorf("must call SetApplicationAuth before calling GenerateToken")
	}
	return a.appAuth.generateToken()
}

func (a *AuthSet) SetAPISecret(apiKey, apiSecret string) {
	a.apiSecret = &apiSecretAuth{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (a *AuthSet) SetApplicationAuth(appID string, key []byte) error {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return err
	}
	a.appAuth = &applicationAuth{
		appID:      appID,
		privateKey: privateKey,
		r:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return nil
}

type apiSecretRequest interface {
	setApiCredentials(apiKey, apiSecret string)
}

type apiSecretAuth struct {
	apiKey    string
	apiSecret string
}

func (a apiSecretAuth) apply(request apiSecretRequest) error {
	request.setApiCredentials(a.apiKey, a.apiSecret)
	return nil
}

type jwtClaims struct {
	ApplicationID string `json:"application_id"`
	jwt.StandardClaims
}

type applicationAuth struct {
	appID      string
	privateKey *rsa.PrivateKey
	r          RandomProvider
}

func (a applicationAuth) generateToken() (string, error) {
	claims := jwtClaims{
		a.appID,
		jwt.StandardClaims{
			Id:       strconv.Itoa(int(a.r.Int31())),
			IssuedAt: time.Now().UTC().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	return token.SignedString(a.privateKey)
}

type Credentials struct {
	APIKey    string `json:"api_key" url:"api_key"`
	APISecret string `json:"api_secret" url:"api_secret"`
}

func (c *Credentials) setApiCredentials(apiKey, apiSecret string) {
	c.APIKey = apiKey
	c.APISecret = apiSecret
}
