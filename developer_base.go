package nexmo

import "github.com/dghubble/sling"

// Developer API allows configuration of account and balance checking. See also: https://developer.nexmo.com/api/account
type DeveloperService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newDeveloperService(base *sling.Sling, authSet *AuthSet) *DeveloperService {
	sling := base.Base("https://rest.nexmo.com/")
	return &DeveloperService{
		sling:   sling,
		authSet: authSet,
	}
}

// Set the base URL for the API request, useful for testing
func (c *DeveloperService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
