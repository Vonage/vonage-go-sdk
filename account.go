package vonage

import (
	"context"

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

type AccountBalanceErrorResponse struct {
	ErrorCode      string
	ErrorCodeLabel string
}

// GetBalance fetches the current balance of the authenticated account, in Euros
func (client *AccountClient) GetBalance() (AccountBalance, AccountBalanceErrorResponse, error) {

	accountClient := account.NewAPIClient(client.Config)

	ctx := context.Background()

	// fetch the balance
	result, _, err := accountClient.BalanceApi.GetAccountBalance(ctx, client.apiKey, client.apiSecret)

	if err != nil {
		return AccountBalance{}, AccountBalanceErrorResponse{}, err
	}

	return AccountBalance(result), AccountBalanceErrorResponse{}, nil
}
