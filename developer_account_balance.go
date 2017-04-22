package nexmo

import (
	"net/http"
)

func (s *DeveloperService) GetBalance() (*GetBalanceResponse, *http.Response, error) {
	request := new(Credentials)
	s.authSet.ApplyAPICredentials(request)
	response := new(GetBalanceResponse)
	httpResponse, err := s.sling.New().
		Get("account/get-balance").
		QueryStruct(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}
