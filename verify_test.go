package vonage

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestVerifyNewVerifyClient(*testing.T) {
	auth := CreateAuthFromKeySecret("123", "456")
	NewVerifyClient(auth)
}

func TestVerifyRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/json",
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
	response, _, _ := client.Request("44777000777", "VonageGoTest", VerifyOpts{})

	message := "Request ID: " + response.RequestId
	if message != "Request ID: abcdef0123456789abcdef0123456789" {
		t.Errorf("Verify request failed")
	}
}

func TestVerifyRequestConcurrent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{

    "request_id": "abcdef0123456789abcdef0123456789",
    "status": "10",
    "error_text": "Concurrent verifications to the same number are not allowed"

}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	_, errResp, _ := client.Request("44777000777", "VonageGoTest", VerifyOpts{})

	message := "Error status " + errResp.Status + ": " + errResp.ErrorText
	if message != "Error status 10: Concurrent verifications to the same number are not allowed" {
		t.Error("Unexpected error response")
	}
}

func TestVerifyRequestFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(401, `
Go away
	`,
			)

			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	_, _, err := client.Request("44777000777", "VonageGoTest", VerifyOpts{})

	if err == nil {
		t.Errorf("This should have produced an error")
	}

}

func TestVerifyCheck(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/check/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "request_id": "abcdef0123456789abcdef0123456789",
  "event_id": "0A00000012345678",
  "status": "0",
  "price": "0.10000000",
  "currency": "EUR",
  "estimated_price_messages_sent": "0.03330000"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	response, _, _ := client.Check("abcdef0123456789abcdef0123456789", "9876")

	message := "Request ID: " + response.RequestId
	if message != "Request ID: abcdef0123456789abcdef0123456789" {
		t.Errorf("Verify check failed")
	}
}

func TestVerifyCheckError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/check/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{

    "status": "16",
    "error_text": "The code inserted does not match the expected value"

}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	_, errResp, _ := client.Check("abcdef0123456789abcdef0123456789", "9876")

	message := "Error status " + errResp.Status + ": " + errResp.ErrorText
	if message != "Error status 16: The code inserted does not match the expected value" {
		t.Error("Unexpected error response")
	}
}

func TestVerifySearch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/verify/search/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{

    "request_id": "abcdef0123456789abcdef0123456789",
    "account_id": "abcdef01",
    "number": "447700900000",
    "sender_id": "verify",
    "date_submitted": "2020-01-01 12:00:00",
    "date_finalized": "2020-01-01 12:00:00",
    "checks":

[

{

    "date_received": "2020-01-01 12:00:00",
    "code": "1111",
    "status": "INVALID",
    "ip_address": ""

},

    {
        "date_received": "2020-01-01 12:02:00",
        "code": "1234",
        "status": "VALID",
        "ip_address": ""
    }

],
"first_event_date": "2020-01-01 12:00:00",
"last_event_date": "2020-01-01 12:00:00",
"price": "0.1000000",
"currency": "EUR",
"status": "SUCCESS",
"estimated_price_messages_sent": "0.06660000",
"events":
[

{

    "id": "0A00000012345678",
    "type": "sms"

},

        {
            "id": "0A00000012345679",
            "type": "sms"
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
	client := NewVerifyClient(auth)
	response, _, _ := client.Search("abcdef0123456789abcdef0123456789")

	message := "Request ID: " + response.RequestId
	if message != "Request ID: abcdef0123456789abcdef0123456789" {
		t.Errorf("Verify request failed")
	}
}

func TestVerifySearchError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/verify/search/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "status": "101",
  "error_text": "No response found"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	_, errResp, _ := client.Search("abcdef0123456789abcdef0123456789")

	message := "Error status " + errResp.Status + ": " + errResp.ErrorText
	if message != "Error status 101: No response found" {
		t.Error("Unexpected error response")
	}
}

func TestVerifyCancel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/control/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "status": "0",
  "command": "cancel"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	response, _, _ := client.Cancel("abcdef0123456789abcdef0123456789")

	message := "Status: " + response.Status
	if message != "Status: 0" {
		t.Errorf("Verify cancel failed")
	}
}

func TestVerifyCancelTooSoon(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/control/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
    "status": "19",
    "error_text": "Verification request ['abcdef0123456789abcdef0123456789'] can't be cancelled within the first 30 seconds."

}`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	_, errResp, _ := client.Cancel("abcdef0123456789abcdef0123456789")

	message := "Status: " + errResp.Status
	if message != "Status: 19" {
		t.Errorf("Verify cancel 'too soon' failed")
	}
}

func TestVerifyCancelNotNow(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/control/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "status": "19",
  "error_text": "Verification request  ['abcdef0123456789abcdef0123456789'] can't be cancelled now. Too many attempts to re-deliver have already been made."
}
`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	_, errResp, _ := client.Cancel("abcdef0123456789abcdef0123456789")

	message := "Status: " + errResp.Status
	if message != "Status: 19" {
		t.Errorf("Verify cancel 'not now' failed")
	}
}

func TestVerifyTriggerNextEvent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/control/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "status": "0",
  "command": "trigger_next_event"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	response, _, _ := client.TriggerNextEvent("abcdef0123456789abcdef0123456789")

	message := "Status: " + response.Status
	if message != "Status: 0" {
		t.Errorf("Verify trigger next event failed")
	}
}

func TestVerifyTriggerNextEventFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/control/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
    "status": "1",
    "error_text": "Throttled"
}
`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewVerifyClient(auth)
	_, errResp, _ := client.Cancel("abcdef0123456789abcdef0123456789")

	message := "Status: " + errResp.Status
	if message != "Status: 1" {
		t.Errorf("Verify throttled trigger next event failed")
	}
}
