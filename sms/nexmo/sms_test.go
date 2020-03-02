package nexmo

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestNewNexmoSMSClient(*testing.T) {
	auth := CreateAuthFromKeySecret("123", "456")
	NewNexmoSMSClient(auth)
}

func TestSend(*testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/sms/json",
		httpmock.NewStringResponder(200, `{
				"message-count": "1",
				"messages": [{
					"to": "447520615146",
					"message-id": "140000005494DDEB",
					"status": "0",
					"remaining-balance": "54.42941782",
					"message-price": "0.03330000",
					"network": "23409"
				}]
			}`))

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNexmoSMSClient(auth)
	client.Send("44777000777", "44777000888", "hello", SMSClientOpts{})
}

func TestSendFail(*testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/sms/json",
		httpmock.NewStringResponder(200, `{
				"message-count": "1",
				"messages": [{
					"status": "4",
				}]
			}`))

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNexmoSMSClient(auth)
	_, err := client.Send("44777000777", "44777000888", "hello", SMSClientOpts{})

	if err == nil {
		panic("This should not work")
	}
}
