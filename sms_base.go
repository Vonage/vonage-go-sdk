package nexmo

import "github.com/nexmo-community/nexmo-go/sling"

type SMSService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newSMSService(base *sling.Sling, authSet *AuthSet) *SMSService {
	sling := base.Base("https://rest.nexmo.com/sms/")
	return &SMSService{
		sling:   sling,
		authSet: authSet,
	}
}

func (c *SMSService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
