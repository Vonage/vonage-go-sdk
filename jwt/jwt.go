package jwt

import (
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Path struct {
	Path string
}

type Generator struct {
	ApplicationID string
	PrivateKey    []byte
	TTL           int
	Subject       string
	Paths         []Path
	JTI           string
	NBF           int64
	token         *jwt.Token
}

func NewGenerator(ApplicationID string, PrivateKey []byte) *Generator {
	g := new(Generator)
	g.ApplicationID = ApplicationID
	g.PrivateKey = PrivateKey
	return g
}

func NewGeneratorFromFilename(ApplicationID string, PrivateKeyFileName string) (*Generator, error) {
	key, err := ioutil.ReadFile(PrivateKeyFileName)
	if err != nil {
		return &Generator{}, err
	}

	g := NewGenerator(ApplicationID, key)
	return g, nil
}

func (g *Generator) AddPath(path Path) *Generator {
	g.Paths = append(g.Paths, path)
	return g
}

func (g *Generator) SetPaths(paths []Path) *Generator {
	g.Paths = paths
	return g
}

func (g *Generator) SetTTL(ttl int) *Generator {
	g.TTL = ttl
	return g
}

func (g *Generator) SetJTI(jti string) *Generator {
	g.JTI = jti
	return g
}

func (g *Generator) SetNBF(nbf int64) *Generator {
	g.NBF = nbf
	return g
}

func (g *Generator) SetSubject(subject string) *Generator {
	g.Subject = subject
	return g
}

func (g *Generator) GenerateToken() (string, error) {
	// build the JWT up: First: non-editable fields
	atClaims := jwt.MapClaims{}
	atClaims["iat"] = time.Now().Unix()

	// fields set on the generator
	atClaims["application_id"] = g.ApplicationID

	ttl := 15 // recommended token lifetime, in minutes
	if g.TTL != 0 {
		ttl = g.TTL
	}
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(ttl)).Unix()

	initJTI := uuid.New()
	atClaims["jti"] = initJTI
	if g.JTI != "" {
		jti, jtiErr := uuid.Parse(g.JTI)
		if jtiErr != nil {
			jti = uuid.New()
		}
		atClaims["jti"] = jti
	}

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

func (g *Generator) getACL() map[string]map[string]map[string]string {
	// aiming for {"paths": ["thing": {}, "otherThing": {}]}
	mymap := map[string]map[string]map[string]string{}
	mymap["paths"] = make(map[string]map[string]string)

	for _, path := range g.Paths {
		mymap["paths"][path.Path] = make(map[string]string)
	}

	return mymap
}

func (g *Generator) GetHeader() map[string]interface{} {
	return g.token.Header
}

func (g *Generator) GetClaims() jwt.MapClaims {
	claims := g.token.Claims

	// cast the interface to our map type
	return claims.(jwt.MapClaims)
}
