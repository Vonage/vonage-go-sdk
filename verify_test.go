package nexmo

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestNewVerifyClient(*testing.T) {
	auth := CreateAuthFromKeySecret("123", "456")
	NewVerifyClient(auth)
}

func TestRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/verify/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
	{
		"request_id": "abcdef0123456789abcdef0123456789",
		"status": "0"
	}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	_, err := client.Request("44777000777", "NexmoGoTest")

	if err != nil {
		t.Error("Verification failed")
	}
}
