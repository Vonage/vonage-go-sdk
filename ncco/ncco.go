package ncco

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Action is an interface to ensure all Actions have a prepare()
type Action interface {
	prepare() Action
}

// Ncco is a parent type to hold the actions
type Ncco struct {
	actions []interface{}
}

// AddAction to add a next action, in sequence, to the Ncco
func (n *Ncco) AddAction(action Action) {
	n.actions = append(n.actions, action.prepare())
}

// GetActions to return the NCCO array, ready to be JSON Marshalled
// This calls the prepare() actions for any additional transforms needed
func (n *Ncco) GetActions() []interface{} {
	return n.actions
}

// MarshalJSON to return the NCCO array, ready to be JSON Marshalled
func (n Ncco) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.GetActions())
}

// TalkAction is a text-to-speech feature. Beware that the "Loop"
// field is a string here, and the CalculatedLoopValue is used to
// assemble the correct value when sending
type TalkAction struct {
	Action              string `json:"action"`
	Text                string `json:"text"`
	Loop                string `json:"-"`
	BargeIn             bool   `json:"bargeIn"`
	Level               int    `json:"level,omitempty"`
	VoiceName           string `json:"voiceName,omitempty"`
	Style               int    `json:"style,omitempty"`
	Language            string `json:"language,omitempty"`
	CalculatedLoopValue int    `json:"loop"`
}

// prepare for TalkAction calculates a loop value
func (a TalkAction) prepare() Action {
	a.Action = "talk"
	// int fields default to 0, which isn't the default for looping
	// look at the value of string Loop, and set field for JSON if it's a number
	a.CalculatedLoopValue = 1
	if a.Loop != "" {
		loop, err := strconv.Atoi(a.Loop)
		if err == nil {
			a.CalculatedLoopValue = loop
		}
	}
	return a
}

// NotifyAction to represent a Notify Action
type NotifyAction struct {
	Action      string            `json:"action"`
	Payload     map[string]string `json:"payload,omitempty"`
	EventUrl    []string          `json:"eventUrl,omitempty"`
	EventMethod string            `json:"eventMethod,omitempty"`
}

// prepare for the NotifyAction
func (a NotifyAction) prepare() Action {
	a.Action = "notify"
	return a
}

// RecordAction to start a recording at this point in the call
type RecordAction struct {
	Action       string   `json:"action"`
	Format       string   `json:"format,omitempty"`
	Split        string   `json:"split,omitempty"`
	Channels     int      `json:"channels,omitempty"`
	EndOnSilence int      `json:"endOnSilence,omitempty"`
	EndOnKey     string   `json:"endOnKey,omitempty"`
	TimeOut      int      `json:"timeOut,omitempty"`
	BeepStart    bool     `json:"beepStart,omitempty"`
	EventUrl     []string `json:"eventUrl,omitempty"`
	EventMethod  string   `json:"eventMethod,omitempty"`
}

// prepare for the RecordAction
func (a RecordAction) prepare() Action {
	a.Action = "record"
	return a
}

// ConversationAction sets up a conference that calls can be added to
type ConversationAction struct {
	Action                      string   `json:"action"`
	Name                        string   `json:"name,omitempty"`
	MusicOnHoldUrl              []string `json:"musicOnHoldUrl,omitempty"`
	StartOnEnter                string   `json:"-"`
	EndOnExit                   bool     `json:"endOnExit,omitempty"`
	Record                      bool     `json:"record,omitempty"`
	CanSpeak                    []string `json:"canSpeak,omitempty"`
	CanHear                     []string `json:"canHear,omitempty"`
	CalculatedStartOnEnterValue bool     `json:"startOnEnter"`
}

// prepare for the ConversationAction
func (a ConversationAction) prepare() Action {
	a.Action = "conversation"
	// boolean fields default to false, but startOnEnter defaults to true
	// look at the value of string StartOnEnter, set false if string "false" is given
	a.CalculatedStartOnEnterValue = true
	if strings.ToLower(a.StartOnEnter) == "false" {
		a.CalculatedStartOnEnterValue = false
	}
	return a
}

// StreamAction plays audio stream from URL. Beware that the "Loop"
// field is a string here, and the CalculatedLoopValue is used to
// assemble the correct value when sending
type StreamAction struct {
	Action              string   `json:"action"`
	StreamUrl           []string `json:"streamUrl,omitempty"`
	Level               int      `json:"level,omitempty"`
	Loop                string   `json:"-"`
	BargeIn             bool     `json:"bargeIn"`
	CalculatedLoopValue int      `json:"loop"`
}

// prepare for StreamAction calculates a loop value
func (a StreamAction) prepare() Action {
	a.Action = "stream"
	// int fields default to 0, which isn't the default for looping
	// look at the value of string Loop, and set field for JSON if it's a number
	a.CalculatedLoopValue = 1
	if a.Loop != "" {
		loop, err := strconv.Atoi(a.Loop)
		if err == nil {
			a.CalculatedLoopValue = loop
		}
	}
	return a
}

// InputAction uses pointers for the optional dtmf input
type InputAction struct {
	Action      string     `json:"action"`
	Dtmf        *DtmfInput `json:"dtmf,omitempty"`
	EventUrl    []string   `json:"eventUrl,omitempty"`
	EventMethod string     `json:"eventMethod,omitempty"`
}

// prepare for the InputAction
func (a InputAction) prepare() Action {
	a.Action = "input"
	return a
}

// DtmfInput captures digits pressed on the keypad
type DtmfInput struct {
	TimeOut      int  `json:"timeOut,omitempty"`
	MaxDigits    int  `json:"maxDigits,omitempty"`
	SubmitOnHash bool `json:"submitOnHash,omitempty"`
}

// ConnectAction takes an Endpoint (of which there are many) and joins
// it into the current call
type ConnectAction struct {
	Action           string     `json:"action"`
	Endpoint         []Endpoint `json:"endpoint"`
	From             string     `json:"from,omitempty"`
	Timeout          int        `json:"timeout,omitempty"`
	Limit            int        `json:"limit,omitempty"`
	MachineDetection string     `json:"machineDetection,omitempty"`
	EventType        string     `json:"eventType,omitempty"`
	EventUrl         []string   `json:"eventUrl,omitempty"`
	EventMethod      string     `json:"eventMethod,omitempty"`
	RingbackTone     string     `json:"ringbackTone,omitempty"`
}

// prepare for the ConnectAction
func (a ConnectAction) prepare() Action {
	a.Action = "connect"
	// organise the endpoint
	a.Endpoint[0] = a.Endpoint[0].prepareEndpoint()
	return a
}

//--- Connect Endpoints ---

// Endpoint is a mostly dummy interface to let us typehint on it
type Endpoint interface {
	prepareEndpoint() Endpoint
}

type PhoneEndpoint struct {
	Type       string `json:"type"`
	Number     string `json:"number"`
	DtmfAnswer string `json:"dtmfAnswer,omitempty"`
	OnAnswer   string `json:"onAnswer,omitempty"`
}

func (e PhoneEndpoint) prepareEndpoint() Endpoint {
	e.Type = "phone"
	return e
}
