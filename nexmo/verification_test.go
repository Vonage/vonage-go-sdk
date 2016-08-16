package nexmo

import "testing"

func TestParseVerifyResponseSuccess(t *testing.T) {
	response, err := parseVerifyResponse([]byte(`{"request_id":"e80c552b22054bbc8b95e521520b0e1c","status":"0"}`))
	if err != nil {
		t.Errorf("err was not nil, was %s instead", err)
		return
	}
	if response == nil {
		t.Errorf("Response was empty!")
		return
	}
	if response.RequestID != "e80c552b22054bbc8b95e521520b0e1c" {
		t.Errorf("Returned RequestID was incorrect: %s", response)
		return
	}
}

func TestParseVerifyResponseFailure(t *testing.T) {
	response, err := parseVerifyResponse([]byte(`{"status":"2","error_text":"Missing username"}`))
	if err != nil {
		t.Errorf("err was not nil! Instead it was %s", err)
	}
	if response == nil {
		t.Errorf("response was empty!")
		return
	}

	if response.Status != "2" {
		t.Errorf("response.Status should have been \"2\". Instead it was %s", response.Status)
	}

	if response.ErrorText != "Missing username" {
		t.Errorf("response.ErrorText had the wrong value: %s", response.ErrorText)
	}

}

func TestParseVerifySearchResponseSuccess(t *testing.T) {
	bytes := []byte(`{
	"request_id": "ad883d3ba753473694d9d6c70f529124",
	"account_id": "accountID",
	"number": "447720444444",
	"sender_id": "verify",
	"date_submitted": "2016-08-14 10:45:26",
	"date_finalized": "2016-08-14 10:45:37",
	"checks": [
		{
		"date_received": "2016-08-14 10:45:37",
		"code": "1111",
		"status": "VALID",
		"ip_address": ""
		}
	],
	"first_event_date": "2016-08-14 10:45:26",
	"last_event_date": "2016-08-14 10:45:26",
	"price": "0.10000000",
	"currency": "EUR",
	"status": "SUCCESS"
	}`)

	response, err := parseVerifySearchResponse(bytes)
	if err != nil {
		t.Errorf("err was non-nil! %s", err)
		return
	}
	assertString(t, response.RequestID, "ad883d3ba753473694d9d6c70f529124")
	assertString(t, response.Checks[0].Code, "1111")
}

func TestParseVerifySearchResponseFailure(t *testing.T) {
	bytes := []byte(`{"status":"101","error_text":"No response found"}`)

	response, err := parseVerifySearchResponse(bytes)
	if response != nil {
		t.Errorf("response was non-nil! %s", response)
	}

	if err == nil {
		t.Errorf("Error should have been non-nil!")
	}

	assertString(t, err.Error(), "101: No response found")
}

func assertString(t *testing.T, value, expected string) {
	if value != expected {
		t.Errorf("Expected %s to be %s", value, expected)
	}
}

func TestParseVerifySearchResponseMultiple(t *testing.T) {
	bytes := []byte(`{
	"verification_requests": [
		{
		"request_id": "ad883d3ba753473694d9d6c70f529124",
		"account_id": "accountID",
		"number": "447720444444",
		"sender_id": "verify",
		"date_submitted": "2016-08-14 10:45:26",
		"date_finalized": "2016-08-14 10:45:37",
		"checks": [
			{
			"date_received": "2016-08-14 10:45:37",
			"code": "1111",
			"status": "VALID",
			"ip_address": ""
			}
		],
		"first_event_date": "2016-08-14 10:45:26",
		"last_event_date": "2016-08-14 10:45:26",
		"price": "0.10000000",
		"currency": "EUR",
		"status": "SUCCESS"
		}
	]
	}`)

	response, err := parseVerifySearchResponseMultiple(bytes)
	if err != nil {
		t.Errorf("err was non-nil! %s", err)
		return
	}

	if len(response) != 1 {
		t.Errorf("Length of response expected to be 1. Instead was %d", len(response))
		return
	}
	assertString(t, response[0].RequestID, "ad883d3ba753473694d9d6c70f529124")
	assertString(t, response[0].Checks[0].Code, "1111")
}
