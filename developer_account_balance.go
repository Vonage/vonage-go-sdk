package nexmo

import (
	"net/http"
)

func (s *DeveloperService) GetBalance() (*GetBalanceResponse, *http.Response, error) {
	request := new(Credentials)
	s.authSet.ApplyAPICredentials(request)
	sling := s.sling.New().Get("account/get-balance").QueryStruct(request)

	response := new(GetBalanceResponse)
	httpResponse, err := sling.ReceiveSuccess(response)
	return response, httpResponse, err
}
