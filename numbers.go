package nexmo

import (
	"context"
	"encoding/json"
	"runtime"

	"github.com/antihax/optional"
	"github.com/nexmo-community/nexmo-go/numbers"
)

// NumbersClient for working with the Numbers API
type NumbersClient struct {
	Config    *numbers.Configuration
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
	client.Config = numbers.NewConfiguration()
	client.Config.UserAgent = "nexmo-go/0.15-dev Go/" + runtime.Version()
	transport := &APITransport{APISecret: client.apiSecret}
	client.Config.HTTPClient = transport.Client()
	return client
}

// NumbersErrorResponse is the error format for the Numbers API
type NumbersErrorResponse struct {
	ErrorCode      string `json:"error-code,omitempty"`
	ErrorCodeLabel string `json:"error-code-label,omitempty"`
}

// NumbersOpts sets the options to use in finding the numbers already in the user's account
type NumbersOpts struct {
	ApplicationId  string
	HasApplication string // string because it's tri-state, not boolean
	Country        string
	Pattern        string
	SearchPattern  int32
	Size           int32
	Index          int32
}

// List shows the numbers you already own, filters and pagination are available
func (client *NumbersClient) List(opts NumbersOpts) (numbers.InboundNumbers, error) {

	numbersClient := numbers.NewAPIClient(client.Config)

	// set up the options and parse them
	numbersOpts := numbers.GetOwnedNumbersOpts{}

	if opts.ApplicationId != "" {
		numbersOpts.ApplicationId = optional.NewString(opts.ApplicationId)
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
	ctx := context.WithValue(context.Background(), numbers.ContextAPIKey, numbers.APIKey{
		Key: client.apiKey,
	})

	result, _, err := numbersClient.DefaultApi.GetOwnedNumbers(ctx, &numbersOpts)

	if err != nil {
		return numbers.InboundNumbers{}, err
	}

	return result, nil

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

func (client *NumbersClient) Search(country string, opts NumberSearchOpts) (numbers.AvailableNumbers, error) {

	numbersClient := numbers.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), numbers.ContextAPIKey, numbers.APIKey{
		Key: client.apiKey,
	})

	numbersSearchOpts := numbers.GetAvailableNumbersOpts{}

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
		return numbers.AvailableNumbers{}, err
	}

	return result, nil
}

// NumberBuyOpts enables users to set the Target API Key (and any future params)
type NumberBuyOpts struct {
	TargetApiKey string
}

func (client *NumbersClient) Buy(country string, msisdn string, opts NumberBuyOpts) (numbers.Response, NumbersErrorResponse, error) {

	numbersClient := numbers.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), numbers.ContextAPIKey, numbers.APIKey{
		Key: client.apiKey,
	})

	numbersBuyOpts := numbers.BuyANumberOpts{}

	if opts.TargetApiKey != "" {
		numbersBuyOpts.TargetApiKey = optional.NewString(opts.TargetApiKey)
	}

	result, resp, err := numbersClient.DefaultApi.BuyANumber(ctx, country, msisdn, &numbersBuyOpts)
	// check for non-200 status codes first, err will be set but we handle these specifically
	if resp.StatusCode != 200 {
		// handle a 4xx error
		e := err.(numbers.GenericOpenAPIError)
		data := e.Body()

		var errResp NumbersErrorResponse
		json.Unmarshal(data, &errResp)
		if errResp.ErrorCode == "420" && errResp.ErrorCodeLabel == "method failed" {
			errResp.ErrorCodeLabel = "method failed. This can also indicate that you already own this number"
		}
		return result, errResp, nil
	}

	if err != nil {
		return numbers.Response{}, NumbersErrorResponse{}, err
	}

	return result, NumbersErrorResponse{}, nil
}

// NumberCancelOpts enables users to set the Target API Key (and any future params)
type NumberCancelOpts struct {
	TargetApiKey string
}

func (client *NumbersClient) Cancel(country string, msisdn string, opts NumberCancelOpts) (numbers.Response, NumbersErrorResponse, error) {
	numbersClient := numbers.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), numbers.ContextAPIKey, numbers.APIKey{
		Key: client.apiKey,
	})

	numbersCancelOpts := numbers.CancelANumberOpts{}

	if opts.TargetApiKey != "" {
		numbersCancelOpts.TargetApiKey = optional.NewString(opts.TargetApiKey)
	}

	result, resp, err := numbersClient.DefaultApi.CancelANumber(ctx, country, msisdn, &numbersCancelOpts)
	if resp.StatusCode != 200 {
		// handle a 4xx error
		e := err.(numbers.GenericOpenAPIError)
		data := e.Body()

		var errResp NumbersErrorResponse
		json.Unmarshal(data, &errResp)
		if errResp.ErrorCode == "420" && errResp.ErrorCodeLabel == "method failed" {
			// expand on this error code, it's commonly because you don't own the number
			errResp.ErrorCodeLabel = "method failed. This can also indicate that the number is not associated with this key"
		}
		return result, errResp, nil
	}

	if err != nil {
		return numbers.Response{}, NumbersErrorResponse{}, err
	}

	return result, NumbersErrorResponse{}, nil
}

// NumberUpdateOpts sets all the various fields for the number config
type NumberUpdateOpts struct {
	AppId                 string
	MoHttpUrl             string
	VoiceCallbackType     string
	VoiceCallbackValue    string
	VoiceStatusCallback   string
	MessagesCallbackType  string
	MessagesCallbackValue string
}

func (client *NumbersClient) Update(country string, msisdn string, opts NumberUpdateOpts) (numbers.Response, NumbersErrorResponse, error) {
	numbersClient := numbers.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), numbers.ContextAPIKey, numbers.APIKey{
		Key: client.apiKey,
	})

	numbersUpdateOpts := numbers.UpdateANumberOpts{}

	if opts.AppId != "" {
		numbersUpdateOpts.AppId = optional.NewString(opts.AppId)
	}

	if opts.MoHttpUrl != "" {
		numbersUpdateOpts.MoHttpUrl = optional.NewString(opts.MoHttpUrl)
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
		return numbers.Response{}, NumbersErrorResponse{}, err
	}

	if resp.StatusCode != 200 {
		// handle a 4xx error
		e := err.(numbers.GenericOpenAPIError)
		data := e.Body()

		var errResp NumbersErrorResponse
		json.Unmarshal(data, &errResp)
		return result, errResp, nil
	}

	return result, NumbersErrorResponse{}, nil
}
