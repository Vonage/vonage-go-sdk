package nexmo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// VerifyRequest encapsulates all of the possible arguments for a verify call.
//
// Parameter values are documented at https://docs.nexmo.com/verify/api-reference/api-reference
type VerifyRequest struct {
	Number        string
	Brand         string
	Country       string
	SenderID      string
	CodeLength    int
	Locale        string
	RequireType   string
	PinExpiry     time.Duration
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
		Brand:  brand,
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

	bytes, err := urlSlurp(url.String())
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

	if response.Status != "0" {
		return nil, fmt.Errorf("%s: %s", response.Status, response.ErrorText)
	}

	return &VerifyResponse{RequestID: response.RequestID}, err
}

// CheckResponse holds the values returned from a successful check request.
type CheckResponse struct {
	EventID  string
	Price    string
	Currency string
}

type checkResponse struct {
	EventID   string `json:"event_id"`
	Status    string `json:"status"`
	Price     string `json:"price"`
	Currency  string `json:"currency"`
	ErrorText string `json:"error_text"`
}

func (response checkResponse) toPublic() *CheckResponse {
	return &CheckResponse{
		EventID:  response.EventID,
		Price:    response.Price,
		Currency: response.Currency,
	}
}

func buildCheckURL(client *nexmoClient, requestID, code string) (string, error) {
	params := url.Values{}
	params.Set("api_key", client.apiKey)
	params.Set("api_secret", client.apiSecret)
	params.Set("request_id", requestID)
	params.Set("code", code)

	url, err := url.Parse(client.baseURL + pathVerifyCheck)
	if err != nil {
		return "", err
	}
	url.RawQuery = params.Encode()

	return url.String(), nil
}

func parseCheckResponse(data []byte) (*CheckResponse, error) {
	response := checkResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "0" {
		return nil, fmt.Errorf("%s: %s", response.Status, response.ErrorText)
	}

	return response.toPublic(), err
}

func (client nexmoClient) Check(requestID, code string) (*CheckResponse, error) {
	url, err := buildCheckURL(&client, requestID, code)
	if err != nil {
		return nil, err
	}

	bytes, err := urlSlurp(url)
	if err != nil {
		return nil, err
	}

	return parseCheckResponse(bytes)
}

// Utility function for making a GET request and returning the body's bytes.
func urlSlurp(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected response from server: %d %s", response.StatusCode, response.Status)
	}

	return ioutil.ReadAll(response.Body)
}