# Nexmo Server SDK For Go

[![Go Report Card](https://goreportcard.com/badge/github.com/nexmo-community/nexmo-go)](https://goreportcard.com/report/github.com/nexmo-community/nexmo-go)
[![Build Status](https://travis-ci.org/nexmo-community/nexmo-go.svg?branch=master)](https://travis-ci.org/nexmo-community/nexmo-go)
[![Coverage](https://codecov.io/gh/nexmo-community/nexmo-go/branch/master/graph/badge.svg)](https://codecov.io/gh/nexmo-community/nexmo-go)
[![GoDoc](https://godoc.org/github.com/nexmo-community/nexmo-go?status.svg)](https://godoc.org/github.com/nexmo-community/nexmo-go) 

<img src="https://developer.nexmo.com/assets/images/Vonage_Nexmo.svg" height="48px" alt="Nexmo is now known as Vonage" />

This is the community-supported Golang library for [Nexmo](https://nexmo.com). It has support for most of our APIs, but is still under active development. Issues, pull requests and other input is very welcome.

If you don't already know Nexmo: We make telephony APIs. If you need to make a call, check a phone number, or send an SMS then you are in the right place! If you don't have a Nexmo yet, you can [sign up for a Nexmo account](https://dashboard.nexmo.com/sign-up?utm_source=DEV_REL&amp;utm_medium=github&amp;utm_campaign=nexmo-go) and get some free credit to get you started.

## Installation

Find current and past releases on the [releases page](https://github.com/nexmo-community/nexmo-go/releases).

## Recommended process (Go 1.13+)

Import the package and use it:

```
import ("github.com/nexmo-community/nexmo-go")
```

## Older versions of Go (<= 1.12)

To install the package, use `go get`:

```
go get github.com/nexmo-community/nexmo-go
```

Or import the package into your project and then do `go get .`.

## Usage

Here are some simple examples to get you started. If there's anything else you'd like to see here, please open an issue and let us know! Be aware that this library is still at an alpha stage so things may change between versions.

### Send SMS

To send an SMS, try the code below:

```golang
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go/nexmo"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := nexmo.NewSMSClient(auth)
	response, err := smsClient.Send("NexmoGolang", "44777000777", "This is a message from golang", nexmo.SMSOpts{})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	}
}
```

### Send Unicode SMS

Add `Type` to the `opts` parameter and set it to "unicode":

```golang
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go/nexmo"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := nexmo.NewSMSClient(auth)
	response, err := smsClient.Send("NexmoGolang", "44777000777", "こんにちは世界", nexmo.SMSOpts{Type: "unicode"})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	}
}
```

### Receive SMS

To receive an SMS, you will need to run a local webserver and expose the URL publicly (you can use a tool such as [ngrok](https://ngrok.com).

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/webhooks/inbound-sms", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		fmt.Println("SMS from " + params["msisdn"][0] + ": " + string(params["text"][0]))
	})

	http.ListenAndServe(":8080", nil)
}
```

### Verify a User's Phone Number

This is a multi-step process. First: request that the number be verified and state what "brand" is asking.

```golang
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := nexmo.NewVerifyClient(auth)

    response, errResp, err := verifyClient.Request("447846810475", "GoTest", nexmo.VerifyOpts{CodeLength: 6, Lg: "es-es", WorkflowId: 4})

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

### Check Verification Code

When a request is in progress, use the `requestId` and the PIN code sent by the user to check if it is correct:

```golang
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := nexmo.NewVerifyClient(auth)

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

### Cancel a Verification

### Trigger the Next Event

### List the Numbers You Own

To check on the numbers already associated with your account:

```
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := nexmo.NewVerifyClient(auth)
	response, err := numbersClient.List()

	if err != nil {
		panic(err)
	}

	for _, number := range response.Numbers {
		fmt.Println("Number: " + number.Msisdn + " (" + number.Country + ") with: " + strings.Join(number.Features, ", "))
	}
}
```

## Tips, Tricks and Troubleshooting

### Changing the Base URL

If you want to point your API calls to an alternative endpoint (for geographical or local testing reasons this can be useful) try this:

```golang
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go/nexmo"
)

func main() {
	fmt.Println("Hello")

	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := nexmo.NewSMSClient(auth)
    smsClient.Config.BasePath = "http://localhost:4010"

	response, err := smsClient.Send("NexmoGolang", "44777000777", "This is a message from golang", nexmo.SMSOpts{})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	}
}

_(The example above shows using the library with [Prism](https://github.com/stoplightio/prism), which we find useful at development time)_

The fields for configuration are:
- `BasePath` (shown in the example above) overrides where the requests should be sent to
- `DefaultHeader` is a map, add any custom headers you need here
- `HTTPClient` is a pointer to an httpClient if you need to change any networking settings

## Getting Help
 
We love to hear from you so if you have questions, comments or find a bug in the project, let us know! You can either:
 
* Open an issue on this repository
* Tweet at us! We're [@VonageDev on Twitter](https://twitter.com/VonageDev)
* Or [join the Vonage Community Slack](https://developer.nexmo.com/community/slack)
 
## Further Reading
 
* Check out the Developer Documentation at <https://developer.nexmo.com> - you'll find the API references for all the APIs there as well
* The documentation for the library: <https://godoc.org/github.com/nexmo-community/nexmo-go>
