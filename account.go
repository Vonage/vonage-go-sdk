package vonage

import (
	"context"
	"encoding/json"

	"github.com/antihax/optional"
	"github.com/vonage/vonage-go-sdk/internal/account"
)

// AccountClient for working with the SMS API
type AccountClient struct {
	Config    *account.Configuration
	apiKey    string
	apiSecret string
}

// NewAccountClient Creates a new Account Client, supplying an Auth to work with
func NewAccountClient(Auth Auth) *AccountClient {
	client := new(AccountClient)
	creds := Auth.GetCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]

	// Use a default set of config but make it accessible
	client.Config = account.NewConfiguration()
	client.Config.UserAgent = GetUserAgent()

	// Does not pick up correct server URL from OpenAPI description
	client.Config.BasePath = "https://rest.nexmo.com"
	return client
}

type AccountBalance struct {
	Value      float32
	AutoReload bool
}

type AccountErrorResponse struct {
	ErrorCode      string `json:"error-code,omitempty"`
	ErrorCodeLabel string `json:"error-code-label,omitempty"`
}

// GetBalance fetches the current balance of the authenticated account, in Euros
func (client *AccountClient) GetBalance() (AccountBalance, AccountErrorResponse, error) {

	accountClient := account.NewAPIClient(client.Config)

	ctx := context.Background()

	// fetch the balance
	result, _, err := accountClient.BalanceApi.GetAccountBalance(ctx, client.apiKey, client.apiSecret)

	if err != nil {
		return AccountBalance{}, AccountErrorResponse{}, err
	}

	return AccountBalance(result), AccountErrorResponse{}, nil
}

type AccountConfigSettings struct {
	MoCallbackUrl string
	DrCallbackUrl string
}

type AccountConfigResponse struct {
	MoCallbackUrl      string
	DrCallbackUrl      string
	MaxOutboundRequest int32
	MaxInboundRequest  int32
	MaxCallsPerSecond  int32
}

// SetConfig allows the user to set the URLs for incoming SMS (mo) and delivery receipt (dr) payloads
func (client *AccountClient) SetConfig(config AccountConfigSettings) (AccountConfigResponse, AccountErrorResponse, error) {

	accountClient := account.NewAPIClient(client.Config)

	ctx := context.Background()

	opts := account.ChangeAccountSettingsOpts{}
	if config.MoCallbackUrl != "" {
		opts.MoCallBackUrl = optional.NewString(config.MoCallbackUrl)
	}
	if config.DrCallbackUrl != "" {
		opts.DrCallBackUrl = optional.NewString(config.DrCallbackUrl)
	}

	// update Account settings
	result, _, err := accountClient.ConfigurationApi.ChangeAccountSettings(ctx, client.apiKey, client.apiSecret, &opts)

	if err != nil {
		e := err.(account.GenericOpenAPIError)
		data := e.Body()

		var errResp AccountErrorResponse
		jsonErr := json.Unmarshal(data, &errResp)
		if jsonErr == nil {
			return AccountConfigResponse{}, errResp, err
		}

	}

	return AccountConfigResponse(result), AccountErrorResponse{}, nil
}
