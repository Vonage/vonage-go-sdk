package nexmo

import (
	"context"
	"runtime"

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

func (client *NumbersClient) List() (numbers.InboundNumbers, error) {

	numbersClient := numbers.NewAPIClient(client.Config)

	numbersOpts := numbers.GetOwnedNumbersOpts{}

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
