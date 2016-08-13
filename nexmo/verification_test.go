package nexmo

import "testing"

func TestParseVerifyResponseSuccess(t *testing.T) {
	response, err := parseVerifyResponse([]byte(`{"request_id":"e80c552b22054bbc8b95e521520b0e1c","status":"0"}`))
	if err != nil {
		t.Errorf("err was not nil, was %s instead", err)
		return
	}
	if response == "" {
		t.Errorf("Response was empty!")
		return
	}
	if response != "e80c552b22054bbc8b95e521520b0e1c" {
		t.Errorf("Returned RequestID was incorrect: %s", response)
		return
	}
}

func TestParseVerifyResponseFailure(t *testing.T) {
	response, err := parseVerifyResponse([]byte(`{"status":"2","error_text":"Missing username"}`))
	if response != "" {
		t.Errorf("response was not empty! Instead it was %s", response)
		return
	}
	if err == nil {
		t.Errorf("err was nil!")
		return
	}
	if err.Error() != "2: Missing username" {
		t.Errorf("Error message was incorrect: \"%s\"", err)
		return
	}
}
