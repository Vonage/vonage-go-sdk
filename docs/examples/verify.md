---
title: Verify API
permalink: examples/verify
---

* [Verify a User's Phone Number](#verify-a-users-phone-number)
* [Verify a Payment via Phone (PSD2)](#verify-a-payment-via-phone-psd2)
* [Check Verification Code](#check-verification-code)
* [Cancel a Verification](#cancel-a-verification)
* [Trigger the Next Event in a Verification](#trigger-the-next-event-in-a-verification)
* [Search for a Verification](#search-for-a-verification)

Verify API is great for checking contact numbers and for 2fa. Read more on the [developer portal](https://developer.nexmo.com/verify/overview) and [API reference](https://developer.nexmo.com/api/verify).

## Verify a User's Phone Number

This is a multi-step process. First: request that the number be verified and state what "brand" is asking.

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := vonage.NewVerifyClient(auth)

    response, errResp, err := verifyClient.Request("44777000777", "GoTest", vonage.VerifyOpts{CodeLength: 6, Lg: "es-es", WorkflowID: 4})

    if err != nil {
        fmt.Printf("%#v\n", err)
    } else if response.Status != "0" {
        fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
    } else {
        fmt.Println("Request started: " + response.RequestId)
    }
}
```

Copy the request ID; the user will receive a PIN code and when they have it, you can check the code (see [Check PIN Code section](#check-verification-code))

## Verify a Payment via Phone (PSD2)

Just like verifying a user's number, this is a multi-step process. First: send a code to the user alongside information about the amount and destination of the payment.

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := vonage.NewVerifyClient(auth)

    response, errResp, err := verifyClient.Psd2("44777000777", "GoTestRetail", 45.67, vonage.VerifyPsd2Opts{})

    if err != nil {
        fmt.Printf("%#v\n", err)
    } else if response.Status != "0" {
        fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
    } else {
        fmt.Println("Request started: " + response.RequestId)
    }
}
```

Copy the request ID; the user will receive a PIN code and when they have it, you can check the code (see next section)

## Check Verification Code

When a request is in progress, use the `requestId` and the PIN code sent by the user to check if it is correct:

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := vonage.NewVerifyClient(auth)

	response, errResp, err := verifyClient.Check(REQUEST_ID, PIN_CODE)

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		// all good
		fmt.Println("Request complete: " + response.RequestId)
	}
}
```

If status is zero, the code was correct and you have confirmed the user owns the number

## Cancel a Verification

If you have a verification in progress, and no longer wish to proceed, you can cancel it. This can be done from 30 seconds after the verification was requested, until the second event occurs.

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := vonage.NewVerifyClient(auth)
   	response, errResp, err := verifyClient.Cancel(REQUEST_ID)

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		// all good
		fmt.Println("Request cancelled: " + response.RequestId)
	}
}
```

## Trigger the Next Event in a Verification

If for example, an SMS has been sent, and you'd immediately like to have the user get a TTS call (depending on the [workflow](https://developer.nexmo.com/verify/guides/workflows-and-events) in use), it's possible to make the next event happen on demand:

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := vonage.NewVerifyClient(auth)
   	response, errResp, err := verifyClient.TriggerNextEvent(REQUEST_ID)

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		// all good
		fmt.Println("Next event triggered for request: " + response.RequestId)
	}
}
```

## Search for a Verification

You can check on an in-progress or completely (either successfully or not) verification by its request ID:

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := vonage.NewVerifyClient(auth)
   	response, errResp, err := verifyClient.Search(REQUEST_ID)

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		// all good
		fmt.Println("Next event triggered for request: " + response.RequestId)
	}
}
```


