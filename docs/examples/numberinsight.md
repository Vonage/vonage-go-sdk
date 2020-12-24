---
title: Number Insight API
permalink: examples/numberinsight
---

* [Number Insight Basic/Standard](#number-insight-basic-standard)
* [Number Insight Advanced](#number-insight-advanced)

Use Number Insight to get additional data about a phone number and check its validity. See also:
* Docs: https://developer.nexmo.com/number-insight/overview
* API reference: https://developer.nexmo.com/api/number-insight

## Number Insight Basic/Standard


This is available in Basic/Standard versions:

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	niClient := vonage.NewNumberInsightClient(auth)

	result, _, _ := niClient.Basic("44777000777", vonage.NiOpts{})
    // or for standard:
	// result, _, _ := niClient.Standard("44777000777", vonage.NiOpts{})
    fmt.Println("International Format: " + result.InternationalFormatNumber)
}
```

## Number Insight Advanced

For more detail, try the `AsyncAdvanced` endpoint which will accept your request and then send much more detailed information to the callback endpoint you specified:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	niClient := vonage.NewNumberInsightClient(auth)

	result, _, _ := niClient.AsyncAdvanced("447770007777", "https://example.com/number-insight-data", vonage.NiOpts{})

	if result.Status == 0 {
		http.HandleFunc("/number-insight-data", func(w http.ResponseWriter, r *http.Request) {
			data, _ := ioutil.ReadAll(r.Body)
			fmt.Println(string(data))
		})

		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println("Request status " + string(result.Status) + ": " + result.StatusMessage)
	}
}
```
