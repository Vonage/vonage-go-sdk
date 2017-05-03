# Nexmo Client Library For Go

[![Go Report Card](https://goreportcard.com/badge/github.com/judy2k/nexmo-go)](https://goreportcard.com/report/github.com/judy2k/nexmo-go)
[![Build Status](https://travis-ci.org/judy2k/nexmo-go.svg?branch=master)](https://travis-ci.org/judy2k/nexmo-go)
[![Coverage](https://codecov.io/gh/judy2k/nexmo-go/branch/master/graph/badge.svg)](https://codecov.io/gh/judy2k/nexmo-go)


This library is moving towards a full client implementation of the
[Nexmo](https://www.nexmo.com/) APIs. The library is not currently officially
supported by Nexmo, but the author, [@judy2k](https://twitter.com/judy2k)
works in Nexmo's Developer Relations team. The hope is that this library will 
become popular enough to justify becoming an officially supported
Nexmo library.

The library currently has good coverage for the following APIs:

API         | Coverage
------------|---------:
Voice       | (9/9)
SMS         | (1/4) 
Insight     | (3/4)
Verify      | (4/4)
Application | (5/5)

It currently has only a handful of Developer (5/15) endpoints implemented, and no
webhook support.

Current API Coverage can be found in [this spreadsheet](https://docs.google.com/spreadsheets/d/19lsAoW2oiGMK7Xg0dOw5KPdOOix1Oo-GaTWkRyVRMXI/pubhtml#)

## Usage

Usage looks a bit like this:

```golang
package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/judy2k/nexmo-go"
)

func main() {
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(API_KEY, API_SECRET)
	client := nexmo.NewClient(http.DefaultClient, auth)
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
