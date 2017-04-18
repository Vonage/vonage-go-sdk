package nexmo

import (
	"fmt"
	"net/http"
)

func (c *CallErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", c.ErrorTitle, c.Type)
}

func (c *CallService) CreateCall(request CreateCallRequest) (*CreateCallResponse, *http.Response, error) {
	sling := c.sling.New().Post("").BodyJSON(request)
	err := c.authSet.ApplyJWT(sling)
	if err != nil {
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

func (c *CallService) SearchCalls(request SearchCallsRequest) (*SearchCallsResponse, *http.Response, error) {
	sling := c.sling.New().Get("").QueryStruct(request)
	err := c.authSet.ApplyJWT(sling)
	if err != nil {
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
