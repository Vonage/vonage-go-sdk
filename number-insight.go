package vonage

import (
	"context"
	"runtime"

	"github.com/antihax/optional"
	"github.com/vonage/vonage-go-sdk/numberInsight"
)

// NumberInsightClient for working with the NumberInsight API
type NumberInsightClient struct {
	Config    *numberInsight.Configuration
	apiKey    string
	apiSecret string
}

// NewNumberInsightClient Creates a new NumberInsight Client, supplying an Auth to work with
func NewNumberInsightClient(Auth Auth) *NumberInsightClient {
	client := new(NumberInsightClient)
	creds := Auth.GetCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]

	client.Config = numberInsight.NewConfiguration()
	client.Config.UserAgent = "vonage-go/0.15-dev Go/" + runtime.Version()
	return client
}

type NiErrorResponse struct {
	Status        int32
	StatusMessage string
}

type NiOpts struct {
	Country string
}

type NiResponseJsonBasic struct {
	Status                    numberInsight.NiBasicStatus
	StatusMessage             string
	RequestId                 string
	InternationalFormatNumber string
	NationalFormatNumber      string
	CountryCode               string
	CountryCodeIso3           string
	CountryName               string
	CountryPrefix             string
}

// Basic does a basic-level lookup for data about a number
func (client *NumberInsightClient) Basic(number string, opts NiOpts) (NiResponseJsonBasic, NiErrorResponse, error) {
	// create the client
	numberInsightClient := numberInsight.NewAPIClient(client.Config)

	niOpts := numberInsight.GetNumberInsightBasicOpts{}

	if opts.Country != "" {
		niOpts.Country = optional.NewString(opts.Country)
	}

	// we need context for the API key
	ctx := context.Background()
	ctx = context.WithValue(ctx, numberInsight.ContextAPIKey, numberInsight.APIKey{Key: client.apiKey})
	ctx = context.WithValue(ctx, numberInsight.ContextAPISecret, numberInsight.APIKey{Key: client.apiSecret})

	result, _, err := numberInsightClient.DefaultApi.GetNumberInsightBasic(ctx, "json", number, &niOpts)

	// catch HTTP errors
	if err != nil {
		return NiResponseJsonBasic{}, NiErrorResponse{}, err
	}

	if result.Status != 0 {
		errResp := NiErrorResponse{
			Status:        int32(result.Status),
			StatusMessage: result.StatusMessage,
		}
		return NiResponseJsonBasic(result), errResp, nil
	}

	return NiResponseJsonBasic(result), NiErrorResponse{}, nil
}

type NiResponseJsonStandard struct {
	Status                    numberInsight.NiBasicStatus
	StatusMessage             string
	RequestId                 string
	InternationalFormatNumber string
	NationalFormatNumber      string
	CountryCode               string
	CountryCodeIso3           string
	CountryName               string
	CountryPrefix             string
	RequestPrice              string
	RefundPrice               string
	RemainingBalance          string
	CurrentCarrier            numberInsight.NiCurrentCarrierProperties
	OriginalCarrier           numberInsight.NiInitialCarrierProperties
	Ported                    string
	Roaming                   numberInsight.NiRoaming
	CallerIdentity            numberInsight.NiCallerIdentity
	CallerName                string
	LastName                  string
	FirstName                 string
	CallerType                string
}

// Standard does a Standard-level lookup for data about a number
func (client *NumberInsightClient) Standard(number string, opts NiOpts) (NiResponseJsonStandard, NiErrorResponse, error) {
	// create the client
	numberInsightClient := numberInsight.NewAPIClient(client.Config)

	niOpts := numberInsight.GetNumberInsightStandardOpts{}

	// we need context for the API key
	ctx := context.Background()
	ctx = context.WithValue(ctx, numberInsight.ContextAPIKey, numberInsight.APIKey{Key: client.apiKey})
	ctx = context.WithValue(ctx, numberInsight.ContextAPISecret, numberInsight.APIKey{Key: client.apiSecret})

	result, _, err := numberInsightClient.DefaultApi.GetNumberInsightStandard(ctx, "json", number, &niOpts)

	// catch HTTP errors
	if err != nil {
		return NiResponseJsonStandard{}, NiErrorResponse{}, err
	}

	if result.Status != 0 {
		errResp := NiErrorResponse{
			Status:        int32(result.Status),
			StatusMessage: result.StatusMessage,
		}
		return NiResponseJsonStandard(result), errResp, nil
	}

	return NiResponseJsonStandard(result), NiErrorResponse{}, nil
}

type NiResponseAsync struct {
	RequestId        string
	Number           string
	RemainingBalance string
	RequestPrice     string
	Status           numberInsight.NiStandardAdvancedStatus
	StatusMessage    string
}

// AdvancedAsync requests a callback with advanced-level information about a number
func (client *NumberInsightClient) AdvancedAsync(number string, callback string, opts NiOpts) (NiResponseAsync, NiErrorResponse, error) {
	// create the client
	numberInsightClient := numberInsight.NewAPIClient(client.Config)

	niOpts := numberInsight.GetNumberInsightAsyncOpts{}

	// we need context for the API key
	ctx := context.Background()
	ctx = context.WithValue(ctx, numberInsight.ContextAPIKey, numberInsight.APIKey{Key: client.apiKey})
	ctx = context.WithValue(ctx, numberInsight.ContextAPISecret, numberInsight.APIKey{Key: client.apiSecret})

	result, _, err := numberInsightClient.DefaultApi.GetNumberInsightAsync(ctx, "json", callback, number, &niOpts)

	// catch HTTP errors
	if err != nil {
		return NiResponseAsync{}, NiErrorResponse{}, err
	}

	if result.Status != 0 {
		errResp := NiErrorResponse{
			Status:        int32(result.Status),
			StatusMessage: result.StatusMessage,
		}
		return NiResponseAsync(result), errResp, nil
	}

	return NiResponseAsync(result), NiErrorResponse{}, nil
}
