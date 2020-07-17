package nexmo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTalkSimple(t *testing.T) {
	ncco := Ncco{}
	talk := TalkAction{Text: "Hello"}
	ncco.AddAction(talk)

	// should get one element
	if len(ncco.GetActions()) != 1 {
		t.Errorf("Unexpected number of ncco items")
	}

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"talk\",\"text\":\"Hello\",\"bargeIn\":false,\"loop\":1}]" {
		t.Errorf("Unexpected JSON format for: simple Talk")
	}
}

func TestTalkAll(t *testing.T) {
	ncco := Ncco{}
	talk := TalkAction{Text: "Hello", Loop: "4", BargeIn: true, Level: 1, VoiceName: "Nicole"}
	ncco.AddAction(talk)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"talk\",\"text\":\"Hello\",\"bargeIn\":true,\"level\":1,\"voiceName\":\"Nicole\",\"loop\":4}]" {
		t.Errorf("Unexpected JSON format for: All Talk fields")
	}
}

func TestNotify(t *testing.T) {
	ncco := Ncco{}
	url := []string{"https://example.com/notify"}
	data := make(map[string]string)
	data["key"] = "If that is a key, this is a value"
	data["count"] = "7"
	ping := NotifyAction{EventUrl: url, EventMethod: "POST", Payload: data}
	ncco.AddAction(ping)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"notify\",\"payload\":{\"count\":\"7\",\"key\":\"If that is a key, this is a value\"},\"eventUrl\":[\"https://example.com/notify\"],\"eventMethod\":\"POST\"}]" {
		t.Errorf("Unexpected JSON format for: Notify action")
	}
}

func TestRecordSimple(t *testing.T) {
	ncco := Ncco{}
	record := RecordAction{}
	ncco.AddAction(record)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"record\"}]" {
		t.Errorf("Unexpected JSON format for: simple Record")
	}
}

func TestRecordAll(t *testing.T) {
	ncco := Ncco{}
	url := []string{"https://example.com/record"}
	record := RecordAction{Format: "ogg", Split: "conversation", Channels: 8, EndOnSilence: 5, EndOnKey: "#", TimeOut: 10, BeepStart: true, EventUrl: url, EventMethod: "GET"}
	ncco.AddAction(record)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"record\",\"format\":\"ogg\",\"split\":\"conversation\",\"channels\":8,\"endOnSilence\":5,\"endOnKey\":\"#\",\"timeOut\":10,\"beepStart\":true,\"eventUrl\":[\"https://example.com/record\"],\"eventMethod\":\"GET\"}]" {
		t.Errorf("Unexpected JSON format for: Record all")
	}
}

func TestConversationSimple(t *testing.T) {
	ncco := Ncco{}
	conversation := ConversationAction{Name: "convo1"}
	ncco.AddAction(conversation)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"conversation\",\"name\":\"convo1\",\"startOnEnter\":true}]" {
		t.Errorf("Unexpected JSON format for: simple Conversation")
	}
}

func TestConversationAll(t *testing.T) {
	ncco := Ncco{}
	url := []string{"https://example.com/music.mp3"}
	conversation := ConversationAction{Name: "convo1", MusicOnHoldUrl: url, StartOnEnter: "False", EndOnExit: true, Record: true}
	ncco.AddAction(conversation)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"conversation\",\"name\":\"convo1\",\"musicOnHoldUrl\":[\"https://example.com/music.mp3\"],\"endOnExit\":true,\"record\":true,\"startOnEnter\":false}]" {
		t.Errorf("Unexpected JSON format for: Conversation all")
	}
}

func TestConnectSimplePhone(t *testing.T) {
	ncco := Ncco{}
	endpoint := make([]Endpoint, 1)
	endpoint[0] = PhoneEndpoint{Number: "447770007777"}
	connect := ConnectAction{Endpoint: endpoint, From: "447770008888"}
	ncco.AddAction(connect)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"connect\",\"endpoint\":[{\"type\":\"phone\",\"number\":\"447770007777\"}],\"from\":\"447770008888\"}]" {
		t.Errorf("Unexpected JSON format for: simple Connect phone")
	}
}

func TestConnectAllPhone(t *testing.T) {
	ncco := Ncco{}
	url := []string{"https://example.com/event"}
	endpoint := make([]Endpoint, 1)
	endpoint[0] = PhoneEndpoint{Number: "447770007777", DtmfAnswer: "41"}
	connect := ConnectAction{Endpoint: endpoint, From: "447770008888", Timeout: 3, Limit: 5, MachineDetection: "continue", EventMethod: "GET", EventUrl: url}
	ncco.AddAction(connect)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"connect\",\"endpoint\":[{\"type\":\"phone\",\"number\":\"447770007777\",\"dtmfAnswer\":\"41\"}],\"from\":\"447770008888\",\"limit\":5,\"machineDetection\":\"continue\",\"eventUrl\":[\"https://example.com/event\"],\"eventMethod\":\"GET\"}]" {
		t.Errorf("Unexpected JSON format for: simple Connect phone")
	}
}
