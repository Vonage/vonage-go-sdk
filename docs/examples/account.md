---
title: Account API
permalink: examples/account
---

Account API gives access to check your balance, rotate your account secrets and configure the default behaviour of your account. Check out the [documentation](https://developer.nexmo.com/account/overview) and [API reference](https://developer.nexmo.com/api/account) for more details.

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

