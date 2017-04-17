package nexmo

import (
	"crypto/rsa"
	"math/rand"
	"strconv"
	"time"

	"fmt"

	"github.com/dghubble/sling"
	"github.com/dgrijalva/jwt-go"
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

type AuthSet struct {
	apiSecret *apiSecretAuth
	appAuth   *applicationAuth
}

func NewAuthSet() *AuthSet {
	return new(AuthSet)
}

func (a *AuthSet) SetAPISecret(apiKey, apiSecret string) {
	a.apiSecret = &apiSecretAuth{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (a *AuthSet) GenerateToken() (string, error) {
	if a.appAuth == nil {
		return "", fmt.Errorf("must call SetApplicationAuth before calling GenerateToken")
	}
	return a.appAuth.generateToken()
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

func (a *AuthSet) ApplyAuth(acceptableAuths []AuthType, sling *sling.Sling, request apiSecretRequest) error {
	for _, acceptableAuth := range acceptableAuths {
		switch acceptableAuth {
		case ApiSecretAuth:
			if a.apiSecret != nil {
				return a.apiSecret.apply(request)
			}
		case ApiSecretPathAuth:
			if a.apiSecret != nil {
				sling.Path(fmt.Sprintf("%s/%s", a.apiSecret.apiKey, a.apiSecret.apiSecret))
			}
		case JwtAuth:
			token, err := a.appAuth.generateToken()
			if err != nil {
				return err
			}
			if a.appAuth != nil {
				sling.Set("Authorization", fmt.Sprintf("Bearer %v", token))
			}
		}
	}
	return fmt.Errorf("AuthSet not configured with credentials for %x", acceptableAuths)
}

type apiSecretRequest interface {
	applyAPISecret(apiKey, apiSecret string)
}

type apiSecretAuth struct {
	apiKey    string
	apiSecret string
}

func (a apiSecretAuth) apply(request apiSecretRequest) error {
	if request == nil {
		return fmt.Errorf("cannot apply api_key and api_secret to a nil request")
	}
	request.applyAPISecret(a.apiKey, a.apiSecret)
	return nil
}

type applicationAuth struct {
	appID      string
	privateKey *rsa.PrivateKey
	r          RandomProvider
}

type jwtClaims struct {
	ApplicationID string `json:"application_id"`
	jwt.StandardClaims
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
