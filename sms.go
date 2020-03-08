package nexmo

import (
	"errors"
	"runtime"

	"github.com/antihax/optional"
	"github.com/nexmo-community/nexmo-go/sms"
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
	client.Config.UserAgent = "nexmo-go/0.15-dev Go/" + runtime.Version()
	return client
}

// SMSOpts holds all the optional values that can be set when sending an SMS, check the https://developer.nexmo.com/api/sms API reference for more information
type SMSOpts struct {
	// **Advanced**: The duration in milliseconds the delivery of an SMS will be attempted.§§ By default Nexmo attempt delivery for 72 hours, however the maximum effective value depends on the operator and is typically 24 - 48 hours. We recommend this value should be kept at its default or at least 30 minutes.
	Ttl int32
	// **Advanced**: Boolean indicating if you like to receive a [Delivery Receipt](https://developer.nexmo.com/messaging/sms/building-blocks/receive-a-delivery-receipt).
	StatusReportReq bool
	// **Advanced**: The webhook endpoint the delivery receipt for this sms is sent to. This parameter overrides the webhook endpoint you set in Dashboard.
	Callback string
	// **Advanced**: The Data Coding Scheme value of the message
	MessageClass int32
	// **Advanced**: The format of the message body
	Type string
	// **Advanced**: Hex encoded binary data. Depends on `type` parameter having the value `binary`.
	Body string
	// **Advanced**: Your custom Hex encoded [User Data Header](https://en.wikipedia.org/wiki/User_Data_Header). Depends on `type` parameter having the value `binary`.
	Udh string
	// **Advanced**: The value of the [protocol identifier](https://en.wikipedia.org/wiki/GSM_03.40#Protocol_Identifier) to use. Ensure that the value is aligned with `udh`.
	ProtocolId int32
	// **Advanced**: You can optionally include your own reference of up to 40 characters.
	ClientRef string
	// **Advanced**: An optional string used to identify separate accounts using the SMS endpoint for billing purposes. To use this feature, please email [support@nexmo.com](mailto:support@nexmo.com)
	AccountRef string
}

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
