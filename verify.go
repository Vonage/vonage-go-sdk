package nexmo

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"runtime"

	"github.com/antihax/optional"
	"github.com/nexmo-community/nexmo-go/verify"
)

// VerifyClient for working with the Verify API
type VerifyClient struct {
	Config    *verify.Configuration
	apiKey    string
	apiSecret string
}

// NewVerifyClient Creates a new Verify Client, supplying an Auth to work with
func NewVerifyClient(Auth Auth) *VerifyClient {
	client := new(VerifyClient)
	creds := Auth.GetCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]

	client.Config = verify.NewConfiguration()
	client.Config.UserAgent = "nexmo-go/0.15-dev Go/" + runtime.Version()
	return client
}

// VerifyOpts holds all the optional arguments for the verify request function
type VerifyOpts struct {
	Country       string
	SenderId      string
	CodeLength    int32
	Lg            string
	PinExpiry     int32
	NextEventWait int32
	WorkflowId    int32
}

// Request a number is verified for ownership
func (client *VerifyClient) Request(number string, brand string, opts VerifyOpts) (verify.RequestResponse, verify.RequestErrorResponse, error) {
	// create the client
	verifyClient := verify.NewAPIClient(client.Config)

	// set up and then parse the options
	verifyOpts := verify.VerifyRequestOpts{}

	if opts.CodeLength != 0 {
		verifyOpts.CodeLength = optional.NewInt32(opts.CodeLength)
	}

	if opts.Lg != "" {
		verifyOpts.Lg = optional.NewString(opts.Lg)
	}

	if opts.WorkflowId != 0 {
		verifyOpts.WorkflowId = optional.NewInt32(opts.WorkflowId)
	}

	// we need context for the API key
	ctx := context.WithValue(context.Background(), verify.ContextAPIKey, verify.APIKey{
		Key: client.apiKey,
	})

	result, resp, err := verifyClient.DefaultApi.VerifyRequest(ctx, "json", client.apiSecret, number, brand, &verifyOpts)

	// catch HTTP errors
	if err != nil {
		return verify.RequestResponse{}, verify.RequestErrorResponse{}, err
	}

	// non-zero statuses are also errors
	if result.Status != "0" {
		data, _ := ioutil.ReadAll(resp.Body)

		var errResp verify.RequestErrorResponse
		json.Unmarshal(data, &errResp)
		return result, errResp, nil
	}
	return result, verify.RequestErrorResponse{}, nil
}

// Check the user-supplied code for this request ID
func (client *VerifyClient) Check(requestId string, code string) (verify.CheckResponse, verify.CheckErrorResponse, error) {
	// create the client
	verifyClient := verify.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), verify.ContextAPIKey, verify.APIKey{
		Key: client.apiKey,
	})

	// set up and then parse the options
	verifyOpts := verify.VerifyCheckOpts{}
	result, resp, err := verifyClient.DefaultApi.VerifyCheck(ctx, "json", client.apiSecret, requestId, code, &verifyOpts)

	// catch HTTP errors
	if err != nil {
		return verify.CheckResponse{}, verify.CheckErrorResponse{}, err
	}

	// non-zero statuses are also errors
	if result.Status != "0" {
		data, _ := ioutil.ReadAll(resp.Body)

		var errResp verify.CheckErrorResponse
		json.Unmarshal(data, &errResp)
		return result, errResp, nil
	}

	return result, verify.CheckErrorResponse{}, nil
}
