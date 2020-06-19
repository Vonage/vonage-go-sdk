package nexmo

import "strconv"

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

// GetActions to get all the actions
func (n *Ncco) GetActions() []interface{} {
	return n.actions
}

// TalkAction is a text-to-speech feature. Beware that the "Loop"
// field is a string here, and the CalculatedLoopValue is used to
// assemble the correct value when sending
type TalkAction struct {
	Action              string `json:"action,omitempty"`
	Text                string `json:"text"`
	Loop                string `json:"-"`
	BargeIn             bool   `json:"bargeIn"`
	Level               int    `json:"level"`
	VoiceName           string `json:"voiceName,omitempty"`
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
	Action      string            `json:"action,omitempty"`
	Payload     map[string]string `json:"payload,omitempty"`
	EventUrl    []string          `json:"eventUrl,omitempty"`
	EventMethod string            `json:"eventMethod,omitempty"`
}

// prepare for the NotifyAction
func (a NotifyAction) prepare() Action {
	a.Action = "notify"
	return a
}
