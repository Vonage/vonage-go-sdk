package nexmo

import (
	"net/http"

	"fmt"
)

type BasicInsightRequest struct {
	Credentials
	Number  string `json:"number,omitempty"`
	Country string `json:"country,omitempty"`
}

type BasicInsightResponse struct {
	Status                    int64  `json:"status,omitempty"`
	StatusMessage             string `json:"status_message,omitempty"`
	ErrorText                 string `json:"error_text,omitempty"`
	RequestID                 string `json:"request_id,omitempty"`
	InternationalFormatNumber string `json:"international_format_number,omitempty"`
	NationalFormatNumber      string `json:"national_format_number,omitempty"`
	CountryCode               string `json:"country_code,omitempty"`
	CountryCodeIso3           string `json:"country_code_iso3,omitempty"`
	CountryName               string `json:"country_name,omitempty"`
	CountryPrefix             string `json:"country_prefix,omitempty"`
}

// Format an error if we got one
func (r *BasicInsightResponse) responseError() error {
	if r.Status != 0 && r.Status != 43 && r.Status != 44 && r.Status != 45 {
		if r.StatusMessage != "" {
			return fmt.Errorf("%d: %s", r.Status, r.StatusMessage)
		}
		return fmt.Errorf("%d: %s", r.Status, r.ErrorText)
	}
	return nil
}

// Peform a Number Insight request at a basic level of detail
func (c *InsightService) GetBasicInsight(request BasicInsightRequest) (BasicInsightResponse, *http.Response, error) {
	c.authSet.ApplyAPICredentials(&request)

	insightResponse := new(BasicInsightResponse)
	resp, err := c.sling.New().Post("basic/json").BodyJSON(request).ReceiveSuccess(insightResponse)
	if err != nil {
		return *insightResponse, resp, err
	}
	return *insightResponse, resp, insightResponse.responseError()
}

type StandardInsightRequest struct {
	Credentials
	Number  string `json:"number,omitempty"`
	Country string `json:"country,omitempty"`
	CNAM    bool   `json:"cnam,omitempty"`
}

type StandardInsightResponse struct {
	BasicInsightResponse
	RequestPrice     string         `json:"request_price"`
	RefundPrice      string         `json:"refund_price"`
	RemainingBalance string         `json:"remaining_balance"`
	Ported           string         `json:"ported"`
	CurrentCarrier   *CarrierRecord `json:"current_carrier,omitempty"`
	OriginalCarrier  *CarrierRecord `json:"original_carrier,omitempty"`

	// CNAM fields:
	CallerName string `json:"caller_name"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	CallerType string `json:"caller_type"`
}

type CarrierRecord struct {
	NetworkCode string `json:"network_code"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	NetworkType string `json:"network_type"`
}

// Peform a Number Insight request at a standard level of detail
func (c *InsightService) GetStandardInsight(request StandardInsightRequest) (StandardInsightResponse, *http.Response, error) {
	c.authSet.ApplyAPICredentials(&request)

	insightResponse := new(StandardInsightResponse)
	resp, err := c.sling.New().Post("standard/json").BodyJSON(request).ReceiveSuccess(insightResponse)
	if err != nil {
		return *insightResponse, resp, err
	}
	return *insightResponse, resp, insightResponse.responseError()
}

type AdvancedInsightRequest struct {
	Credentials
	Number  string `json:"number,omitempty"`
	Country string `json:"country,omitempty"`
	CNAM    bool   `json:"cnam,omitempty"`
	IP      string `json:"ip,omitempty"`
}

type AdvancedInsightResponse struct {
	StandardInsightResponse
	ValidNumber string `json:"valid_number"`
	Reachable   string `json:"reachable"`
	Ported      string `json:"ported"`
	Roaming     struct {
		Status             string `json:"status"`
		RoamingCountryCode string `json:"roaming_country_code"`
		RoamingNetworkCode string `json:"roaming_network_code"`
		RoamingNetworkName string `json:"roaming_network_name"`
	} `json:"roaming"`
	LookupOutcome        int    `json:"lookup_outcome"`
	LookupOutcomeMessage string `json:"lookup_outcome_message"`
	IP                   string `json:"ip"`
	IPWarnings           string `json:"ip_warnings"`
	IPMatchLevel         string `json:"ip_match_level"`
	IPCountry            string `json:"ip_country"`
}

// Peform a Number Insight request at an advanced level of detail
func (c *InsightService) GetAdvancedInsight(request AdvancedInsightRequest) (AdvancedInsightResponse, *http.Response, error) {
	c.authSet.ApplyAPICredentials(&request)

	insightResponse := new(AdvancedInsightResponse)
	resp, err := c.sling.New().Post("advanced/json").BodyJSON(request).ReceiveSuccess(insightResponse)
	if err != nil {
		return *insightResponse, resp, err
	}
	return *insightResponse, resp, insightResponse.responseError()
}
