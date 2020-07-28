package nexmo

import (
	"encoding/json"
	"testing"
)

func TestNccoTalkSimple(t *testing.T) {
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

func TestNccoTalkAll(t *testing.T) {
	ncco := Ncco{}
	talk := TalkAction{Text: "Hello", Loop: "4", BargeIn: true, Level: 1, VoiceName: "Nicole"}
	ncco.AddAction(talk)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"talk\",\"text\":\"Hello\",\"bargeIn\":true,\"level\":1,\"voiceName\":\"Nicole\",\"loop\":4}]" {
		t.Errorf("Unexpected JSON format for: All Talk fields")
	}
}

func TestNccoNotify(t *testing.T) {
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

func TestNccoRecordSimple(t *testing.T) {
	ncco := Ncco{}
	record := RecordAction{}
	ncco.AddAction(record)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"record\"}]" {
		t.Errorf("Unexpected JSON format for: simple Record")
	}
}

func TestNccoRecordAll(t *testing.T) {
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

func TestNccoConversationSimple(t *testing.T) {
	ncco := Ncco{}
	conversation := ConversationAction{Name: "convo1"}
	ncco.AddAction(conversation)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"conversation\",\"name\":\"convo1\",\"startOnEnter\":true}]" {
		t.Errorf("Unexpected JSON format for: simple Conversation")
	}
}

func TestNccoConversationAll(t *testing.T) {
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

func TestNccoConnectSimplePhone(t *testing.T) {
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

func TestNccoConnectAllPhone(t *testing.T) {
	ncco := Ncco{}
	url := []string{"https://example.com/event"}
	endpoint := make([]Endpoint, 1)
	endpoint[0] = PhoneEndpoint{Number: "447770007777", DtmfAnswer: "41"}
	connect := ConnectAction{Endpoint: endpoint, From: "447770008888", Timeout: 3, Limit: 5, MachineDetection: "continue", EventMethod: "GET", EventUrl: url}
	ncco.AddAction(connect)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"connect\",\"endpoint\":[{\"type\":\"phone\",\"number\":\"447770007777\",\"dtmfAnswer\":\"41\"}],\"from\":\"447770008888\",\"timeout\":3,\"limit\":5,\"machineDetection\":\"continue\",\"eventUrl\":[\"https://example.com/event\"],\"eventMethod\":\"GET\"}]" {
		t.Errorf("Unexpected JSON format for: simple Connect phone")
	}
}

func TestNccoStreamSimple(t *testing.T) {
	ncco := Ncco{}
	stream := StreamAction{StreamUrl: []string{"https://example.com/music.mp3"}}
	ncco.AddAction(stream)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"stream\",\"streamUrl\":[\"https://example.com/music.mp3\"],\"bargeIn\":false,\"loop\":1}]" {
		t.Errorf("Unexpected JSON format for: simple Stream")
	}
}

func TestNccoStreamOptions(t *testing.T) {
	ncco := Ncco{}
	stream := StreamAction{StreamUrl: []string{"https://example.com/music.mp3"}, Level: 1, Loop: "4", BargeIn: true}
	ncco.AddAction(stream)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"stream\",\"streamUrl\":[\"https://example.com/music.mp3\"],\"level\":1,\"bargeIn\":true,\"loop\":4}]" {
		t.Errorf("Unexpected JSON format for: Stream options")
	}
}

func TestNccoInputDtmfEmpty(t *testing.T) {
	ncco := Ncco{}
	dtmf := InputAction{Dtmf: &DtmfInput{}}
	ncco.AddAction(dtmf)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	if string(j) != "[{\"action\":\"input\",\"dtmf\":{}}]" {
		t.Errorf("Unexpected JSON format for: Input DTMF empty")
	}
}

func TestNccoInputDtmfOptions(t *testing.T) {
	ncco := Ncco{}
	dtmf := InputAction{EventUrl: []string{"https://example.com/event"}, EventMethod: "GET", Dtmf: &DtmfInput{TimeOut: 8, MaxDigits: 2, SubmitOnHash: true}}
	ncco.AddAction(dtmf)

	// check the JSON
	j, _ := json.Marshal(ncco.GetActions())
	// fmt.Println(string(j))
	if string(j) != "[{\"action\":\"input\",\"dtmf\":{\"timeOut\":8,\"maxDigits\":2,\"submitOnHash\":true},\"eventUrl\":[\"https://example.com/event\"],\"eventMethod\":\"GET\"}]" {
		t.Errorf("Unexpected JSON format for: Input DTMF options")
	}
}
