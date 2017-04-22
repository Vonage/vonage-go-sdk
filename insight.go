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

func (r *BasicInsightResponse) responseError() error {
	if r.Status != 0 {
		if r.StatusMessage != "" {
			return fmt.Errorf("%d: %s", r.Status, r.StatusMessage)
		}
		return fmt.Errorf("%d: %s", r.Status, r.ErrorText)
	}
	return nil
}

func (c *InsightService) GetBasicInsight(request BasicInsightRequest) (BasicInsightResponse, *http.Response, error) {
	c.authSet.ApplyAPICredentials(&request)

	insightResponse := new(BasicInsightResponse)
	resp, err := c.sling.New().Post("basic/json").BodyJSON(request).ReceiveSuccess(insightResponse)
	if err != nil {
		return *insightResponse, resp, err
	}
	return *insightResponse, resp, insightResponse.responseError()
}
