package nexmo

import "github.com/nexmo-community/nexmo-go/sling"

// Number Insights provides information at varying levels of detail (basic/standard/advanced) about a phone number. For more information, visit the developer documentation https://developer.nexmo.com/number-insight
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
