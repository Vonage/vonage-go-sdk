package nexmo

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
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

func (c *CallService) GetCallInfo(uuid string) (*CallInfo, *http.Response, error) {
	sling := c.sling.New().Get(uuid)

	callResponse := new(CallInfo)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

func (c *CallService) ModifyCall(uuid string, request interface{}) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Put(uuid).BodyJSON(request)

	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

func (c *CallService) Stream(uuid string, request StreamRequest) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Put(fmt.Sprintf("%s/stream", uuid)).BodyJSON(request)

	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

func (c *CallService) StopStream(uuid string) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Delete(fmt.Sprintf("%s/stream", uuid))
	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

func (c *CallService) Talk(uuid string, request TalkRequest) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Put(fmt.Sprintf("%s/talk", uuid)).BodyJSON(request)

	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

func (c *CallService) StopTalk(uuid string) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Delete(fmt.Sprintf("%s/talk", uuid))
	callResponse := new(ModifyCallResponse)
	httpResponse, err := c.makeRequest(sling, callResponse)
	return callResponse, httpResponse, err
}

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
