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
