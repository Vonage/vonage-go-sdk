---
title: Voice API
permalink: examples/voice
---

* [Make a Phone Call](#make-a-phone-call)
* [Answer a Call or Return an NCCO Response](#answer-a-call-or-return-an-ncco-response)
* [List all Calls](#list-all-calls)
* [Call Detail](#call-detail)
* [End a Call](#end-a-call)
* [Transfer a Call](#transfer-a-call)
* [Mute or Earmuff a Call](#mute-or-earmuff-a-call)
* [Stream Audio into a Call](#stream-audio-into-a-call)
* [Play Text\-To\-Speech into a Call](#play-text-to-speech-into-a-call)
* [Play DTMF Tones into a Call](#play-dtmf-tones-into-a-call)
* [Error Handling](#error-handling)

Check out more resources on Voice API including guides and code snippets on the [developer portal](https://developer.nexmo.com/voice/voice-api/overview). The [API reference](https://developer.nexmo.com/api/voice) will be useful, and there's a [section on NCCOs](/examples/ncco) in these docs too.

## Make a Phone Call

Start a call (the from number should be a Vonage number you own), supplying either `AnswerUrl` *or* `Ncco`:

```go
package main

import (
	"fmt"
	"github.com/vonage/vonage-go-sdk/ncco"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

	from := vonage.CallFrom{Type: "phone", Number: "447770007777"}
	to := vonage.CallTo{Type: "phone", Number: "447770007788"}

	MyNcco := ncco.Ncco{}
	talk := ncco.TalkAction{Text: "Go library calling to say hello", VoiceName: "Nicole"}
	MyNcco.AddAction(talk)

    // NCCO example
	result, _, _ := client.CreateCall(vonage.CreateCallOpts{From: from, To: to, Ncco: MyNcco})
    // alternate version with answer URL
    //result, _, _ := client.CreateCall(CreateCallOpts{From: from, To: to, AnswerUrl: []string{"https://example.com/answer"}})
	fmt.Println(result.Uuid + " call ID started")
}

```

See [NCCO](#nccos) for more information and examples for all other supported NCCO types.

## Answer a Call or Return an NCCO Response

Often, you will want to return an NCCO as an HTTP response rather than pass the object into an API call. Here's an example of serving an NCCO as a response:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vonage/vonage-go-sdk/ncco"
)

func answer(w http.ResponseWriter, req *http.Request) {

	MyNcco := ncco.Ncco{}

	talk := ncco.TalkAction{Text: "Thank you for calling."}
	MyNcco.AddAction(talk)

	data, _ := json.Marshal(MyNcco)
	fmt.Fprintf(w, "%s", data)
}

func main() {

	http.HandleFunc("/answer", answer)

	http.ListenAndServe(":8081", nil)
}
```

This is useful for answering a call (as shown above) but also for handling webhooks that expect an NCCO response such as notify, input, and so on.

## List all Calls

A list of all the calls associated with your account.

```go
package main

import (
	"fmt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

	response, _, _ := client.GetCalls()
	fmt.Println(response.Embedded.Calls[0].Uuid + " status: " + response.Embedded.Calls[0].Status)
}
```

## Call Detail

If you have the UUID of the call, fetch the details of it:

```go
package main

import (
	"fmt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

	response, _, _ := client.GetCall("aaaabbbb-0000-1111-2222-abcdef01234567")
    t1, _ := time.Parse(time.RFC3339, response.StartTime)
	date_string := t1.Format("Jan _2 2006 at 15:04:05")
	fmt.Println("Call started: " + date_string + ", duration " + result1.Duration + " seconds and status: " + result1.Status)
}
```

The example includes how to parse and then format a date.


## End a Call

End a call using the `hangup()` method on the client:

```go
package main

import (
	"fmt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)
	result, _, _ := client.Hangup("aaaabbbb-0000-1111-2222-abcdef01234567")
	fmt.Println("Status: " + result.Status) // Status: 0 is good
}

```


## Transfer a Call

This requires the Uuid of an existing call. The example below follows the "Make a Phone Call" example and assumes a `result` variable from that example.

```go
package main

import (
	"fmt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

	MyNcco := ncco.Ncco{}
	talk := ncco.TalkAction{Text: "Go library calling to interrupt", VoiceName: "Nicole"}
	MyNcco.AddAction(talk)


    // NCCO example
	result, _, _ := client.TransferCall(vonage.TransferCallOpts{Uuid: result.Uuid, Ncco: MyNcco})
    // handy AnswerUrl example
	// result, _, _ := client.TransferCall(TransferCallOpts{Uuid: result.Uuid, AnswerUrl: []string{"https://raw.githubusercontent.com/nexmo-community/ncco-examples/gh-pages/talk.json"}})
	fmt.Println("Status: " + result.Status)
}

```

See [NCCO](#nccos) for more information and examples for all other supported NCCO types.

## Mute or Earmuff a Call

These actions are similar to one another. To "earmuff" a call makes the call inaudible to the user. To "mute" the call makes the user inaudible to the call. The library offers the following methods:
 * `Mute()`
 * `Unmute()`
 * `Earmuff()`
 * `Unearmuff()`

They all accept the UUID of the in-progress call, so the code looks like this:

```go
package main

import (
	"fmt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

	result, _, _ := client.Mute("aaaabbbb-0000-1111-2222-abcdef01234567")
	fmt.Println("Status: " + result.Status) // Status: 0 is good
}

```

Replace `Mute()` with your desired method name.

## Stream Audio into a Call

You can stream (and stop streaming) audio from a public URL into an in-progress call, like this:

```go
package main

import (
	"fmt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

    result, _, _:= client.PlayAudioStream("aaaabbbb-0000-1111-2222-abcdef01234567",
        "https://raw.githubusercontent.com/nexmo-community/ncco-examples/gh-pages/assets/welcome_to_nexmo.mp3",
        vonage.PlayAudioOpts{}
    )
    // or to stop the audio
    // result, _, _:= client.StopAudioStream("aaaabbbb-0000-1111-2222-abcdef01234567")
	fmt.Println("Update message: " + result.Message)
}

```

## Play Text-To-Speech into a Call

You can send (and stop sending) TTS (Text To Speech) into an in-progress call. Here's an example:

```go
package main

import (
	"fmt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

    result, _, _:= client.PlayTts("aaaabbbb-0000-1111-2222-abcdef01234567",
        "Hello, my friend",
        vonage.PlayTtsOpts{Loop: 2, Style: 0, Language: "en-US"}
    )
    // or to stop an in-progress TTS
    // result, _, _:= client.StopTts("aaaabbbb-0000-1111-2222-abcdef01234567")
	fmt.Println("Update message: " + result.Message)
}

```

## Play DTMF Tones into a Call

You can send DTMF (keypad tones) into an in-progress call. Here's an example:

```go
package main

import (
	"fmt"
)

func main() {
    privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

    result, _, _:= client.PlayDtmf("aaaabbbb-0000-1111-2222-abcdef01234567", "123")
	fmt.Println("Update message: " + result.Message)
}

```

## Error Handling

There are three return values on most methods. The first two are structs representing the fields in the success and error response for the API endpoint involved. The final value is an error, but in many cases this can be type asserted to a more useful `GenericOpenAPIError`, like this:

```
	response, _, http_error := client.GetCalls()

	if http_error != nil {
        e := http_error.(voice.GenericOpenAPIError)
        // output the status code
        fmt.Println(e.Error())
        // print the whole API response
        fmt.Println(string(e.Body()))
	}

```
