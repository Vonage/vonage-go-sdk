package nexmo

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestNewNumbersClient(*testing.T) {
	auth := CreateAuthFromKeySecret("123", "456")
	NewNumbersClient(auth)
}

func TestList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://rest.nexmo.com/account/numbers",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "count": 1,
  "numbers": [
    {
      "country": "GB",
      "msisdn": "447700900000",
      "moHttpUrl": "https://example.com/webhooks/inbound-sms",
      "type": "mobile-lvn",
      "features": [
        "VOICE",
        "SMS"
      ],
      "messagesCallbackType": "app",
      "messagesCallbackValue": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab",
      "voiceCallbackType": "app",
      "voiceCallbackValue": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab"
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
	client := NewNumbersClient(auth)
	response, _ := client.List(NumbersOpts{})

	for _, number := range response.Numbers {
		message := "Number: " + number.Msisdn
		if message != "Number: 447700900000" {
			t.Errorf("Number list failed")
		}
	}
}

func TestListNone(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://rest.nexmo.com/account/numbers",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumbersClient(auth)
	response, _ := client.List(NumbersOpts{})

	if response.Count > 0 {
		t.Errorf("Empty number list failed")
	}
}

func TestFilteredList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://rest.nexmo.com/account/numbers",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "count": 1,
  "numbers": [
    {
      "country": "GB",
      "msisdn": "447700900000",
      "moHttpUrl": "https://example.com/webhooks/inbound-sms",
      "type": "mobile-lvn",
      "features": [
        "VOICE",
        "SMS"
      ],
      "messagesCallbackType": "app",
      "messagesCallbackValue": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab",
      "voiceCallbackType": "app",
      "voiceCallbackValue": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab"
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
	client := NewNumbersClient(auth)
	opts := NumbersOpts{
		HasApplication: "false",
		Pattern:        "447",
		SearchPattern:  0,
	}
	response, _ := client.List(opts)

	for _, number := range response.Numbers {
		message := "Number: " + number.Msisdn
		if message != "Number: 447700900000" {
			t.Errorf("Number list failed")
		}
	}
}

func TestNumberSearchOptions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://rest.nexmo.com/number/search",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "count": 1234,
  "numbers": [
    {
      "country": "GB",
      "msisdn": "447700900000",
      "type": "mobile-lvn",
      "cost": "1.25",
      "features": [
        "VOICE",
        "SMS"
      ]
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
	client := NewNumbersClient(auth)
	opts := NumberSearchOpts{Index: 1, Size: 1, Pattern: "900", SearchPattern: 1, Features: "SMS"}
	response, _ := client.Search("GB", opts)

	for _, number := range response.Numbers {
		message := "Number: " + number.Msisdn
		if message != "Number: 447700900000" {
			t.Errorf("Number search failed")
		}
	}
}

func TestNumberSearch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://rest.nexmo.com/number/search",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "count": 1234,
  "numbers": [
    {
      "country": "GB",
      "msisdn": "447700900000",
      "type": "mobile-lvn",
      "cost": "1.25",
      "features": [
        "VOICE",
        "SMS"
      ]
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
	client := NewNumbersClient(auth)
	opts := NumberSearchOpts{}
	response, _ := client.Search("ES", opts)

	for _, number := range response.Numbers {
		message := "Number: " + number.Msisdn
		if message != "Number: 447700900000" {
			t.Errorf("Number search failed")
		}
	}
}

func TestNumberBuy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/number/buy",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "error-code": "200",
  "error-code-label": "success"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumbersClient(auth)
	opts := NumberBuyOpts{}
	response, _, _ := client.Buy("ES", "44770080000", opts)

	message := "Result: " + response.ErrorCodeLabel
	if message != "Result: success" {
		t.Errorf("Number buy failed")
	}
}

func TestNumberCannotBuy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/number/buy",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "error-code": "420",
  "error-code-label": "Numbers from this country can be requested from the Dashboard (https://dashboard.nexmo.com/buy-numbers) as they require a valid local address to be provided before being purchased."
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumbersClient(auth)
	opts := NumberBuyOpts{}
	response, _, _ := client.Buy("ES", "44770080000", opts)

	message := "Status: " + response.ErrorCode
	if message != "Status: 420" {
		t.Errorf("Number cannot buy failed")
	}
}

func TestNumberCancel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/number/cancel",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "error-code": "200",
  "error-code-label": "success"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumbersClient(auth)
	opts := NumberCancelOpts{}
	response, _, _ := client.Cancel("ES", "44770080000", opts)

	message := "Result: " + response.ErrorCodeLabel
	if message != "Result: success" {
		t.Errorf("Number cancel failed")
	}
}

func TestNumberCancelFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/number/cancel",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "error-code": "420",
  "error-code-label": "method failed"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumbersClient(auth)
	opts := NumberCancelOpts{}
	response, _, _ := client.Cancel("ES", "44770080000", opts)

	message := "Status: " + response.ErrorCode
	if message != "Status: 420" {
		t.Errorf("Number cannot cancel failed")
	}
}

func TestNumberUpdate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/number/update",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "error-code": "200",
  "error-code-label": "success"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumbersClient(auth)
	opts := NumberUpdateOpts{}
	response, _, _ := client.Update("GB", "44770080000", opts)

	message := "Result: " + response.ErrorCodeLabel
	if message != "Result: success" {
		t.Errorf("Number update failed")
	}
}

func TestNumberUpdateFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/number/update",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "error-code": "420",
  "error-code-label": "method failed"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumbersClient(auth)
	opts := NumberUpdateOpts{}
	response, _, _ := client.Update("GB", "44770080000", opts)

	message := "Status: " + response.ErrorCode
	if message != "Status: 420" {
		t.Errorf("Number cannot update failed")
	}
}
