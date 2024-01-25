package vonage

import (
	"context"
	"encoding/json"

	"github.com/antihax/optional"
	"github.com/vonage/vonage-go-sdk/internal/number"
)

// NumbersClient for working with the Numbers API
type NumbersClient struct {
	Config    *number.Configuration
	apiKey    string
	apiSecret string
}

// NewNumbersClient Creates a new Numbers Client, supplying an Auth to work with
func NewNumbersClient(Auth Auth) *NumbersClient {
	client := new(NumbersClient)
	creds := Auth.GetCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]

	// Use a default set of config but make it accessible
	client.Config = number.NewConfiguration()
	client.Config.UserAgent = GetUserAgent()
	transport := &APITransport{APISecret: client.apiSecret}
	client.Config.HTTPClient = transport.Client()
	return client
}

// NumbersResponse is the response format for the Numbers API
type NumbersResponse struct {
	ErrorCode      string `json:"error-code,omitempty"`
	ErrorCodeLabel string `json:"error-code-label,omitempty"`
}

// NumbersErrorResponse is the error format for the Numbers API
type NumbersErrorResponse struct {
	ErrorCode      string `json:"error-code,omitempty"`
	ErrorCodeLabel string `json:"error-code-label,omitempty"`
}

// NumbersOpts sets the options to use in finding the numbers already in the user's account
type NumbersOpts struct {
	ApplicationID  string
	HasApplication string // string because it's tri-state, not boolean
	Country        string
	Pattern        string
	SearchPattern  int32
	Size           int32
	Index          int32
}

type NumberDetail struct {
	Country               string
	Msisdn                string
	MoHttpUrl             string
	Type                  string
	Features              []string
	MessagesCallbackType  string
	MessagesCallbackValue string
	VoiceCallbackType     string
	VoiceCallbackValue    string
}

type NumberCollection struct {
	Count   int32
	Numbers []NumberDetail
}

// List shows the numbers you already own, filters and pagination are available
func (client *NumbersClient) List(opts NumbersOpts) (NumberCollection, NumbersErrorResponse, error) {
	return client.ListContext(context.Background(), opts)
}

// ListContext shows the numbers you already own, filters and pagination are available
func (client *NumbersClient) ListContext(ctx context.Context, opts NumbersOpts) (NumberCollection, NumbersErrorResponse, error) {

	numbersClient := number.NewAPIClient(client.Config)

	// set up the options and parse them
	numbersOpts := number.GetOwnedNumbersOpts{}

	if opts.ApplicationID != "" {
		numbersOpts.ApplicationId = optional.NewString(opts.ApplicationID)
	}

	if opts.HasApplication != "" {
		// if it's set at all, use the value and set it
		if opts.HasApplication == "true" {
			numbersOpts.HasApplication = optional.NewBool(true)
		} else if opts.HasApplication == "false" {
			numbersOpts.HasApplication = optional.NewBool(false)
		}
	}

	if opts.Country != "" {
		numbersOpts.Country = optional.NewString(opts.Country)
	}

	if opts.Pattern != "" {
		numbersOpts.Pattern = optional.NewString(opts.Pattern)
	}

	if opts.SearchPattern != 0 {
		numbersOpts.SearchPattern = optional.NewInt32(opts.SearchPattern)
	}

	if opts.Size != 0 {
		numbersOpts.Size = optional.NewInt32(opts.Size)
	}

	if opts.Index != 0 {
		numbersOpts.Index = optional.NewInt32(opts.Index)
	}

	// we need context for the API key
	ctx = context.WithValue(ctx, number.ContextAPIKey, number.APIKey{
		Key: client.apiKey,
	})

	result, _, err := numbersClient.DefaultApi.GetOwnedNumbers(ctx, &numbersOpts)

	if err != nil {
		e := err.(number.GenericOpenAPIError)
		data := e.Body()

		var errResp NumbersErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			return NumberCollection{}, errResp, err
		}
		return NumberCollection{}, NumbersErrorResponse{}, err
	}

	// deep-convert the numbers collection
	var collection NumberCollection
	var numbers []NumberDetail
	for _, num := range result.Numbers {
		numbers = append(numbers, NumberDetail(num))
	}
	collection.Numbers = numbers
	collection.Count = result.Count

	return collection, NumbersErrorResponse{}, nil

}

// NumberSearchOpts sets the optional values in the Search method
type NumberSearchOpts struct {
	Type          string
	Features      string
	Pattern       string
	SearchPattern int32
	Size          int32
	Index         int32
}

type NumberSearch struct {
	Count   int32
	Numbers []NumberAvailable
}

type NumberAvailable struct {
	Country  string
	Msisdn   string
	Type     string
	Cost     string
	Features []string
}

// Search lets you find a great phone number to use in your application
func (client *NumbersClient) Search(country string, opts NumberSearchOpts) (NumberSearch, NumbersErrorResponse, error) {
	return client.SearchContext(context.Background(), country, opts)
}

// SearchContext lets you find a great phone number to use in your application
func (client *NumbersClient) SearchContext(ctx context.Context, country string, opts NumberSearchOpts) (NumberSearch, NumbersErrorResponse, error) {

	numbersClient := number.NewAPIClient(client.Config)

	// we need context for the API key
	ctx = context.WithValue(ctx, number.ContextAPIKey, number.APIKey{
		Key: client.apiKey,
	})

	numbersSearchOpts := number.GetAvailableNumbersOpts{}

	if opts.Type != "" {
		numbersSearchOpts.Type_ = optional.NewString(opts.Type)
	}

	if opts.Features != "" {
		numbersSearchOpts.Features = optional.NewString(opts.Features)
	}

	if opts.Pattern != "" {
		numbersSearchOpts.Pattern = optional.NewString(opts.Pattern)
	}

	if opts.SearchPattern != 0 {
		numbersSearchOpts.SearchPattern = optional.NewInt32(opts.SearchPattern)
	}

	if opts.Size != 0 {
		numbersSearchOpts.Size = optional.NewInt32(opts.Size)
	}

	if opts.Index != 0 {
		numbersSearchOpts.Index = optional.NewInt32(opts.Index)
	}

	result, _, err := numbersClient.DefaultApi.GetAvailableNumbers(ctx, country, &numbersSearchOpts)

	if err != nil {
		return NumberSearch{}, NumbersErrorResponse{}, err
	}

	// deep-convert the numbers collection
	var collection NumberSearch
	var numbers []NumberAvailable
	for _, num := range result.Numbers {
		numbers = append(numbers, NumberAvailable(num))
	}
	collection.Numbers = numbers
	collection.Count = result.Count

	return collection, NumbersErrorResponse{}, nil
}

// NumberBuyOpts enables users to set the Target API Key (and any future params)
type NumberBuyOpts struct {
	TargetAPIKey string
}

// Buy the best phone number to use in your app
func (client *NumbersClient) Buy(country string, msisdn string, opts NumberBuyOpts) (NumbersResponse, NumbersErrorResponse, error) {
	return client.BuyContext(context.Background(), country, msisdn, opts)
}

// BuyContext the best phone number to use in your app
func (client *NumbersClient) BuyContext(ctx context.Context, country string, msisdn string, opts NumberBuyOpts) (NumbersResponse, NumbersErrorResponse, error) {

	numbersClient := number.NewAPIClient(client.Config)

	// we need context for the API key
	ctx = context.WithValue(ctx, number.ContextAPIKey, number.APIKey{
		Key: client.apiKey,
	})

	numbersBuyOpts := number.BuyANumberOpts{}

	if opts.TargetAPIKey != "" {
		numbersBuyOpts.TargetApiKey = optional.NewString(opts.TargetAPIKey)
	}

	result, resp, err := numbersClient.DefaultApi.BuyANumber(ctx, country, msisdn, &numbersBuyOpts)
	// check for non-200 status codes first, err will be set but we handle these specifically
	if resp.StatusCode != 200 {
		// handle a 4xx error
		e := err.(number.GenericOpenAPIError)
		data := e.Body()

		var errResp NumbersErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			if errResp.ErrorCode == "420" && errResp.ErrorCodeLabel == "method failed" {
				errResp.ErrorCodeLabel = "method failed. This can also indicate that you already own this number"
			}
			return NumbersResponse(result), errResp, nil
		}
	}

	if err != nil {
		return NumbersResponse{}, NumbersErrorResponse{}, err
	}

	return NumbersResponse(result), NumbersErrorResponse{}, nil
}

// NumberCancelOpts enables users to set the Target API Key (and any future params)
type NumberCancelOpts struct {
	TargetAPIKey string
}

// Cancel a number already in your account
func (client *NumbersClient) Cancel(country string, msisdn string, opts NumberCancelOpts) (NumbersResponse, NumbersErrorResponse, error) {
	return client.CancelContext(context.Background(), country, msisdn, opts)
}

// CancelContext a number already in your account
func (client *NumbersClient) CancelContext(ctx context.Context, country string, msisdn string, opts NumberCancelOpts) (NumbersResponse, NumbersErrorResponse, error) {
	numbersClient := number.NewAPIClient(client.Config)

	// we need context for the API key
	ctx = context.WithValue(ctx, number.ContextAPIKey, number.APIKey{
		Key: client.apiKey,
	})

	numbersCancelOpts := number.CancelANumberOpts{}

	if opts.TargetAPIKey != "" {
		numbersCancelOpts.TargetApiKey = optional.NewString(opts.TargetAPIKey)
	}

	result, resp, err := numbersClient.DefaultApi.CancelANumber(ctx, country, msisdn, &numbersCancelOpts)
	if resp.StatusCode != 200 {
		// handle a 4xx error
		e := err.(number.GenericOpenAPIError)
		data := e.Body()

		var errResp NumbersErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			if errResp.ErrorCode == "420" && errResp.ErrorCodeLabel == "method failed" {
				// expand on this error code, it's commonly because you don't own the number
				errResp.ErrorCodeLabel = "method failed. This can also indicate that the number is not associated with this key"
			}
			return NumbersResponse(result), errResp, nil
		}
	}

	if err != nil {
		return NumbersResponse{}, NumbersErrorResponse{}, err
	}

	return NumbersResponse(result), NumbersErrorResponse{}, nil
}

// NumberUpdateOpts sets all the various fields for the number config
type NumberUpdateOpts struct {
	AppID                 string
	MoHTTPURL             string
	VoiceCallbackType     string
	VoiceCallbackValue    string
	VoiceStatusCallback   string
	MessagesCallbackType  string
	MessagesCallbackValue string
}

// Update the configuration for your number
func (client *NumbersClient) Update(country string, msisdn string, opts NumberUpdateOpts) (NumbersResponse, NumbersErrorResponse, error) {
	return client.UpdateContext(context.Background(), country, msisdn, opts)
}

// UpdateContext the configuration for your number
func (client *NumbersClient) UpdateContext(ctx context.Context, country string, msisdn string, opts NumberUpdateOpts) (NumbersResponse, NumbersErrorResponse, error) {
	numbersClient := number.NewAPIClient(client.Config)

	// we need context for the API key
	ctx = context.WithValue(ctx, number.ContextAPIKey, number.APIKey{
		Key: client.apiKey,
	})

	numbersUpdateOpts := number.UpdateANumberOpts{}

	if opts.AppID != "" {
		numbersUpdateOpts.AppId = optional.NewString(opts.AppID)
	}

	if opts.MoHTTPURL != "" {
		numbersUpdateOpts.MoHttpUrl = optional.NewString(opts.MoHTTPURL)
	}

	if opts.VoiceCallbackType != "" {
		numbersUpdateOpts.VoiceCallbackType = optional.NewString(opts.VoiceCallbackType)
	}

	if opts.VoiceCallbackValue != "" {
		numbersUpdateOpts.VoiceCallbackValue = optional.NewString(opts.VoiceCallbackValue)
	}

	if opts.VoiceStatusCallback != "" {
		numbersUpdateOpts.VoiceStatusCallback = optional.NewString(opts.VoiceStatusCallback)
	}

	result, resp, err := numbersClient.DefaultApi.UpdateANumber(ctx, country, msisdn, &numbersUpdateOpts)
	if err != nil {
		return NumbersResponse{}, NumbersErrorResponse{}, err
	}

	if resp.StatusCode != 200 {
		// handle a 4xx error
		e := err.(number.GenericOpenAPIError)
		data := e.Body()

		var errResp NumbersErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			return NumbersResponse(result), errResp, nil
		}
	}

	return NumbersResponse(result), NumbersErrorResponse{}, nil
}
