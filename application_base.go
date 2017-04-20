package nexmo

import "github.com/dghubble/sling"

type ApplicationService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newApplicationService(base *sling.Sling, authSet *AuthSet) *ApplicationService {
	sling := base.Base("https://api.nexmo.com/v1/applications/")
	return &ApplicationService{
		sling:   sling,
		authSet: authSet,
	}
}

func (c *ApplicationService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}

