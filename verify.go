package nexmo

import (
	"fmt"
	"net/http"
)

type StartVerificationRequest struct {
	Credentials
	Number        string `json:"number"`
	Brand         string `json:"brand"`
	Country       string `json:"country,omitempty"`
	SenderID      string `json:"sender_id,omitempty"`
	CodeLength    int8   `json:"code_length,omitempty"`
	LG            string `json:"lg,omitempty"`
	RequireType   string `json:"require_type,omitempty"`
	PINExpiry     int16  `json:"pin_expiry,omitempty"`
	NextEventWait int16  `json:"next_event_wait,omitempty"`
}

type StartVerificationResponse struct {
	RequestID string `json:"request_id"`
	Status    string `json:"status"`
	ErrorText string `json:"error_text"`
}

// Begin the process of verifying a phone number, you probably want to capture the request_id
func (s *VerifyService) Start(request StartVerificationRequest) (*StartVerificationResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(StartVerificationResponse)
	httpResponse, err := s.sling.New().
		Post("json").
		BodyJSON(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

type CheckVerificationRequest struct {
	Credentials
	RequestID string `json:"request_id"`
	Code      string `json:"code"`
	IPAddress string `json:"ip_address,omitempty"`
}

type CheckVerificationResponse struct {
	RequestID string `json:"event_id"`
	Status    string `json:"status"`
	Price     string `json:"price"`
	Currency  string `json:"currency"`
	ErrorText string `json:"error_text"`
}

// Check if the code the user supplied is correct for this request
func (s *VerifyService) Check(request CheckVerificationRequest) (*CheckVerificationResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(CheckVerificationResponse)
	httpResponse, err := s.sling.New().
		Post("check/json").
		BodyJSON(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

type SearchVerificationRequest struct {
	Credentials
	RequestIDs []string `json:"request_ids" url:"request_ids"`
}

type SearchVerificationResponse struct {
	Status               string `json:"status"`
	ErrorText            string `json:"error_text"`
	VerificationRequests []struct {
		RequestID      string `json:"request_id"`
		AccountID      string `json:"account_id"`
		Number         string `json:"number"`
		SenderID       string `json:"sender_id"`
		DateSubmitted  string `json:"date_submitted"`
		DateFinalized  string `json:"date_finalized"`
		FirstEventDate string `json:"first_event_date"`
		LastEventDate  string `json:"last_event_date"`
		Status         string `json:"status"`
		Price          string `json:"price"`
		Currency       string `json:"currency"`
		Checks         []struct {
			DateReceived string `json:"date_received"`
			Code         string `json:"code"`
			Status       string `json:"status"`
			IPAddress    string `json:"ip_address"`
		} `json:"checks"`
	} `json:"verification_requests"`
}

// Search for current or past verify requests, their costs and statuses
func (s *VerifyService) Search(request SearchVerificationRequest) (*SearchVerificationResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(SearchVerificationResponse)
	httpResponse, err := s.sling.New().
		Get("search/json").
		QueryStruct(request).
		ReceiveSuccess(response)
	if response.Status != "" {
		err = fmt.Errorf("%s: %s", response.Status, response.ErrorText)
	}
	return response, httpResponse, err
}

type ControlVerificationRequest struct {
	Credentials
	RequestID string `json:"request_id"`
	Command   string `json:"cmd"`
}

type ControlVerificationResponse struct {
	Status    string `json:"status"`
	Command   string `json:"command"`
	ErrorText string `json:"error_text"`
}

// The control endpoint allows cancellation of a request or moving to the next verification stage
func (s *VerifyService) Control(request ControlVerificationRequest) (*ControlVerificationResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(ControlVerificationResponse)
	httpResponse, err := s.sling.New().
		Post("control/json").
		BodyJSON(request).
		ReceiveSuccess(response)
	if response.Status != "" {
		err = fmt.Errorf("%s: %s", response.Status, response.ErrorText)
	}
	return response, httpResponse, err
}
