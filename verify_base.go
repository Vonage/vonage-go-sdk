package nexmo

import "github.com/judy2k/nexmo-go/sling"

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
