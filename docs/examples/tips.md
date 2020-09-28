---
title: Tips, Tricks and Troubleshooting
permalink: examples/tips
---

* [Changing the Base URL](#changing-the-base-url)
* [Handling Date Fields](#handling-date-fields)

## Changing the Base URL

If you want to point your API calls to an alternative endpoint (for geographical or local testing reasons this can be useful) try this:

```golang
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	fmt.Println("Hello")

	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := vonage.NewSMSClient(auth)
    smsClient.Config.BasePath = "http://localhost:4010"

	response, err := smsClient.Send("VonageGolang", "44777000777", "This is a message from golang", vonage.SMSOpts{})

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

## Handling Date Fields

Many of our APIs use dates but they come from the API as strings that Go understands as RFC3339 format. Convert to a Go time object with something like:

```go
	start_time, _ := time.Parse(time.RFC3339, response.StartTime)
```

You can then go ahead and use the time object as you usually would.
