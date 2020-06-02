package nexmo

import (
	"context"
	"runtime"

	"github.com/nexmo-community/nexmo-go/voice"
)

// VoiceClient for working with the Voice API
type VoiceClient struct {
	Config *voice.Configuration
	JWT    string
}

// NewVoiceClient Creates a new Voice Client, supplying an Auth to work with
func NewVoiceClient(Auth Auth) *VoiceClient {
	client := new(VoiceClient)
	creds := Auth.GetCreds()
	client.JWT = creds[0]

	client.Config = voice.NewConfiguration()
	client.Config.UserAgent = "nexmo-go/0.15-dev Go/" + runtime.Version()
	client.Config.AddDefaultHeader("Authorization", "Bearer "+client.JWT)
	// client.Config.BasePath = "http://localhost:4010"
	return client
}

// List your calls
func (client *VoiceClient) GetCalls() (voice.GetCallsResponse, error) {
	// create the client
	voiceClient := voice.NewAPIClient(client.Config)

	// set up and then parse the options
	voiceOpts := voice.GetCallsOpts{}

	ctx := context.Background()
	result, _, err := voiceClient.DefaultApi.GetCalls(ctx, &voiceOpts)

	// catch HTTP errors
	if err != nil {
		e := err.(voice.GenericOpenAPIError)
		return voice.GetCallsResponse{}, e
	}

	return result, nil
}
