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

	result, _, _ := niClient.Basic("44777000777")
    // or for standard:
	// result, _, _ := niClient.Standard("44777000777")
    fmt.Println("International Format: " + result.InternationalFormatNumber)
}
```

## Number Insight Advanced

For more detail, try the Async endpoint which will accept your request and then send much more detailed information to the callback endpoint you specified:

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	niClient := vonage.NewNumberInsightClient(auth)

	result, _, _ := niClient.Async("44777000777", "https://example.com/number-insight-data")
    fmt.Println("Status: " + result.Status) // 0 is good! Data will arrive to the callback
}
```
