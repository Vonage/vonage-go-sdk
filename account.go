package vonage

import (
	"context"
	"encoding/json"
	"time"

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
	// default config is not correct
	// so override if it's still using the default setting.
	defaultConfig := account.NewConfiguration()
	if defaultConfig.BasePath == client.Config.BasePath {
		client.Config.BasePath = "https://rest.nexmo.com"
	}

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
	// default config is not correct
	// so override if it's still using the default setting.
	defaultConfig := account.NewConfiguration()
	if defaultConfig.BasePath == client.Config.BasePath {
		client.Config.BasePath = "https://rest.nexmo.com"
	}

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

type AccountSecret struct {
	ID        string
	CreatedAt time.Time
}

type AccountSecretCollection struct {
	Secrets []AccountSecret
}

type AccountSecretErrorResponse struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func (client *AccountClient) ListSecrets() (AccountSecretCollection, AccountSecretErrorResponse, error) {
	accountClient := account.NewAPIClient(client.Config)

	ctx := context.WithValue(context.Background(), account.ContextBasicAuth, account.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	// get secrets
	result, _, err := accountClient.SecretManagementApi.RetrieveAPISecrets(ctx, client.apiKey)

	if err != nil {
		e, ok := err.(account.GenericOpenAPIError)
		if ok {
			data := e.Body()

			var errResp AccountSecretErrorResponse
			jsonErr := json.Unmarshal(data, &errResp)
			if jsonErr == nil {
				return AccountSecretCollection{}, errResp, err
			}
			// if we didn't get the expected format but it was an openapi error
			return AccountSecretCollection{}, AccountSecretErrorResponse{}, e
		}
		// something else went wrong
		return AccountSecretCollection{}, AccountSecretErrorResponse{}, err
	}

	// lots and lots of type assertions needed here because of how the API description is structured
	data := result["_embedded"].(map[string]interface{})
	list := data["secrets"].([]interface{})
	var collection AccountSecretCollection

	for i := 0; i < len(list); i++ {
		secret := list[i].(map[string]interface{})
		createdAt, _ := time.Parse(time.RFC3339, secret["created_at"].(string))
		collection.Secrets = append(
			collection.Secrets,
			AccountSecret{ID: secret["id"].(string), CreatedAt: createdAt},
		)
	}
	return collection, AccountSecretErrorResponse{}, nil
}
