package nexmo

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type Client struct {
	sling       *sling.Sling
	Insight     *InsightService
	SMS         *SMSService
	Call        *CallService
	Verify      *VerifyService
	Developer   *DeveloperService
	Application *ApplicationService
}

func NewClient(httpClient *http.Client, authSet *AuthSet) *Client {
	base := sling.New().
		Client(httpClient).
		Set("User-Agent", "nexmo-go/2.0 (judy2k)")
	return &Client{
		sling:       base,
		Insight:     newInsightService(base.New(), authSet),
		SMS:         newSMSService(base.New(), authSet),
		Call:        newCallService(base.New(), authSet),
		Verify:      newVerifyService(base.New(), authSet),
		Developer:   newDeveloperService(base.New(), authSet),
		Application: newApplicationService(base.New(), authSet),
	}
}

type APIError struct {
	Status       int64
	ErrorMessage string
}

func (a APIError) Error() string {
	return fmt.Sprintf("%d: %s", a.Status, a.ErrorMessage)
}
