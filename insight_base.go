package nexmo

import "github.com/dghubble/sling"

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

// Set the base URL for the API request. This is mostly useful for testing.
func (c *InsightService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
