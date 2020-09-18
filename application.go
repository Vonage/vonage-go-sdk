package vonage

import (
	"context"
	"fmt"
	"runtime"

	"github.com/vonage/vonage-go-sdk/application"
)

// ApplicationClient for working with the Application API
type ApplicationClient struct {
	Config    *application.Configuration
	apiKey    string
	apiSecret string
}

// NewApplicationClient Creates a new Application Client, supplying an Auth to work with
func NewApplicationClient(Auth Auth) *ApplicationClient {
	client := new(ApplicationClient)
	creds := Auth.GetCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]

	client.Config = application.NewConfiguration()
	client.Config.UserAgent = "vonage-go/0.15-dev Go/" + runtime.Version()
	return client
}

// List your Applications
func (client *ApplicationClient) GetApplications() (application.ApplicationResponseCollection, error) {
	// create the client
	applicationClient := application.NewAPIClient(client.Config)

	AppOpts := application.ListApplicationOpts{}

	ctx := context.WithValue(context.Background(), application.ContextBasicAuth, application.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	result, _, err := applicationClient.DefaultApi.ListApplication(ctx, &AppOpts)
	// fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", err)

	return result, nil
}

// GetApplication returns one application, by app ID
func (client *ApplicationClient) GetApplication(app_id string) (application.ApplicationResponse, error) {
	// create the client
	applicationClient := application.NewAPIClient(client.Config)

	ctx := context.WithValue(context.Background(), application.ContextBasicAuth, application.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	result, _, err := applicationClient.DefaultApi.GetApplication(ctx, app_id)
	// fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", err)

	return result, nil
}
