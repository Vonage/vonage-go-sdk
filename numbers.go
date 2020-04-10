package nexmo

import (
	"context"
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
