---
title: SMS API
permalink: examples/sms
---

* [Send SMS](#send-sms)
* [Send Unicode SMS](#send-unicode-sms)
* [Receive SMS](#receive-sms)

SMS API is one of our most-used APIs. Check out the [documentation](https://developer.nexmo.com/messaging/sms/overview) and [API reference](https://developer.nexmo.com/api/sms) for more details.

## Send SMS

To send an SMS, try the code below:

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
    auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
    smsClient := vonage.NewSMSClient(auth)
    response, errResp, err := smsClient.Send("44777000000", "44777000777", "This is a message from golang", vonage.SMSOpts{})

    if err != nil {
        panic(err)
    }

    if response.Messages[0].Status == "0" {
        fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
    } else {
        fmt.Println("Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
    }
}
```

## Send Unicode SMS

Add `Type` to the `opts` parameter and set it to "unicode":

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
    auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
    smsClient := vonage.NewSMSClient(auth)
    response, errResp, err := smsClient.Send("44777000000", "44777000777", "こんにちは世界", vonage.SMSOpts{Type: "unicode"})

    if err != nil {
        panic(err)
    }

    if response.Messages[0].Status == "0" {
        fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
    } else {
        fmt.Println("Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
    }
}
```

## Receive SMS

To receive an SMS, you will need to run a local webserver and expose the URL publicly (you can use a tool such as [ngrok](https://ngrok.com).

```golang
package main

import (
	"fmt"
	"net/http"

	"github.com/vonage/vonage-go-sdk"
)

func main() {

	http.HandleFunc("/webhooks/inbound-sms", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		fmt.Println("SMS from " + params["msisdn"][0] + ": " + string(params["text"][0]))
	})

	http.ListenAndServe(":8080", nil)
}
```


