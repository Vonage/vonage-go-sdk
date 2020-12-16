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

func TestAccountSetConfig(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/account/settings",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "mo-callback-url": "https://example.com/webhooks/inbound-sms",
  "dr-callback-url": "https://example.com/webhooks/delivery-receipt",
  "max-outbound-request": 30,
  "max-inbound-request": 30,
  "max-calls-per-second": 30
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewAccountClient(auth)
	result, _, _ := client.SetConfig(AccountConfigSettings{})

	if result.MoCallbackUrl != "https://example.com/webhooks/inbound-sms" {
		t.Error("Test account set config failed")
	}
}

func TestAccountSetConfigNoAuth(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/account/settings",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(420, `
{"max-outbound-request":0,"max-inbound-request":0,"max-calls-per-second":0,"error-code":"420","error-code-label":"API key is required"}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("", "")
	client := NewAccountClient(auth)
	_, resp, _ := client.SetConfig(AccountConfigSettings{})

	if resp.ErrorCode != "420" {
		t.Error("Test account set config missing auth behaviour failed")
	}
}

func TestAccountListSecrets(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/accounts/12345678/secrets",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "_links": {
    "self": {
      "href": "abc123"
    }
  },
  "_embedded": {
    "secrets": [
      {
        "_links": {
          "self": {
            "href": "abc123"
          }
        },
        "id": "ad6dc56f-07b5-46e1-a527-85530e625800",
        "created_at": "2017-03-02T16:34:49Z"
      }
    ]
  }
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewAccountClient(auth)
	result, _, _ := client.ListSecrets()

	if result.Secrets[0].ID != "ad6dc56f-07b5-46e1-a527-85530e625800" {
		t.Error("Test account list secrets failed")
	}
}
