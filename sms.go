package nexmo

import (
	"errors"

	"github.com/antihax/optional"
	"github.com/nexmo-community/nexmo-go/sms"
)

// Client for working with the SMS API
type NexmoSMSClient struct {
	apiKey    string
	apiSecret string
}

// Create a new SMS Client, supplying an Auth to work with
func NewNexmoSMSClient(Auth Auth) *NexmoSMSClient {
	client := new(NexmoSMSClient)
	creds := Auth.GetCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]
	return client
}

type SMSClientOpts struct {
	Config *sms.Configuration
}

// Send an SMS. Give some text to send and the number to send to - there are
// some restrictions on what you can send from, to be safe try using a Nexmo
// number associated with your account
func (client *NexmoSMSClient) Send(from string, to string, text string, opts SMSClientOpts) (sms.Sms, error) {

	config := sms.NewConfiguration()

	// but use the one passed in if we got one
	if opts.Config != nil {
		config = opts.Config
	}

	smsClient := sms.NewAPIClient(config)

	smsOpts := sms.SendAnSmsOpts{}
	smsOpts.Text = optional.NewString(text)
	smsOpts.ApiSecret = optional.NewString(client.apiSecret)

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
