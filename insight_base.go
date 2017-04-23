package nexmo

import "github.com/judy2k/nexmo/sling"

type InsightService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newInsightService(base *sling.Sling, authSet *AuthSet) *InsightService {
	sling := base.Base("https://api.nexmo.com/ni/")
	return &InsightService{
		sling:   sling,
		authSet: authSet,
	}
}

func (c *InsightService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}

