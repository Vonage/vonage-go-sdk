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
	if err := c.authSet.ApplyJWT(sling); err != nil {
		return nil, nil, fmt.Errorf("%s - error applying JWT", err)
	}

	callResponse := new(CreateCallResponse)
	errorResponse := new(CallErrorResponse)
	httpResponse, err := sling.Receive(callResponse, errorResponse)
	if err != nil {
		return nil, httpResponse, fmt.Errorf("%s - error receiving data from server", err)
	}
	if errorResponse.Type != "" {
		return nil, httpResponse, errorResponse
	}
	return callResponse, httpResponse, nil
}

// SearchCalls returns information about calls matching the filter in the provided SearchCallsRequest
func (c *CallService) SearchCalls(request SearchCallsRequest) (*SearchCallsResponse, *http.Response, error) {
	sling := c.sling.New().Get("").QueryStruct(request)
	if err := c.authSet.ApplyJWT(sling); err != nil {
		return nil, nil, fmt.Errorf("%s - error applying JWT", err)
	}

	callResponse := new(SearchCallsResponse)
	errorResponse := new(CallErrorResponse)
	httpResponse, err := sling.Receive(callResponse, errorResponse)
	if err != nil {
		return nil, httpResponse, fmt.Errorf("%s - error receiving data from server", err)
	}
	if errorResponse.Type != "" {
		return nil, httpResponse, errorResponse
	}
	return callResponse, httpResponse, nil
}

func (c *CallService) GetCallInfo(uuid string) (*CallInfo, *http.Response, error) {
	sling := c.sling.New().Get(uuid)
	if err := c.authSet.ApplyJWT(sling); err != nil {
		return nil, nil, fmt.Errorf("%s - error applying JWT", err)
	}

	callResponse := new(CallInfo)
	errorResponse := new(CallErrorResponse)
	httpResponse, err := sling.Receive(callResponse, errorResponse)
	if err != nil {
		return nil, httpResponse, fmt.Errorf("%s - error receiving data from server", err)
	}
	if errorResponse.Type != "" {
		return nil, httpResponse, errorResponse
	}
	return callResponse, httpResponse, nil
}

func (c *CallService) ModifyCall(uuid string, request interface{}) (*ModifyCallResponse, *http.Response, error) {
	sling := c.sling.New().Put(uuid).BodyJSON(request)
	if err := c.authSet.ApplyJWT(sling); err != nil {
		return nil, nil, fmt.Errorf("%s - error applying JWT", err)
	}

	callResponse := new(ModifyCallResponse)
	errorResponse := new(CallErrorResponse)
	httpResponse, err := sling.Receive(callResponse, errorResponse)
	if err != nil {
		return nil, httpResponse, fmt.Errorf("%s - error receiving data from server", err)
	}
	if errorResponse.Type != "" {
		return nil, httpResponse, errorResponse
	}
	return callResponse, httpResponse, nil
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
