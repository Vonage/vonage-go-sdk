package nexmo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// VerifyRequest encapsulates all of the possible arguments for a verify call.
//
// Parameter values are documented at https://docs.nexmo.com/verify/api-reference/api-reference
type VerifyRequest struct {
	Number string
	Brand string
	Country string
	SenderID string
	CodeLength int
	Locale string
	RequireType string
	PinExpiry	time.Duration
	NextEventWait time.Duration
}

func (request VerifyRequest) toURLValues(params *url.Values) {
	params.Set("number", request.Number)
	params.Set("brand", request.Brand)

	setStringParam := func(key, value string) {
		if value != "" {
			params.Set(key, value)
		}
	}

	if request.CodeLength > 0 {
		params.Set("code_length", string(request.CodeLength))
	}

	setStringParam("country", request.Country)
	setStringParam("sender_id", request.SenderID)
	setStringParam("locale", request.Locale)
	setStringParam("require_type", request.RequireType)
	if request.PinExpiry > 0 {
		params.Set("pin_expiry", string(int(request.PinExpiry.Seconds())))
	}
	if request.NextEventWait > 0 {
		params.Set("next_event_wait", string(int(request.NextEventWait.Seconds())))
	}
}

// NewVerifyRequest creates a new VerifyRequest with compulsory values.
func NewVerifyRequest(number, brand string) VerifyRequest {
	return VerifyRequest{
		Number: number,
		Brand: brand,
	}
}

// VerifyResponse holds the values returned by a successful verify call.
type VerifyResponse struct {
	RequestID string
}

// verificationResponse holds the raw API struct returned by a verify call.
type verificationResponse struct {
	RequestID string `json:"request_id"`
	Status    string `json:"status"`
	ErrorText string `json:"error_text"`
}

func (response verificationResponse) StatusCode() (int, error) {
	return strconv.Atoi(response.Status)
}

func (client nexmoClient) Verify(request VerifyRequest) (*VerifyResponse, error) {
	// TODO: Needs more validation.
	// TODO: Timeouts are currently ignored.

	length := request.CodeLength
	if length > 0 && length != 4 && length != 6 {
		return nil, fmt.Errorf("code length must be 4 or 6")
	}

	params := url.Values{}
	params.Set("api_key", client.apiKey)
	params.Set("api_secret", client.apiSecret)
	request.toURLValues(&params)

	url, err := url.Parse(client.baseURL + pathVerify)
	if err != nil {
		return nil, err
	}
	url.RawQuery = params.Encode()

	httpRequest, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected response from server: %d %s", response.StatusCode, response.Status)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return parseVerifyResponse(bytes)
}

func parseVerifyResponse(data []byte) (*VerifyResponse, error) {
	response := verificationResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	statusCode, err := response.StatusCode()
	if err != nil {
		return nil, err
	}
	if statusCode != 0 {
		return nil, fmt.Errorf("%s: %s", response.Status, response.ErrorText)
	}

	return &VerifyResponse{RequestID: response.RequestID}, err
}
