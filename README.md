# Nexmo Client Library For Go

This library is moving towards a full client implementation of the
[Nexmo](https://www.nexmo.com/) APIs. The library is not currently officially
supported by Nexmo, but the author, [@judy2k] works in Nexmo's Developer
Relations team. The hope is that this library will become popular enough to
justify becoming an officially supported Nexmo library.

The library currently has high coverage for the following APIs:

* Voice
* SMS
* Insight
* Verify
* Application

It currently has only a handful of Developer endpoints implemented, and no
webhook support.

## Usage

Usage looks a bit like this:

```golang
package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/judy2k/nexmo"
)

func main() {
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(API_KEY, API_SECRET)
	httpClient := http.Client{}
	client := nexmo.NewClient(&httpClient, auth)
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

More documentation will be coming as the API stabilises! Things are still in flux.

## To Do

Lots has been done, but there's still lots left to do! If you'd like to help,
please get in touch **first**! Progress is moving swiftly and I wouldn't want
to waste your time!

### Testing

There isn't any testing yet! A test harness is currently in the works, to
allow some quality tests to be written. Every API call has been tested
manually during development, so they *do* work, but there's more work to
be done.

### Error Handling

Error responses from Nexmo APIs are not currently dealt with very gracefully.
Fortunately, much of the work can be done in one place (inside the custom fork
of [Sling](https://github.com/dghubble/sling) - which is why I forked it).
This comes directly after testing, so I can ensure that all the different
error responses in different parts of the API are dealt with properly.

### Remaining Endpoints

Coverage is actually really high. 90% of Voice, SMS, Insight, Verify &
Application APIs are covered, with some of the Developer API as well.

### Webhook Support

Support for webhook parsing and validation. Made easier because Golang has
a widely-used request/response API!