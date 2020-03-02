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

### Sending SMS

```golang
package main

import (
	"fmt"

	"github.com/lornajane/goclient-lib/nexmo"
	"github.com/spf13/viper"
)

func main() {
	auth := nexmo.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := nexmo.NewNexmoSMSClient(auth)
	response, err := smsClient.Send("NexmoGolang", "44777000777", "This is a message from golang", nexmo.SMSClientOpts{})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	}
}
```

### Receiving SMS

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

## Getting Help
 
We love to hear from you so if you have questions, comments or find a bug in the project, let us know! You can either:
 
* Open an issue on this repository
* Tweet at us! We're [@NexmoDev on Twitter](https://twitter.com/NexmoDev)
* Or [join the Nexmo Community Slack](https://developer.nexmo.com/community/slack)
 
## Further Reading
 
* Check out the Developer Documentation at <https://developer.nexmo.com> - you'll find the API references for all the APIs there as well
* The documentation for the library: <https://godoc.org/github.com/nexmo-community/nexmo-go>
