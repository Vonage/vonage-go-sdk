package nexmo

import (
	"net/http"
)

func (c *SMSService) SendSMS(request SendSMSRequest) (*SendSMSResponse, *http.Response, error) {
	c.authSet.ApplyAPICredentials(&request)
	sling := c.sling.New().Post("json").BodyJSON(request)

	callResponse := new(SendSMSResponse)
	httpResponse, err := sling.ReceiveSuccess(callResponse)
	return callResponse, httpResponse, err
}
