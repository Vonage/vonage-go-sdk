package vonage

import (
	"context"
	"runtime"

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
	Status        numberInsight.NiBasicStatus
	StatusMessage string
}

// Basic does a basic-level lookup for data about a number
func (client *NumberInsightClient) Basic(number string) (numberInsight.NiResponseJsonBasic, NiErrorResponse, error) {
	// create the client
	numberInsightClient := numberInsight.NewAPIClient(client.Config)

	niOpts := numberInsight.GetNumberInsightBasicOpts{}

	// we need context for the API key
	ctx := context.Background()
	ctx = context.WithValue(ctx, numberInsight.ContextAPIKey, numberInsight.APIKey{Key: client.apiKey})
	ctx = context.WithValue(ctx, numberInsight.ContextAPISecret, numberInsight.APIKey{Key: client.apiSecret})

	result, _, err := numberInsightClient.DefaultApi.GetNumberInsightBasic(ctx, "json", number, &niOpts)

	// catch HTTP errors
	if err != nil {
		return numberInsight.NiResponseJsonBasic{}, NiErrorResponse{}, err
	}

	if result.Status != 0 {
		errResp := NiErrorResponse{
			Status:        result.Status,
			StatusMessage: result.StatusMessage,
		}
		return result, errResp, nil
	}

	return result, NiErrorResponse{}, nil
}
