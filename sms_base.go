package nexmo

import "github.com/judy2k/nexmo-go/sling"

type SMSService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

type MessageType string

const (
	Text    MessageType = "text"
	Binary  MessageType = "binary"
	WAPPush MessageType = "wappush"
	Unicode MessageType = "unicode"
)

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
