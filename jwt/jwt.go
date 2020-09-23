package jwt

import (
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Path represents each path in the ACL structure
type Path struct {
	Path string
}

// Generator is what makes a token. Set the fields you need, then generate. The token is also stored in `token` once generated.
type Generator struct {
	ApplicationID string
	PrivateKey    []byte
	TTL           time.Duration
	Subject       string
	Paths         []Path
	JTI           uuid.UUID
	NBF           int64
	token         *jwt.Token
}

// NewGenerator takes your application ID and private key to create a generator
func NewGenerator(ApplicationID string, PrivateKey []byte) *Generator {
	g := new(Generator)
	g.ApplicationID = ApplicationID
	g.PrivateKey = PrivateKey
	return g
}

// NewGeneratorFromFilename takes your application ID and the filename of your private key to create a token generator
func NewGeneratorFromFilename(ApplicationID string, PrivateKeyFileName string) (*Generator, error) {
	key, err := ioutil.ReadFile(PrivateKeyFileName)
	if err != nil {
		return &Generator{}, err
	}

	g := NewGenerator(ApplicationID, key)
	return g, nil
}

// AddPath adds an entry to the ACL "paths" field
func (g *Generator) AddPath(path Path) *Generator {
	g.Paths = append(g.Paths, path)
	return g
}

// GenerateToken assembles and returns the JWT token
func (g *Generator) GenerateToken() (string, error) {
	// build the JWT up: First: non-editable fields
	atClaims := jwt.MapClaims{}
	atClaims["iat"] = time.Now().Unix()

	// fields set on the generator
	atClaims["application_id"] = g.ApplicationID

	ttl := time.Minute * time.Duration(15) // recommended token lifetime 15 min
	if g.TTL != 0 {
		ttl = g.TTL
	}
	atClaims["exp"] = time.Now().Add(ttl).Unix()

	jti := uuid.New()
	if g.JTI != uuid.Nil {
		jti = g.JTI
	}
	atClaims["jti"] = jti

	if g.NBF != 0 {
		atClaims["nbf"] = g.NBF
	}

	if g.Subject != "" {
		atClaims["sub"] = g.Subject
	}

	if len(g.Paths) > 0 {
		// add these paths to ACL
		atClaims["acl"] = g.getACL()
	}

	signWith, keyErr := jwt.ParseRSAPrivateKeyFromPEM(g.PrivateKey)
	if keyErr != nil {
		return "", keyErr
	}

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	token, tokenErr := at.SignedString(signWith)
	if tokenErr != nil {
		return "", tokenErr
	}

	// store the token and return a string
	g.token = at
	return token, nil
}

// getACL is a helper function to build the ACL claim from our Paths structs
func (g *Generator) getACL() map[string]map[string]map[string]string {
	// aiming for {"paths": ["thing": {}, "otherThing": {}]}
	mymap := map[string]map[string]map[string]string{}
	mymap["paths"] = make(map[string]map[string]string)

	for _, path := range g.Paths {
		mymap["paths"][path.Path] = make(map[string]string)
	}

	return mymap
}

// GetHeader gives access to the header fields `alg` and `typ` of the generated token
func (g *Generator) GetHeader() map[string]interface{} {
	return g.token.Header
}

// GetClaims returns the body claims for the generated token
func (g *Generator) GetClaims() jwt.MapClaims {
	claims := g.token.Claims

	// cast the interface to our map type
	return claims.(jwt.MapClaims)
}
