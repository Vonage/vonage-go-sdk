---
title: Application API
permalink: examples/application
---

* [List Applications:](#list-applications)
* [Get One Application By ID](#get-one-application-by-id)
* [Create a New Application](#create-a-new-application)
* [Create a New Application with a Public Key](#create-a-new-application-with-a-public-key)
* [Creating an Application for All Capabilities](#creating-an-application-for-all-capabilities)
* [Update an Application](#update-an-application)
* [Delete an Application](#delete-an-application)

Working with the Voice, Messages or Conversation APIs relies on using applications. Check out the [docs](https://developer.nexmo.com/application/overview) and [API reference](https://developer.nexmo.com/api/application.v2) for more details.

## List Applications

To see all the applications associated with your account:

```go
package main

import (
    "fmt"
    "strconv"

    "github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	appClient := vonage.NewApplicationClient(auth)

    result, _, _:= appClient.GetApplications(vonage.GetApplicationsOpts{})
    fmt.Println("Application count: " + strconv.FormatInt(int64(result.TotalItems), 10))
    for _, app := range result.Embedded.Applications {
        fmt.Println(app.Name)
    }
}
```

## Get One Application By ID

For details on one application, you can request it by its ID with code like this:

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	appClient := vonage.NewApplicationClient(auth)

	result, _, _ := appClient.GetApplication(app_id)
    fmt.Println(result.Name)
```


## Create a New Application

To create an application, supply the configuration for the types of operations you'll need. The example below shows how to create an application for Voice API:

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	appClient := vonage.NewApplicationClient(auth)

    voice := vonage.ApplicationVoice{
        Webhooks: vonage.ApplicationVoiceWebhooks{
            AnswerUrl:         vonage.ApplicationUrl{Address: "https://example.com/answer", HttpMethod: "POST"},
            EventUrl:          vonage.ApplicationUrl{Address: "https://example.com/event", HttpMethod: "POST"},
        },
    }

    opts := vonage.CreateApplicationOpts{Capabilities: vonage.ApplicationCapabilities{Voice: &voice}}
    result, _, _:= appClient.CreateApplication("MyGoVoiceApp", opts)
    fmt.Println("App ID: " + result.Id)
    if result.Keys.PrivateKey != "" {
        // if the user supplied a public key, the private key isn't returned because they have it already
        fmt.Println("Private Key (save this for later):")
        fmt.Println(result.Keys.PrivateKey)
    }
    
}
```

## Create a New Application with a Public Key

By default, Vonage will generate the public/private key pair for your application. If you'd prefer to generate your own and supply the public key, you can do that with code like this example:

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	appClient := vonage.NewApplicationClient(auth)

    voice := vonage.ApplicationVoice{
        Webhooks: vonage.ApplicationVoiceWebhooks{
            AnswerUrl:         vonage.ApplicationUrl{Address: "https://example.com/answer", HttpMethod: "POST"},
            EventUrl:          vonage.ApplicationUrl{Address: "https://example.com/event", HttpMethod: "POST"},
        },
    }

    publicKey, _ := ioutil.ReadFile("keyfile.pub")

    opts := vonage.CreateApplicationOpts{Capabilities: vonage.ApplicationCapabilities{Voice: &voice}, Keys: vonage.ApplicationKeys{PublicKey: string(publicKey)}}

    result, _, _:= appClient.CreateApplication("MyGoVoiceApp", opts)
    fmt.Println("App ID: " + result.Id)
}
```

## Creating an Application for All Capabilities

The various capabilities each have their own data structure. Here's an example that uses all of them; add as many as you need for your application:

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	appClient := vonage.NewApplicationClient(auth)

    voice := vonage.ApplicationVoice{
        Webhooks: vonage.ApplicationVoiceWebhooks{
            AnswerUrl:         vonage.ApplicationUrl{Address: "https://example.com/answer", HttpMethod: "POST"},
            EventUrl:          vonage.ApplicationUrl{Address: "https://example.com/event", HttpMethod: "POST"},
        },
    }

    rtc := vonage.ApplicationRtc{
        Webhooks: vonage.ApplicationRtcWebhooks{
            EventUrl: vonage.ApplicationUrl{Address: "https://example.com/rtc-event", HttpMethod: "POST"},
        },
    }

	mesg := vonage.ApplicationMessages{
		Webhooks: vonage.ApplicationMessagesWebhooks{
			StatusUrl:  vonage.ApplicationUrl{Address: "https://example.com/status", HttpMethod: "POST"},
			InboundUrl: vonage.ApplicationUrl{Address: "https://example.com/inbound", HttpMethod: "POST"},
		},
	}

    opts := vonage.CreateApplicationOpts{Capabilities: vonage.ApplicationCapabilities{Voice: &voice, Messages: &mesg, Vbc: &vonage.ApplicationVbc{}, Rtc: &rtc}}

    result, _, _:= appClient.CreateApplication("MyGoOmniApp", opts)
    fmt.Println("App ID: " + result.Id)
}
```

## Update an Application

Much like the creation feature, this method supports a new private key and any capabilities you wish to add. You must also supply the application's name (or new name if you want to change it) and ID:

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	appClient := vonage.NewApplicationClient(auth)

    voice := vonage.ApplicationVoice{
        Webhooks: vonage.ApplicationVoiceWebhooks{
            AnswerUrl:         vonage.ApplicationUrl{Address: "https://example.com/answer", HttpMethod: "POST"},
            EventUrl:          vonage.ApplicationUrl{Address: "https://example.com/event", HttpMethod: "POST"},
        },
    }

    opts := vonage.CreateApplicationOpts{Capabilities: vonage.ApplicationCapabilities{Voice: &voice}}
    result, _, _:= appClient.UpdateApplication("aaaabbbb-cccc-dddd-eeee-0123456789ff", "MyUpdatedGoApp", opts)
    fmt.Println("App ID: " + result.Id)
}
```

## Delete an Application

To delete the application, supply the application ID:

```go
package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	appClient := vonage.NewApplicationClient(auth)
    result, appErr, _:= appClient.DeleteApplication("aaaabbbb-cccc-dddd-eeee-0123456789ff")

	if !result {
		fmt.Printf("%#v\n", appErr)
	} else {
		fmt.Println("App deleted")
	}
}
```


