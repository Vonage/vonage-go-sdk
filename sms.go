package nexmo

import (
	"net/http"
)

// Send an SMS message
func (c *SMSService) SendSMS(request SendSMSRequest) (*SendSMSResponse, *http.Response, error) {
	c.authSet.ApplyAPICredentials(&request)
	sling := c.sling.New().Post("json").BodyJSON(request)

	callResponse := new(SendSMSResponse)
	httpResponse, err := sling.ReceiveSuccess(callResponse)
	return callResponse, httpResponse, err
}
