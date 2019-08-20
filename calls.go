package nexmo

import (
	"fmt"
	"net/http"

	"github.com/nexmo-community/nexmo-go/sling"
)

func (c *CallErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", c.ErrorTitle, c.Type)
}

// CreateCall starts a voice call, configured using the provided CreateCallRequest.
func (c *CallService) CreateCall(request CreateCallRequest) (*CreateCallResponse, *http.Response, error) {
	sling := c.sling.New().Post("").BodyJSON(request)

	callResponse := new(CreateCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

// SearchCalls returns information about calls matching the filter in the provided SearchCallsRequest
func (c *CallService) SearchCalls(request SearchCallsRequest) (*SearchCallsResponse, *http.Response, error) {
	sling := c.sling.New().Get("").QueryStruct(request)

	callResponse := new(SearchCallsResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

// Get information about a current or past call by call ID
func (c *CallService) GetCallInfo(uuid string) (*CallInfo, *http.Response, error) {
	sling := c.sling.New().Get(uuid)

	callResponse := new(CallInfo)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

// Change the in-progress call by performing an action, such as hangup, transfer, mute, etc. See the API reference: https://developer.nexmo.com/api/voice#updateCall
func (c *CallService) ModifyCall(uuid string, request interface{}) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Put(uuid).BodyJSON(request)

	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

// Start playing an audio file into a call
func (c *CallService) Stream(uuid string, request StreamRequest) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Put(fmt.Sprintf("%s/stream", uuid)).BodyJSON(request)

	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

// Stop the audio stream from playing in a call
func (c *CallService) StopStream(uuid string) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Delete(fmt.Sprintf("%s/stream", uuid))
	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

// Send text-to-speech into a call
func (c *CallService) Talk(uuid string, request TalkRequest) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Put(fmt.Sprintf("%s/talk", uuid)).BodyJSON(request)

	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

// Stop the text-to-speech that is currently being sent into a call
func (c *CallService) StopTalk(uuid string) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Delete(fmt.Sprintf("%s/talk", uuid))

	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

// Play DTMF tones into a call
func (c *CallService) SendDTMF(uuid string, request DTMFRequest) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Put(fmt.Sprintf("%s/dtmf", uuid)).BodyJSON(request)

	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

func (c *CallService) makeRequest(s *sling.Sling, successV interface{}) (*http.Response, error) {
	errorV := new(CallErrorResponse)
	if err := c.authSet.ApplyJWT(s); err != nil {
		return nil, fmt.Errorf("%s - error applying JWT", err)
	}
	httpResponse, err := s.Receive(successV, errorV)
	if err != nil {
		return httpResponse, fmt.Errorf("%s - error receiving data from server", err)
	}
	if errorV.Type != "" {
		return httpResponse, errorV
	}
	return httpResponse, nil
}
