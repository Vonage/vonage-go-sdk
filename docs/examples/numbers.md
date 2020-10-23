---
title: Numbers API
permalink: examples/numbers
---

* [List the Numbers You Own](#list-the-numbers-you-own)
* [Search for a number to buy](#search-for-a-number-to-buy)
* [Buy a number](#buy-a-number)
* [Update number configuration](#update-number-configuration)
* [Cancel a bought number](#cancel-a-bought-number)

Learn more about the Numbers API by visiting the [docs on the developer portal](https://developer.nexmo.com/numbers/overview) and the [API reference](https://developer.nexmo.com/api/numbers).

#### List the Numbers You Own

To check on the numbers already associated with your account:

```
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := vonage.NewNumbersClient(auth)
	response, err := numbersClient.List(vonage.NumbersOpts{})

	if err != nil {
		panic(err)
	}

	for _, number := range response.Numbers {
		fmt.Println("Number: " + number.Msisdn + " (" + number.Country + ") with: " + strings.Join(number.Features, ", "))
	}
}
```

You can also filter by which applications a number is linked to (set the `ApplicationId` param) or as in the example below, by whether it is linked to an application at all:

```
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := vonage.NewNumbersClient(auth)
	response, err := numbersClient.List(vonage.NumbersOpts{HasApplication: "false"})

	if err != nil {
		panic(err)
	}

	for _, number := range response.Numbers {
		fmt.Println("Number: " + number.Msisdn + " (" + number.Country + ") with: " + strings.Join(number.Features, ", "))
	}
}
```

#### Search for a number to buy

You can programmatically add numbers to your account:

```
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := vonage.NewNumbersClient(auth)
    response, respErr, err := numbersClient.Search("GB", vonage.NumberSearchOpts{Size: 10})
    if err != nil {
        panic(err)
    }
    if respErr.ErrorCode != "" {
        fmt.Println("Error " + respErr.ErrorCode + ": " + respErr.ErrorCodeLabel)
    }
    for _, number := range response.Numbers {
        fmt.Println("Number: " + number.Msisdn + " (" + number.Country + ") with: " + strings.Join(number.Features, ", "))
    }
}
```

#### Buy a number

Use the search to find a number that suits your needs (see the previous example), then buy it:

```
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := vonage.NewNumbersClient(auth)

    response, resp, err := numbersClient.Buy("GB", "44777000777", vonage.NumberBuyOpts{})
    if err != nil {
        panic(err)
    }
    fmt.Printf("%#v\n", response)
}
```

Failures in this action can indicate that more information is needed when buying the number. If you get a 420 status code, try buying via the dashboard <https://developer.nexmo.com>.

#### Update number configuration

This endpoint is how you configure the number behaviour. There are a few properties you can set, they are named to match the [Number API Reference](https://developer.nexmo.com/api/number).

* `MoHTTPURL` - The URL for incoming ("mobile originated", hence the name) SMS API webhooks.
* `AppID` - The application ID to use for configuration (this is the most common setup for most apps)
* `VoiceCallbackType` - Can be `tel` or `sip`
* `VoiceCallbackValue` - The value 
f telephone number or sip connection as appropriate

```
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := vonage.NewNumbersClient(auth)
	response, resp, err := numbersClient.Update("GB", "44777000777", vonage.NumberUpdateOpts{AppID: " aaaaaaaa-bbbb-cccc-dddd-0123456789abc"})

	fmt.Printf("%#v\n", response)
}
```

#### Cancel a bought number

If you don't need a number any more, you can cancel it and stop paying the rental for it:

```
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := vonage.NewNumbersClient(auth)

    response, resp, err := numbersClient.Cancel("GB", "44777000777", vonage.NumberCancelOpts{})
    if err != nil {
        panic(err)
    }
    fmt.Printf("%#v\n", response)
}
```


