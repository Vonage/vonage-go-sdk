package nexmo

import "github.com/nexmo-community/nexmo-go/sling"

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

func (c *CallService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
