package nexmo

import "github.com/dghubble/sling"

// For working with the Voice API. More information about Voice: https://developer.nexmo.com/voice/voice-api/
type CallService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newCallService(base *sling.Sling, authSet *AuthSet) *CallService {
	sling := base.Base("https://api.nexmo.com/v1/calls/")
	return &CallService{
		sling:   sling,
		authSet: authSet,
	}
}

// Set the base URL for the API requests. Mostly useful for testing.
func (c *CallService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
