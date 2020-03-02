package nexmo

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestNewNexmoSMSClient(*testing.T) {
	auth := CreateAuthFromKeySecret("123", "456")
	NewNexmoSMSClient(auth)
}

func TestSend(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/sms/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
	{
	  "message-count": "1",
	  "messages": [
	    {
	      "to": "447700900000",
	      "message-id": "0A0000000123ABCD1",
	      "status": "0",
	      "remaining-balance": "3.14159265",
	      "message-price": "0.03330000",
	      "network": "12345",
	      "account-ref": "customer1234"
	    }
	  ]
	}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNexmoSMSClient(auth)
	_, err := client.Send("44777000777", "44777000888", "hello", SMSClientOpts{})

	if err != nil {
		t.Error("Test SMS not sent")
	}
}

func TestSendFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/sms/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
	{
	  "message-count": "1",
	  "messages": [
	    {
	      "status": "4"
	    }
	  ]
	}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNexmoSMSClient(auth)
	_, err := client.Send("44777000777", "44777000888", "hello", SMSClientOpts{})

	if err == nil {
		t.Error("The failure failed")
	}
}
