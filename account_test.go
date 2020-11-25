package vonage

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAccountNewAccountClient(*testing.T) {
	auth := CreateAuthFromKeySecret("123", "456")
	NewAccountClient(auth)
}

func TestAccountGetBalance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://rest.nexmo.com/account/get-balance",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "value": 10.28,
  "autoReload": false
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewAccountClient(auth)
	result, _, _ := client.GetBalance()

	balance := result.Value
	if balance != 10.28 {
		t.Error("Test account get balance failed")
	}
}
