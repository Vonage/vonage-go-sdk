---
title: JWTs
permalink: examples/jwt
---

* [Generate a Basic JWT](#generate-a-basic-jwt)
* [Generate a JWT with more options](#generate-a-jwt-with-more-options)

# JWT Authentication

We use JSON Web Tokens for authentication on some APIs (some are API key and secret). More information about working with JWTs is in the following sections.

## Generate a Basic JWT

Generate a JSON Web Token (JWT) for the APIs that use that. You usually won't need to do this if you're using the library but if you need to make a custom request or want to use a JWT for something else, you can use this.

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk/jwt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
    g := jwt.NewGenerator(APPLICATION_ID, privateKey)

    token, _ := g.GenerateToken()
    fmt.Println(token)
}
```

## Generate a JWT with more options

You can also set up the generator with the options needed on your token, such as expiry time or ACLs.

```go
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
    g := jwt.Generator{
        ApplicationID: APPLICATION_ID,
        PrivateKey:    privateKey,
        TTL:           time.Minute * time.Duration(90),
    }
	g.AddPath(jwt.Path{Path: "/*/users/**"})

    token, _ := g.GenerateToken()
    fmt.Println(token)
```



