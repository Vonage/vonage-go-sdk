package nexmo

import (
	"net/http"

	"fmt"

	"github.com/dghubble/sling"
)

type InsightService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

type BasicInsightRequest struct {
	Number  string
	Country string
	Cnam    *bool
}

type basicInsightJSONRequest struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
	Number    string `json:"number,omitempty"`
	Country   string `json:"country,omitempty"`
}

type BasicInsightResponse struct {
	Status                    int64  `json:"status"`
	StatusMessage             string `json:"status_message,omitempty"`
	ErrorText                 string `json:"error_text,omitempty"`
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

func newInsightService(base *sling.Sling, authSet *AuthSet) *InsightService {
	sling := base.Base("https://api.nexmo.com/ni/")
	return &InsightService{
		sling:   sling,
		authSet: authSet,
	}
}

func (c *InsightService) GetBasicInsight(request BasicInsightRequest) (BasicInsightResponse, *http.Response, error) {
	if c.authSet.apiSecret == nil {
		return BasicInsightResponse{}, nil, fmt.Errorf("Cannot call GetBasicInsight without providing APISecretAuth")
	}
	jsonRequest := basicInsightJSONRequest{
		APIKey:    c.authSet.apiSecret.apiKey,
		APISecret: c.authSet.apiSecret.apiSecret,
		Number:    request.Number,
		Country:   request.Country,
	}

	insightResponse := new(BasicInsightResponse)
	resp, err := c.sling.New().Post("basic/json").BodyJSON(jsonRequest).ReceiveSuccess(insightResponse)
	if err != nil {
		return *insightResponse, resp, err
	}
	err = insightResponse.responseError()
	return *insightResponse, resp, err
}

func (c *InsightService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
