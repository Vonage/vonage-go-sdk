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

func TestAccountGetSecret(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/accounts/12345678/secrets/ad6dc56f-07b5-46e1-a527-85530e625800",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
	"_links": {
	"self": {
		"href": "abc123"
	}
	},
	"id": "ad6dc56f-07b5-46e1-a527-85530e625800",
	"created_at": "2017-03-02T16:34:49Z"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewAccountClient(auth)
	result, _, _ := client.GetSecret("ad6dc56f-07b5-46e1-a527-85530e625800")

	if result.ID != "ad6dc56f-07b5-46e1-a527-85530e625800" {
		t.Error("Test account get one secret failed")
	}
}
func TestAccountGetMissingSecret(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/accounts/12345678/secrets/does-not-exist",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(404, `
{
	"_links": {
	"self": {
		"href": "abc123"
	}
	},
	"id": "ad6dc56f-07b5-46e1-a527-85530e625800",
	"created_at": "2017-03-02T16:34:49Z"
}
	`,
			)

			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewAccountClient(auth)
	_, _, err := client.GetSecret("does-not-exist")

	if err == nil {
		t.Error("Test account get one missing secret failed")
	}
}

func TestAccountCreateSecret(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/accounts/12345678/secrets",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, `
{
    "_links": {
        "self": {
            "href": "/accounts/12345678/secrets/c07c0604-dc5b-4520-94df-dc4964b2fbca"
        }
    },
    "id": "c07c0604-dc5b-4520-94df-dc4964b2fbca",
    "created_at": "2020-10-18T13:37:15Z"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewAccountClient(auth)
	result, _, _ := client.CreateSecret("V3ryS3cr3t!")

	if result.ID != "c07c0604-dc5b-4520-94df-dc4964b2fbca" {
		t.Error("Test account create secret failed")
	}
}

func TestAccountDeleteSecret(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.nexmo.com/accounts/12345678/secrets/ad6dc56f-07b5-46e1-a527-85530e625800",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")

			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewAccountClient(auth)
	result, _, _ := client.DeleteSecret("ad6dc56f-07b5-46e1-a527-85530e625800")

	if result != true {
		t.Error("Test account delete secret failed")
	}
}
