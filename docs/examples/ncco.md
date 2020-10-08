---
title: NCCOs
permalink: examples/ncco
---

* [Talk Action](#talk-action)
* [Notify Action](#notify-action)
* [Record Action](#record-action)
* [Conversation Action](#conversation-action)
* [Stream Action](#stream-action)
* [Connect Action](#connect-action)

NCCO (Nexmo Call Control Object) is the format for describing the various actions that will take place during a call. Check the [NCCO reference on the developer portal](https://developer.nexmo.com/voice/voice-api/ncco-reference) for full details, but examples of each action are included in the sections below.

The basic approach is to create an NCCO, then create actions to add into it:

```go
	MyNcco := ncco.Ncco{}
	talk := ncco.TalkAction{Text: "Greetings from the golang library", VoiceName: "Nicole"}
	MyNcco.AddAction(talk)
```

## Talk Action

Create a `talk` action to read some text into the call:

```go
	talk := ncco.TalkAction{Text: "Greetings from the golang library", VoiceName: "Nicole"}
```

or

```go
	talk := ncco.TalkAction{Text: "Greetings from the golang library", Style: 0, Language: "en-US}
```

## Notify Action

Use `notify` to send a particular data payload to a nominated URL:

```go
	url := []string{"https://example.com/webhooks/notify"}
	data := make(map[string]string)
	data["stage"] = "Registration"
	ping := ncco.NotifyAction{EventUrl: url, Payload: data}
```

This feature is useful for marking progress through a call and that the user is still connected.

## Record Action

Send a `record` action to start a recording:

```go
    record := ncco.RecordAction{BeepStart: true}
```

When the recording completes, Vonage sends a webhook containing the recording URL so that you can download the file.

## Conversation Action

Adds the call to a conversation:

```go
    conversation := ncco.ConversationAction{Name: "convo1"}
```

## Stream Action

Play an mp3 file into a call as an audio stream:

```go
    stream := ncco.StreamAction{StreamUrl: []string{"https://example.com/music.mp3"}}

```

## Connect Action

Connects the current call to another endpoint (currently only phone is supported):

```go
    endpoint := []ncco.PhoneEndpoint{Number: "44777000777"}
	connect := ncco.ConnectAction{Endpoint: endpoint, From: "44777000888"}
```
The `from` field when connecting to a phone endpoint should be a Vonage number that you own.
