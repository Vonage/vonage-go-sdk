package vonage

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/antihax/optional"
	"github.com/vonage/vonage-go-sdk/internal/verify"
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
	client.Config.UserAgent = GetUserAgent()
	return client
}

// VerifyOpts holds all the optional arguments for the verify request function
type VerifyOpts struct {
	Country       string
	SenderID      string
	CodeLength    int32
	Lg            string
	PinExpiry     int32
	NextEventWait int32
	WorkflowID    int32
}

type VerifyRequestResponse struct {
	// The unique ID of the Verify request. You need this `request_id` for the Verify check.
	RequestId string
	Status    string
}

type VerifyErrorResponse struct {
	RequestId string `json:"request_id"`
	Status    string `json:"status"`
	ErrorText string `json:"error_text,omitempty"`
}

// Request a number is verified for ownership
func (client *VerifyClient) Request(number string, brand string, opts VerifyOpts) (VerifyRequestResponse, VerifyErrorResponse, error) {
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

	if opts.WorkflowID != 0 {
		verifyOpts.WorkflowId = optional.NewInt32(opts.WorkflowID)
	}

	if opts.SenderID != "" {
		verifyOpts.SenderId = optional.NewString(opts.SenderID)
	}

	// we need context for the API key
	ctx := context.WithValue(context.Background(), verify.ContextAPIKey, verify.APIKey{
		Key: client.apiKey,
	})

	result, resp, err := verifyClient.DefaultApi.VerifyRequest(ctx, "json", client.apiSecret, number, brand, &verifyOpts)

	// catch HTTP errors
	if err != nil {
		return VerifyRequestResponse{}, VerifyErrorResponse{}, err
	}

	// non-zero statuses are also errors
	if result.Status != "0" {
		data, _ := ioutil.ReadAll(resp.Body)

		var errResp VerifyErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			return VerifyRequestResponse(result), errResp, nil
		}
	}
	return VerifyRequestResponse(result), VerifyErrorResponse{}, nil
}

type VerifyCheckResponse struct {
	RequestId                  string
	EventId                    string
	Status                     string
	Price                      string
	Currency                   string
	EstimatedPriceMessagesSent string
}

// Check the user-supplied code for this request ID
func (client *VerifyClient) Check(requestID string, code string) (VerifyCheckResponse, VerifyErrorResponse, error) {
	// create the client
	verifyClient := verify.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), verify.ContextAPIKey, verify.APIKey{
		Key: client.apiKey,
	})

	// set up and then parse the options
	verifyOpts := verify.VerifyCheckOpts{}
	result, resp, err := verifyClient.DefaultApi.VerifyCheck(ctx, "json", client.apiSecret, requestID, code, &verifyOpts)

	// catch HTTP errors
	if err != nil {
		return VerifyCheckResponse{}, VerifyErrorResponse{}, err
	}

	// non-zero statuses are also errors
	if result.Status != "0" {
		data, _ := ioutil.ReadAll(resp.Body)

		var errResp VerifyErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			return VerifyCheckResponse(result), errResp, nil
		}
	}

	return VerifyCheckResponse(result), VerifyErrorResponse{}, nil
}

type VerifySearchResponse struct {
	RequestId                  string
	AccountId                  string
	Status                     string
	Number                     string
	Price                      string
	Currency                   string
	SenderId                   string
	DateSubmitted              string
	DateFinalized              string
	FirstEventDate             string
	LastEventDate              string
	Checks                     []verify.SearchResponseChecks
	Events                     []verify.SearchResponseEvents
	EstimatedPriceMessagesSent string
}

// Search for an earlier request by id
func (client *VerifyClient) Search(requestID string) (VerifySearchResponse, VerifyErrorResponse, error) {
	// create the client
	verifyClient := verify.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), verify.ContextAPIKey, verify.APIKey{
		Key: client.apiKey,
	})

	// set up and then parse the options
	verifyOpts := verify.VerifySearchOpts{}
	verifyOpts.RequestId = optional.NewString(requestID)
	result, resp, err := verifyClient.DefaultApi.VerifySearch(ctx, "json", client.apiSecret, &verifyOpts)

	// catch HTTP errors
	if err != nil {
		return VerifySearchResponse{}, VerifyErrorResponse{}, err
	}

	// search failed if we didn't get a request ID
	if result.RequestId == "" {
		data, _ := ioutil.ReadAll(resp.Body)

		var errResp VerifyErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			return VerifySearchResponse{}, errResp, nil
		}
	}

	return VerifySearchResponse(result), VerifyErrorResponse{}, nil
}

type VerifyControlResponse struct {
	Status  string
	Command string
}

// Cancel an in-progress request (check API docs for when this is possible)
func (client *VerifyClient) Cancel(requestID string) (VerifyControlResponse, VerifyErrorResponse, error) {
	// create the client
	verifyClient := verify.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), verify.ContextAPIKey, verify.APIKey{
		Key: client.apiKey,
	})

	result, resp, err := verifyClient.DefaultApi.VerifyControl(ctx, "json", client.apiSecret, requestID, "cancel")

	// catch HTTP errors
	if err != nil {
		return VerifyControlResponse{}, VerifyErrorResponse{}, err
	}

	// search statuses are strings
	if result.Status != "0" {
		data, _ := ioutil.ReadAll(resp.Body)

		var errResp VerifyErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			return VerifyControlResponse(result), errResp, nil
		}
	}

	return VerifyControlResponse(result), VerifyErrorResponse{}, nil
}

// TriggerNextEvent moves on to the next event in the workflow
func (client *VerifyClient) TriggerNextEvent(requestID string) (VerifyControlResponse, VerifyErrorResponse, error) {
	// create the client
	verifyClient := verify.NewAPIClient(client.Config)

	// we need context for the API key
	ctx := context.WithValue(context.Background(), verify.ContextAPIKey, verify.APIKey{
		Key: client.apiKey,
	})

	result, resp, err := verifyClient.DefaultApi.VerifyControl(ctx, "json", client.apiSecret, requestID, "trigger_next_event")

	// catch HTTP errors
	if err != nil {
		return VerifyControlResponse{}, VerifyErrorResponse{}, err
	}

	// search statuses are strings
	if result.Status != "0" {
		data, _ := ioutil.ReadAll(resp.Body)

		var errResp VerifyErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			return VerifyControlResponse(result), errResp, nil
		}
	}

	return VerifyControlResponse(result), VerifyErrorResponse{}, nil
}
