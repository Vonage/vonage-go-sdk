# Nexmo Server SDK For Go

[![Go Report Card](https://goreportcard.com/badge/github.com/nexmo-community/nexmo-go)](https://goreportcard.com/report/github.com/nexmo-community/nexmo-go)
[![Build Status](https://travis-ci.org/nexmo-community/nexmo-go.svg?branch=master)](https://travis-ci.org/nexmo-community/nexmo-go)
[![Coverage](https://codecov.io/gh/nexmo-community/nexmo-go/branch/master/graph/badge.svg)](https://codecov.io/gh/nexmo-community/nexmo-go)
[![GoDoc](https://godoc.org/github.com/nexmo-community/nexmo-go?status.svg)](https://godoc.org/github.com/nexmo-community/nexmo-go) 

This is the community-supported Golang library for [Nexmo](https://nexmo.com). It has support for most of our APIs, but is still under active development. Issues, pull requests and other input is very welcome.

If you don't already know Nexmo: We make telephony APIs. If you need to make a call, check a phone number, or send an SMS then you are in the right place! If you don't have a Nexmo yet, you can [sign up for a Nexmo account](https://dashboard.nexmo.com/sign-up?utm_source=DEV_REL&amp;utm_medium=github&amp;utm_campaign=nexmo-go) and get some free credit to get you started.

## Installation

To install the package, use `go get`:

```
go get github.com/nexmo-community/nexmo-go
```

Or import the package into your project and then do `go get .`.

## Usage

Here are some simple examples to get you started. If there's anything else you'd like to see here, please open an issue and let us know! Be aware that this library is still at an alpha stage so things may change between versions.

### Number Insight

```golang
package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(API_KEY, API_SECRET)
	client := nexmo.New(http.DefaultClient, auth)
	insight, _, err := client.Insight.GetBasicInsight(nexmo.BasicInsightRequest{
		Number: PHONE_NUMBER,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Country Name:", insight.CountryName)
	fmt.Println("Local Formatting:", insight.NationalFormatNumber)
	fmt.Println("International Formatting:", insight.InternationalFormatNumber)
}
```

### Sending SMS

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nexmo-community/nexmo-go"
)

func main() {
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(API_KEY, API_SECRET)

	client := nexmo.NewClient(http.DefaultClient, auth)
	smsReq := nexmo.SendSMSRequest {
	    From: FROM_NUMBER,
	    To: TO_NUMBER,
	    Text: "This message comes to you from Nexmo via Golang",
    }

	callR, _, err := client.SMS.SendSMS(smsReq)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", callR.Messages[0].Status)
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

### Starting a Verify Request


```golang
    package main

    import (
        "fmt"
        "github.com/nexmo-community/nexmo-go"
        "log"
        "net/http"
    )

    func verify_start() {
        auth := nexmo.NewAuthSet()
        auth.SetAPISecret(API_KEY, API_SECRET)
        client := nexmo.NewClient(http.DefaultClient, auth)
        verification, _, err := client.Verify.Start(nexmo.StartVerificationRequest{
            Number: PHONE_NUMBER,
            Brand:  "Golang Docs",
        })
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Request ID:", verification.RequestID)
    }

    func main() {
        verify_start()
    }
```

### Confirming a Verify Code

```golang
    package main

    import (
        "fmt"
        "github.com/nexmo-community/nexmo-go"
        "log"
        "net/http"
    )

    func verify_check() {
        auth := nexmo.NewAuthSet()
        auth.SetAPISecret(API_KEY, API_SECRET)
        client := nexmo.NewClient(http.DefaultClient, auth)
        response, _, err := client.Verify.Check(nexmo.CheckVerificationRequest{
            RequestID: REQUEST_ID,
            Code:      CODE,
        })
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Status:", response.Status)
        fmt.Println("Cost:", response.Price)
    }

    func main() {
        verify_check()
    }
```

### Make a Phone Call

The Voice API uses applications and private keys for authentication. To use this snippet you will need a Nexmo number (for the "from" field), and an application with an ID and a private key. You can [learn more about creating applications](https://developer.nexmo.com/application/overview#creating-applications) on the Developer Portal.

```golang
	// get the private key ready, assume file name private.key
	file, file_err := os.Open("private.key")
	if file_err != nil {
		log.Fatal(file_err)
	}
	defer file.Close()

	key, _ := ioutil.ReadAll(file)

	auth := nexmo.NewAuthSet()
	auth.SetApplicationAuth(APPLICATION_ID, key)
	client := nexmo.NewClient(http.DefaultClient, auth)

	to := make([]interface{}, 1)
	to[0] = nexmo.PhoneCallEndpoint{
		Type:   "phone",
		Number: TO_NUMBER,
	}

	response, _, err := client.Call.CreateCall(nexmo.CreateCallRequest{
		From:      nexmo.PhoneCallEndpoint{"phone", NEXMO_NUMBER, ""},
		To:        to,
		AnswerURL: []string{ANSWER_URL},
		EventURL:  []string{EVENT_URL},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", response.Status)
```

## Getting Help
 
We love to hear from you so if you have questions, comments or find a bug in the project, let us know! You can either:
 
* Open an issue on this repository
* Tweet at us! We're [@NexmoDev on Twitter](https://twitter.com/NexmoDev)
* Or [join the Nexmo Community Slack](https://developer.nexmo.com/community/slack)
 
## Further Reading
 
* Check out the Developer Documentation at <https://developer.nexmo.com> - you'll find the API references for all the APIs there as well
* The documentation for the library: <https://godoc.org/github.com/nexmo-community/nexmo-go>
