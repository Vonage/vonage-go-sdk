# Nexmo Server SDK For Go

[![Go Report Card](https://goreportcard.com/badge/github.com/nexmo-community/nexmo-go)](https://goreportcard.com/report/github.com/nexmo-community/nexmo-go)
[![Build Status](https://travis-ci.org/nexmo-community/nexmo-go.svg?branch=master)](https://travis-ci.org/nexmo-community/nexmo-go)
[![Coverage](https://codecov.io/gh/nexmo-community/nexmo-go/branch/master/graph/badge.svg)](https://codecov.io/gh/nexmo-community/nexmo-go)
[![GoDoc](https://godoc.org/github.com/nexmo-community/nexmo-go?status.svg)](https://godoc.org/github.com/nexmo-community/nexmo-go) 

<img src="https://developer.nexmo.com/assets/images/Vonage_Nexmo.svg" height="48px" alt="Nexmo is now known as Vonage" />

This is the community-supported Golang library for [Nexmo](https://nexmo.com). It has support for most of our APIs, but is still under active development. Issues, pull requests and other input is very welcome.

If you don't already know Nexmo: We make telephony APIs. If you need to make a call, check a phone number, or send an SMS then you are in the right place! If you don't have a Nexmo yet, you can [sign up for a Nexmo account](https://dashboard.nexmo.com/sign-up?utm_source=DEV_REL&amp;utm_medium=github&amp;utm_campaign=nexmo-go) and get some free credit to get you started.

> Please note that this project is released with a Contributor Code of Conduct. By participating in this project you agree to abide by its terms.

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

## SMS API
#### Send SMS

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

#### Send Unicode SMS

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

#### Receive SMS

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


### Verify API

#### Verify a User's Phone Number

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

    response, errResp, err := verifyClient.Request("44777000777", "GoTest", nexmo.VerifyOpts{CodeLength: 6, Lg: "es-es", WorkflowId: 4})

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

#### Check Verification Code

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

#### Cancel a Verification

If you have a verification in progress, and no longer wish to proceed, you can cancel it. This can be done from 30 seconds after the verification was requested, until the second event occurs.

```golang
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := nexmo.NewVerifyClient(auth)
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

#### Trigger the Next Event in a Verification

If for example, an SMS has been sent, and you'd immediately like to have the user get a TTS call (depending on the [workflow](https://developer.nexmo.com/verify/guides/workflows-and-events) in use), it's possible to make the next event happen on demand:

```golang
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := nexmo.NewVerifyClient(auth)
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

#### Search for a Verification

You can check on an in-progress or completely (either successfully or not) verification by its request ID:

```golang
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	verifyClient := nexmo.NewVerifyClient(auth)
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


### Number Management

#### List the Numbers You Own

To check on the numbers already associated with your account:

```
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := nexmo.NewNumbersClient(auth)
	response, err := numbersClient.List(nexmo.NumbersOpts{})

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

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := nexmo.NewNumbersClient(auth)
	response, err := numbersClient.List(nexmo.NumbersOpts{HasApplication: "false"})

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

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := nexmo.NewNumbersClient(auth)
    response, respErr, err := numbersClient.Search("GB", nexmo.NumberSearchOpts{Size: 10})
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

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := nexmo.NewNumbersClient(auth)

    response, resp, err := numbersClient.Buy("GB", "44777000777", nexmo.NumberBuyOpts{})
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

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := nexmo.NewNumbersClient(auth)
	response, resp, err := numbersClient.Update("GB", "44777000777", nexmo.NumberUpdateOpts{AppID: " aaaaaaaa-bbbb-cccc-dddd-0123456789abc"})

	fmt.Printf("%#v\n", response)
}
```

#### Cancel a bought number

If you don't need a number any more, you can cancel it and stop paying the rental for it:

```
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	numbersClient := nexmo.NewNumbersClient(auth)

    response, resp, err := numbersClient.Cancel("GB", "44777000777", nexmo.NumberCancelOpts{})
    if err != nil {
        panic(err)
    }
    fmt.Printf("%#v\n", response)
}
```


### JWT Authentication
#### Generate a Basic JWT

Generate a JSON Web Token (JWT) for the APIs that use that. You usually won't need to do this if you're using the library but if you need to make a custom request or want to use a JWT for something else, you can use this.

```go
package main

import (
	"fmt"

	"github.com/nexmo-community/nexmo-go/jwt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
    g := jwt.NewGenerator(APPLICATION_ID, privateKey)

    token, _ := g.GenerateToken()
    fmt.Println(token)
}
```

#### Generate a JWT with more options

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


### NCCOs

NCCO (Nexmo Call Control Object) is the format for describing the various actions that will take place during a call. Check the [NCCO reference on the developer portal](https://developer.nexmo.com/voice/voice-api/ncco-reference) for full details, but examples of each action are included in the sections below.

The basic approach is to create an NCCO, then create actions to add into it:

```go
	ncco := nexmo.Ncco{}
	talk := nexmo.TalkAction{Text: "Greetings from the golang library", VoiceName: "Nicole"}
	ncco.AddAction(talk)
```

#### Talk Action

Create a `talk` action to read some text into the call:

```go
	talk := nexmo.TalkAction{Text: "Greetings from the golang library", VoiceName: "Nicole"}
```

#### Notify Action

Use `notify` to send a particular data payload to a nominated URL:

```go
	url := make([]string, 1)
	url[0] = "https://example.com/webhooks/notify"
	data := make(map[string]string)
	data["stage"] = "Registration"
	ping := nexmo.NotifyAction{EventUrl: url, Payload: data}
```

This feature is useful for marking progress through a call and that the user is still connected.

#### Record Action

Send a `record` action to start a recording:

```go
    record := nexmo.RecordAction{BeepStart: true}
```

When the recording completes, Nexmo sends a webhook containing the recording URL so that you can download the file.

#### Conversation Action

Adds the call to a conversation:

```go
    conversation := nexmo.ConversationAction{Name: "convo1"}
```

#### Connect Action

Connects the current call to another endpoint (currently only phone is supported):

```go
    endpoint := make([]nexmo.Endpoint, 1)
	endpoint[0] = nexmo.PhoneEndpoint{Number: "44777000777"}
	connect := nexmo.ConnectAction{Endpoint: endpoint, From: "44777000888"}
```
The `from` field when connecting to a phone endpoint should be a Nexmo number that you own.

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
```
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
