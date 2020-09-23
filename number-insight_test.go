package vonage

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestNINewClient(*testing.T) {
	auth := CreateAuthFromKeySecret("123", "456")
	NewNumberInsightClient(auth)
}

func TestNIBasic(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/ni/basic/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "status": 0,
  "status_message": "Success",
  "request_id": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab",
  "international_format_number": "447700900000",
  "national_format_number": "07700 900000",
  "country_code": "GB",
  "country_code_iso3": "GBR",
  "country_name": "United Kingdom",
  "country_prefix": "44"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumberInsightClient(auth)
	response, _, _ := client.Basic("44777000777", NiOpts{})

	message := "Status: " + response.StatusMessage
	if message != "Status: Success" {
		t.Errorf("Number insight basic lookup failed")
	}
}

func TestNIBasicWithOpts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/ni/basic/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "status": 0,
  "status_message": "Success",
  "request_id": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab",
  "international_format_number": "447700900000",
  "national_format_number": "07700 900000",
  "country_code": "GB",
  "country_code_iso3": "GBR",
  "country_name": "United Kingdom",
  "country_prefix": "44"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumberInsightClient(auth)
	response, _, _ := client.Basic("44777000777", NiOpts{Country: "US"})

	message := "Status: " + response.StatusMessage
	if message != "Status: Success" {
		t.Errorf("Number insight basic lookup failed")
	}
}

func TestNIAsync(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/ni/advanced/async/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "request_id": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab",
  "number": "447700900000",
  "remaining_balance": "1.23456789",
  "request_price": "0.01500000",
  "status": 0,
  "error_text": "Success"
}

	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumberInsightClient(auth)
	response, _, _ := client.AdvancedAsync("44777000777", "https://example.com/webhook/ni", NiOpts{Country: "US"})

	message := "Price: " + response.RequestPrice
	if message != "Price: 0.01500000" {
		t.Errorf("Number insight advanced async request failed")
	}
}

func TestNIStandard(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/ni/standard/json",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "status": 0,
  "status_message": "Success",
  "request_id": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab",
  "international_format_number": "447700900000",
  "national_format_number": "07700 900000",
  "country_code": "GB",
  "country_code_iso3": "GBR",
  "country_name": "United Kingdom",
  "country_prefix": "44",
  "request_price": "0.04000000",
  "refund_price": "0.01500000",
  "remaining_balance": "1.23456789",
  "current_carrier": {
    "network_code": "12345",
    "name": "Acme Inc",
    "country": "GB",
    "network_type": "mobile"
  },
  "original_carrier": {
    "network_code": "12345",
    "name": "Acme Inc",
    "country": "GB",
    "network_type": "mobile"
  },
  "ported": "not_ported",
  "roaming": {
    "status": "roaming",
    "roaming_country_code": "US",
    "roaming_network_code": "12345",
    "roaming_network_name": "Acme Inc"
  },
  "caller_identity": {
    "caller_type": "consumer",
    "caller_name": "John Smith",
    "first_name": "John",
    "last_name": "Smith"
  },
  "caller_name": "John Smith",
  "last_name": "Smith",
  "first_name": "John",
  "caller_type": "consumer"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewNumberInsightClient(auth)
	response, _, _ := client.Standard("44777000777", NiOpts{Country: "US"})

	message := "Price: " + response.RequestPrice
	if message != "Price: 0.04000000" {
		t.Errorf("Number insight standard lookup failed")
	}
}
