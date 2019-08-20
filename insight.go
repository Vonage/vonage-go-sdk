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

// Returns a formatted error if there is an error, or nil otherwise
func (r *BasicInsightResponse) responseError() error {
	if r.Status != 0 && r.Status != 43 && r.Status != 44 && r.Status != 45 {
		if r.StatusMessage != "" {
			return fmt.Errorf("%d: %s", r.Status, r.StatusMessage)
		}
		return fmt.Errorf("%d: %s", r.Status, r.ErrorText)
	}
	return nil
}

// Perform a Number Insight request at a basic level of detail
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

type PortedStatus string

type CallerType string

const (
	PortedStatusUnknown          PortedStatus = "unknown"
	PortedStatusPorted           PortedStatus = "ported"
	PortedStatusNotPorted        PortedStatus = "not_ported"
	PortedStatusAssumedPorted    PortedStatus = "assumed_ported"
	PortedStatusAssumedNotPorted PortedStatus = "assumed_not_ported"

	CallerTypeBusiness CallerType = "business"
	CallerTypeConsumer CallerType = "consumer"
	CallerTypeUnknown  CallerType = "unknown"
)

type StandardInsightResponse struct {
	BasicInsightResponse
	RequestPrice     string         `json:"request_price"`
	RefundPrice      string         `json:"refund_price"`
	RemainingBalance string         `json:"remaining_balance"`
	Ported           PortedStatus   `json:"ported"`
	CurrentCarrier   *CarrierRecord `json:"current_carrier,omitempty"`
	OriginalCarrier  *CarrierRecord `json:"original_carrier,omitempty"`

	// CNAM fields:
	CallerName string     `json:"caller_name"`
	LastName   string     `json:"last_name"`
	FirstName  string     `json:"first_name"`
	CallerType CallerType `json:"caller_type"`
}

type NetworkType string

const (
	NetworkTypeMobile           NetworkType = "mobile"
	NetworkTypeLandline         NetworkType = "landline"
	NetworkTypeLandlinePremium  NetworkType = "landline_premium"
	NetworkTypeLandlineTollFree NetworkType = "landline_tollfree"
	NetworkTypeVirtual          NetworkType = "virtual"
	NetworkTypeUnknown          NetworkType = "unknown"
	NetworkTypePager            NetworkType = "pager"
)

type CarrierRecord struct {
	NetworkCode string      `json:"network_code"`
	Name        string      `json:"name"`
	Country     string      `json:"country"`
	NetworkType NetworkType `json:"network_type"`
}

// Perform a Number Insight request at a standard level of detail
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

type RoamingStatus string

type ValidNumberStatus string

type ReachableStatus string

type IPMatchLevel string

type IPWarnings string

type LookupOutcome int8

const (
	RoamingStatusRoaming    RoamingStatus = "roaming"
	RoamingStatusUnknown    RoamingStatus = "unkown"
	RoamingStatusNotRoaming RoamingStatus = "not_roaming"

	ValidNumberStatusValid    ValidNumberStatus = "valid"
	ValidNumberStatusUnknown  ValidNumberStatus = "unknown"
	ValidNumberStatusNotValid ValidNumberStatus = "not_valid"

	ReachableStatusUnknown       ReachableStatus = "unknown"
	ReachableStatusReachable     ReachableStatus = "reachable"
	ReachableStatusUndeliverable ReachableStatus = "undeliverable"
	ReachableStatusAbsent        ReachableStatus = "absent"
	ReachableStatusBadNumber     ReachableStatus = "bad_number"
	ReachableStatusBlacklisted   ReachableStatus = "blacklisted"

	IPMatchLevelCountry  IPMatchLevel = "country"
	IPMatchLevelMismatch IPMatchLevel = "mismatch"

	IPWarningsUnknown   IPWarnings = "unknown"
	IPWarningsNoWarning IPWarnings = "no_warning"

	LookupOutcomeSuccess LookupOutcome = iota
	LookupOutcomePartial
	LookupOutcomeFailed
)

type AdvancedInsightResponse struct {
	StandardInsightResponse
	ValidNumber ValidNumberStatus `json:"valid_number"`
	Reachable   ReachableStatus   `json:"reachable"`
	Ported      string            `json:"ported"`
	Roaming     struct {
		Status             RoamingStatus `json:"status"`
		RoamingCountryCode string        `json:"roaming_country_code"`
		RoamingNetworkCode string        `json:"roaming_network_code"`
		RoamingNetworkName string        `json:"roaming_network_name"`
	} `json:"roaming"`
	LookupOutcome        LookupOutcome `json:"lookup_outcome"`
	LookupOutcomeMessage string        `json:"lookup_outcome_message"`
	IP                   string        `json:"ip"`
	IPWarnings           IPWarnings    `json:"ip_warnings"`
	IPMatchLevel         IPMatchLevel  `json:"ip_match_level"`
	IPCountry            string        `json:"ip_country"`
}

// Perform a Number Insight request at an advanced level of detail
func (c *InsightService) GetAdvancedInsight(request AdvancedInsightRequest) (AdvancedInsightResponse, *http.Response, error) {
	c.authSet.ApplyAPICredentials(&request)

	insightResponse := new(AdvancedInsightResponse)
	resp, err := c.sling.New().Post("advanced/json").BodyJSON(request).ReceiveSuccess(insightResponse)
	if err != nil {
		return *insightResponse, resp, err
	}
	return *insightResponse, resp, insightResponse.responseError()
}
