package nexmo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// PathVerify is the endpoint path for submitting verification requests
const pathVerify = "/verify/json"

// PathVerifyCheck is the endpoint path for submitting verification check requests
const pathVerifyCheck = "/verify/check/json"

// PathVerifySearch is the endpoint path for submitting verification search requests
const pathVerifySearch = "/verify/search/json"

// PathVerifyControl is the endpoint path for submitting verification control requests
const pathVerifyControl = "/verify/control/json"

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

// verificationResponse holds the raw API struct returned by a verify call.
type VerifyResponse struct {
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
	response := VerifyResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// CheckResponse holds the values returned from a successful check request.
type CheckResponse struct {
	EventID   string `json:"event_id"`
	Status    string `json:"status"`
	Price     string `json:"price"`
	Currency  string `json:"currency"`
	ErrorText string `json:"error_text"`
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
	response := CheckResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response, err
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

//----------------------------------------------------------------------------

type VerifySearchResponse struct {
	RequestID      string        `json:"request_id"`
	AccountID      string        `json:"account_id"`
	Number         string        `json:"number"`
	SenderID       string        `json:"sender_id"`
	DateSubmitted  string        `json:"date_submitted"`
	DateFinalized  string        `json:"date_finalized"`
	FirstEventDate string        `json:"first_event_date"`
	LastEventDate  string        `json:"last_event_date"`
	Status         string        `json:"status"`
	Price          string        `json:"price"`
	Currency       string        `json:"currency"`
	ErrorText      string        `json:"error_text"`
	Checks         []VerifyCheck `json:"checks"`
}

type verifySearchResponseMultiple struct {
	VerificationRequests []VerifySearchResponse `json:"verification_requests"`
}

type VerifyCheck struct {
	DateReceived string `json:"date_received"`
	Code         string `json:"code"`
	Status       string `json:"status"`
	IPAddress    string `json:"ip_address"`
}

func parseVerifySearchResponse(data []byte) (*VerifySearchResponse, error) {
	response := VerifySearchResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "SUCCESS" {
		return nil, fmt.Errorf("%s: %s", response.Status, response.ErrorText)
	}

	return &response, err
}

func parseVerifySearchResponseMultiple(data []byte) ([]VerifySearchResponse, error) {
	response := verifySearchResponseMultiple{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.VerificationRequests, err
}

func buildVerifySearchURL(client *nexmoClient, requestID []string) (string, error) {
	params := url.Values{}
	params.Set("api_key", client.apiKey)
	params.Set("api_secret", client.apiSecret)

	switch len(requestID) {
	case 0:
		return "", fmt.Errorf("At least one requestID must be provided.")
	case 1:
		params.Set("request_id", requestID[0])
	default:
		for _, id := range requestID {
			params.Add("request_ids", id)
		}
	}

	url, err := url.Parse(client.baseURL + pathVerifyCheck)
	if err != nil {
		return "", err
	}
	url.RawQuery = params.Encode()

	return url.String(), nil
}

func (client nexmoClient) VerifySearch(requestID string) (*VerifySearchResponse, error) {
	url, err := buildVerifySearchURL(&client, []string{requestID})
	if err != nil {
		return nil, err
	}

	bytes, err := urlSlurp(url)
	if err != nil {
		return nil, err
	}

	return parseVerifySearchResponse(bytes)
}

func (client nexmoClient) VerifySearchMultiple(requestID ...string) ([]VerifySearchResponse, error) {
	url, err := buildVerifySearchURL(&client, requestID)
	if err != nil {
		return nil, err
	}

	bytes, err := urlSlurp(url)
	if err != nil {
		return nil, err
	}

	return parseVerifySearchResponseMultiple(bytes)
}
