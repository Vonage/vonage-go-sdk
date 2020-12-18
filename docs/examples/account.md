---
title: Account API
permalink: examples/account
---

Account API gives access to check your balance, rotate your account secrets and configure the default behaviour of your account. Check out the [documentation](https://developer.nexmo.com/account/overview) and [API reference](https://developer.nexmo.com/api/account) for more details.

- [Get Account Balance](#get-account-balance)
- [Configure Account](#configure-account)
- [Fetch All Account Secrets](#fetch-all-account-secrets)
- [Fetch One Account Secret](#fetch-one-account-secret)
- [Create Account Secret](#create-account-secret)
- [Revoke Account Secret](#revoke-account-secret)

## Get Account Balance

Check the current balance on your account.

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
    auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	accountClient := vonage.NewAccountClient(auth)
	response, _, err := accountClient.GetBalance()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Account balance: %f EUR", response.Value)
}
```

## Configure Account

Set the default URLs for incoming SMS and delivery receipt payloads to be sent to. If the number has settings, that will be used but otherwise we fall back to this account API setting.

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
    auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	accountClient := vonage.NewAccountClient(auth)

    opts := vonage.AccountConfigSettings{
        MoCallbackUrl: "https://example.com/webhooks/inbound-sms",
        DrCallbackUrl: "https://example.com/webhooks/delivery-receipt"
    }
	response, _, err := accountClient.SetConfig(opts)

	if err != nil {
		panic(err)
	}

    fmt.Println("Incoming SMS sent to: " + response.MoCallbackUrl)
}
```

## Fetch All Account Secrets

This endpoint returns the secret ID (needed to identify the secret for deletion) and creation date for all the secrets on this account.

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
    auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	accountClient := vonage.NewAccountClient(auth)

	response, _, err := accountClient.ListSecrets()

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(response.Secrets); i++ {
		date := response.Secrets[i].CreatedAt.Format("Jan 2 2006")
		fmt.Println("Secret " + response.Secrets[i].ID + " created " + date)
	}
}
```

## Fetch One Account Secret

This endpoint returns the secret ID and creation date for the secret requested.

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
    auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	accountClient := vonage.NewAccountClient(auth)

	response, _, err := accountClient.GetSecret("abcdefab-0000-1111-2222-0123456789ef")

	if err != nil {
		panic(err)
	}

	if response.ID != "" {
		date := response.CreatedAt.Format("Jan 2 2006")
		fmt.Println("Secret " + response.ID + " created " + date)
	}
}
```

## Create Account Secret

To change account secrets, first add a new secret. Then update your application and when it's using the new secret, revoke the old secret.

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
    auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	accountClient := vonage.NewAccountClient(auth)

	response, errResp, err := accountClient.CreateSecret("T0pS3cr3t!")

	if err != nil {
		fmt.Println("ERROR: " + errResp.Title + ": " + errResp.Detail)
	}

	if response.ID != "" {
		date := response.CreatedAt.Format("Jan 2 2006")
		fmt.Println("Secret " + response.ID + " created " + date)
	}
}
```

## Revoke Account Secret

Revoke an account secret by its ID. Note that you can't revoke the only secret you have, so add a new one before attempting to delete the existing one.

This method returns a boolean to indicate if the deletion was successful.

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
    auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	accountClient := vonage.NewAccountClient(auth)

	ok, errResp, err := accountClient.DeleteSecret("abcdefab-0000-1111-2222-0123456789ef")

	if err != nil {
		fmt.Println("ERROR: " + errResp.Title + ": " + errResp.Detail)
	}

	if ok {
		fmt.Println("Secret deleted")
	}

}
```
