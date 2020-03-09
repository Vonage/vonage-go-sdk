package nexmo

import (
	"context"
	"runtime"

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

func (client *VerifyClient) Request(number string, brand string) (verify.RequestResponse, error) {
	// create the client
	verifyClient := verify.NewAPIClient(client.Config)

	// set up and then parse the options
	verifyOpts := verify.VerifyRequestOpts{}

	// we need context for the API key
	ctx := context.WithValue(context.Background(), verify.ContextAPIKey, verify.APIKey{
		Key: client.apiKey,
	})

	result, _, err := verifyClient.DefaultApi.VerifyRequest(ctx, "json", client.apiSecret, number, brand, &verifyOpts)

	// catch HTTP errors
	if err != nil {
		return verify.RequestResponse{}, err
	}

	return result, nil
}

/*
// Send an SMS. Give some text to send and the number to send to - there are
// some restrictions on what you can send from, to be safe try using a Nexmo
// number associated with your account
func (client *SMSClient) Send(from string, to string, text string, opts SMSOpts) (sms.Sms, error) {

	smsClient := sms.NewAPIClient(client.Config)

	smsOpts := sms.SendAnSmsOpts{}
	smsOpts.Text = optional.NewString(text)
	smsOpts.ApiSecret = optional.NewString(client.apiSecret)

	// check through the opts and send whatever was set
	if opts.ClientRef != "" {
		smsOpts.ClientRef = optional.NewString(opts.ClientRef)
	}

	if opts.Callback != "" {
		smsOpts.Callback = optional.NewString(opts.Callback)
	}

	if opts.Type != "" {
		smsOpts.Type_ = optional.NewString(opts.Type)
	}

	if opts.StatusReportReq != false {
		smsOpts.StatusReportReq = optional.NewBool(opts.StatusReportReq)
	}

	// now send the SMS
	result, _, err := smsClient.DefaultApi.SendAnSms(nil, "json", client.apiKey, from, to, &smsOpts)

	// catch HTTP errors
	if err != nil {
		return sms.Sms{}, err
	}

	// now worry about the status code in the response
	if result.Messages[0].Status != "0" {
		return result, errors.New("Status code: " + result.Messages[0].Status)
	}

	return result, nil
}
*/
