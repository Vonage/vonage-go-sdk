package nexmo

import "github.com/nexmo-community/nexmo-go/sling"

// Use Verify API for 2FA, passwordless login, or confirming that a user has given a correct phone number. More information: https://developer.nexmo.com/verify
type VerifyService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newVerifyService(base *sling.Sling, authSet *AuthSet) *VerifyService {
	sling := base.Base("https://api.nexmo.com/verify/")
	return &VerifyService{
		sling:   sling,
		authSet: authSet,
	}
}

func (c *VerifyService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
