package vonage

import (
	"context"
	"errors"
	"runtime"

	"github.com/antihax/optional"
	"github.com/vonage/vonage-go-sdk/sms"
)

// SMSClient for working with the SMS API
type SMSClient struct {
	Config    *sms.Configuration
	apiKey    string
	apiSecret string
}

// NewSMSClient Creates a new SMS Client, supplying an Auth to work with
func NewSMSClient(Auth Auth) *SMSClient {
	client := new(SMSClient)
	creds := Auth.GetCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]

	// Use a default set of config but make it accessible
	client.Config = sms.NewConfiguration()
	client.Config.UserAgent = "vonage-go/0.15-dev Go/" + runtime.Version()
	return client
}

// SMSOpts holds all the optional values that can be set when sending an SMS, check the https://developer.nexmo.com/api/sms API reference for more information
type SMSOpts struct {
	StatusReportReq bool
	Callback        string
	Type            string
	ClientRef       string
}

type Sms struct {
	// The amount of messages in the request
	MessageCount string
	Messages     []sms.Message
}

// Send an SMS. Give some text to send and the number to send to - there are
// some restrictions on what you can send from, to be safe try using a Nexmo
// number associated with your account
func (client *SMSClient) Send(from string, to string, text string, opts SMSOpts) (Sms, error) {

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

	if opts.StatusReportReq {
		smsOpts.StatusReportReq = optional.NewBool(opts.StatusReportReq)
	}

	ctx := context.Background()

	// now send the SMS
	result, _, err := smsClient.DefaultApi.SendAnSms(ctx, "json", client.apiKey, from, to, &smsOpts)

	// catch HTTP errors
	if err != nil {
		return Sms{}, err
	}

	// now worry about the status code in the response
	if result.Messages[0].Status != "0" {
		return Sms(result), errors.New("Status code: " + result.Messages[0].Status)
	}

	return Sms(result), nil
}
