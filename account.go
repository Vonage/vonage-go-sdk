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

// GetSecret retrieves data about a single account secret
func (client *AccountClient) GetSecret(id string) (AccountSecret, AccountSecretErrorResponse, error) {
	accountClient := account.NewAPIClient(client.Config)

	ctx := context.WithValue(context.Background(), account.ContextBasicAuth, account.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	// get secrets
	result, _, err := accountClient.SecretManagementApi.RetrieveAPISecret(ctx, client.apiKey, id)

	if err != nil {
		e, ok := err.(account.GenericOpenAPIError)
		if ok {
			data := e.Body()

			var errResp AccountSecretErrorResponse
			jsonErr := json.Unmarshal(data, &errResp)
			if jsonErr == nil {
				return AccountSecret{}, errResp, err
			}
			// if we didn't get the expected format but it was an openapi error
			return AccountSecret{}, AccountSecretErrorResponse{}, e
		}
		// something else went wrong
		return AccountSecret{}, AccountSecretErrorResponse{}, err
	}

	createdAt, _ := time.Parse(time.RFC3339, result.CreatedAt)
	secret := AccountSecret{ID: result.Id, CreatedAt: createdAt}
	return secret, AccountSecretErrorResponse{}, nil
}

// CreateSecret adds an additional secret to the account (the number of secrets allowed is limited)
func (client *AccountClient) CreateSecret(secret string) (AccountSecret, AccountSecretErrorResponse, error) {
	accountClient := account.NewAPIClient(client.Config)

	ctx := context.WithValue(context.Background(), account.ContextBasicAuth, account.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	request := account.CreateSecretRequest{Secret: secret}

	result, _, err := accountClient.SecretManagementApi.CreateAPISecret(ctx, client.apiKey, request)

	if err != nil {
		e, ok := err.(account.GenericOpenAPIError)
		if ok {
			data := e.Body()

			var errResp AccountSecretErrorResponse
			jsonErr := json.Unmarshal(data, &errResp)
			if jsonErr == nil {
				return AccountSecret{}, errResp, err
			}
			// if we didn't get the expected format but it was an openapi error
			return AccountSecret{}, AccountSecretErrorResponse{}, e
		}
		// something else went wrong
		return AccountSecret{}, AccountSecretErrorResponse{}, err
	}

	createdAt, _ := time.Parse(time.RFC3339, result.CreatedAt)
	new_secret := AccountSecret{ID: result.Id, CreatedAt: createdAt}
	return new_secret, AccountSecretErrorResponse{}, nil
}

// DeleteSecret adds an additional secret to the account (the number of secrets allowed is limited)
func (client *AccountClient) DeleteSecret(id string) (bool, AccountSecretErrorResponse, error) {
	accountClient := account.NewAPIClient(client.Config)

	ctx := context.WithValue(context.Background(), account.ContextBasicAuth, account.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	// this one only returns two values because it's a 204 so no actual response
	_, err := accountClient.SecretManagementApi.RevokeAPISecret(ctx, client.apiKey, id)

	if err != nil {
		e, ok := err.(account.GenericOpenAPIError)
		if ok {
			data := e.Body()

			var errResp AccountSecretErrorResponse
			jsonErr := json.Unmarshal(data, &errResp)
			if jsonErr == nil {
				return false, errResp, err
			}
			// if we didn't get the expected format but it was an openapi error
			return false, AccountSecretErrorResponse{}, e
		}
		// something else went wrong
		return false, AccountSecretErrorResponse{}, err
	}

	return true, AccountSecretErrorResponse{}, nil
}
