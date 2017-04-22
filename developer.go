package nexmo

import (
	"fmt"
	"net/http"
)

func (s *DeveloperService) GetBalance() (*GetBalanceResponse, *http.Response, error) {
	request := new(Credentials)
	s.authSet.ApplyAPICredentials(request)
	sling := s.sling.New().Get("account/get-balance").QueryStruct(request)

	response := new(GetBalanceResponse)
	httpResponse, err := sling.ReceiveSuccess(response)
	return response, httpResponse, err
}

type GetOutboundPricingForCountryRequest struct {
	Credentials
	Country string `url:"country"`
}

// GetOutboundPricingForCountry requests pricing for a given country
func (s *DeveloperService) GetOutboundPricingForCountry(request GetOutboundPricingForCountryRequest) (*CountryPrices, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(CountryPrices)
	httpResponse, err := s.sling.New().
		Get("account/get-pricing/outbound/").
		QueryStruct(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

type GetOutboundProductPricingRequest struct {
	Credentials
	Product string `url:"-"`
	Country string `url:"country"`
}

type GetOutboundProductPricingResponse struct {
	DialingPrefix      string           `json:"dialingPrefix"`
	DefaultPrice       string           `json:"defaultPrice"`
	Currency           string           `json:"currency"`
	CountryDisplayName string           `json:"countryDisplayName"`
	CountryCode        string           `json:"countryCode"`
	CountryName        string           `json:"countryName"`
	Networks           []NetworkDetails `json:"networks"`
}

type NetworkDetails struct {
	Type        string  `json:"type"`
	Price       string  `json:"price"`
	Currency    string  `json:"currency"`
	Ranges      []int64 `json:"ranges"`
	MNC         string  `json:"mnc"`
	MCC         string  `json:"mcc"`
	NetworkCode string  `json:"networkCode"`
	NetworkName string  `json:"networkName"`
}

// GetOutboundProductPricing requests prices for a product in a given country
func (s *DeveloperService) GetOutboundProductPricing(request GetOutboundProductPricingRequest) (*GetOutboundProductPricingResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(GetOutboundProductPricingResponse)
	httpResponse, err := s.sling.New().
		Get(fmt.Sprintf("account/get-pricing/outbound/%s/", request.Product)).
		QueryStruct(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

type GetPrefixOutboundPricingRequest struct {
	Credentials
	Prefix string `url:"prefix"`
}

type GetPrefixOutboundPricingResponse struct {
	Count  int64           `json:"count"`
	Prices []CountryPrices `json:"prices"`
}

type CountryPrices struct {
	MT       string         `json:"mt"`
	Country  string         `json:"country"`
	Prefix   string         `json:"prefix"`
	Name     string         `json:"name"`
	Networks []NetworkPrice `json:"networks"`
}

type NetworkPrice struct {
	Network string `json:"network"`
	Code    string `json:"code"`
	MTPrice string `json:"mtPrice"`
}

// GetPrefixOutboundPricing requests outbound pricing for a given international prefix
func (s *DeveloperService) GetPrefixOutboundPricing(request GetPrefixOutboundPricingRequest) (*GetPrefixOutboundPricingResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(GetPrefixOutboundPricingResponse)
	httpResponse, err := s.sling.New().
		Get("account/get-prefix-pricing/outbound/").
		QueryStruct(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

type GetPhoneOutboundPricingRequest struct {
	// This is defined here because Product must not be serialized, and I haven't implemented that in the json-schema
	Credentials
	Phone   string `url:"phone"`
	Product string `url:"-"`
}

// GetPhoneOutboundPricing requests outbound pricing for a given phone number
func (s *DeveloperService) GetPhoneOutboundPricing(request GetPhoneOutboundPricingRequest) (*GetPhoneOutboundPricingResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(GetPhoneOutboundPricingResponse)
	httpResponse, err := s.sling.New().
		Get(fmt.Sprintf("account/get-phone-pricing/outbound/%s", request.Product)).
		QueryStruct(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}
