package nexmo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateApplicationRequest struct {
	Credentials
	Name         string `json:"name"`
	Type         string `json:"type"`
	AnswerURL    string `json:"answer_url"`
	AnswerMethod string `json:"answer_method,omitempty"`
	EventURL     string `json:"event_url"`
	EventMethod  string `json:"event_method,omitempty"`
}

type ApplicationConfiguration struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Voice struct {
		Webhooks []struct {
			EndpointType string `json:"endpoint_type"`
			Endpoint     string `json:"endpoint"`
			HTTPMethod   string `json:"http_method"`
		} `json:"webhooks"`
	}
	Keys struct {
		PublicKey  string `json:"public_key"`
		PrivateKey string `json:"private_key"`
	} `json:"keys"`
	Links Links `json:"_links"`
}

type CreateApplicationResponse ApplicationConfiguration

// Create a new application in your Nexmo account
func (s *ApplicationService) CreateApplication(request CreateApplicationRequest) (*CreateApplicationResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(CreateApplicationResponse)
	httpResponse, err := s.sling.New().
		Post("").
		BodyJSON(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

type ListApplicationsRequest struct {
	Credentials
	PageSize  int64 `url:"page_size,omitempty"`
	PageIndex int64 `url:"page_index,omitempty"`
}

type ListApplicationsResponse struct {
	Count    int64 `json:"count,omitempty"`
	PageSize int64 `json:"page_size,omitempty"`
	Embedded struct {
		Applications []ApplicationConfiguration `json:"applications"`
	} `json:"_embedded"`
	Links Links `json:"_links"`
}

// List the applications on the Nexmo account
func (s *ApplicationService) ListApplications(request ListApplicationsRequest) (*ListApplicationsResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(ListApplicationsResponse)
	httpResponse, err := s.sling.New().
		Get("").
		QueryStruct(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

type GetApplicationRequest struct {
	// Created with embedded Credentials, so this will support setApiCredentials
	// (if we alias to Credentials, we lose the implementation)
	Credentials
}

type GetApplicationResponse ApplicationConfiguration

// Fetch a specific application's details
func (s *ApplicationService) GetApplication(id string) (*GetApplicationResponse, *http.Response, error) {
	request := GetApplicationRequest{}
	s.authSet.ApplyAPICredentials(&request)
	response := new(GetApplicationResponse)
	httpResponse, err := s.sling.New().
		Get(id).
		QueryStruct(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

type ModifyApplicationRequest CreateApplicationRequest

type ModifyApplicationResponse ApplicationConfiguration

// Update an existing application by applying changed config to it
func (s *ApplicationService) ModifyApplication(id string, request ModifyApplicationRequest) (*ModifyApplicationResponse, *http.Response, error) {
	s.authSet.ApplyAPICredentials(&request)
	response := new(ModifyApplicationResponse)

	httpResponse, err := s.sling.New().
		Put(id).
		BodyJSON(request).
		ReceiveSuccess(response)
	return response, httpResponse, err
}

// Destroy an application
func (s *ApplicationService) DeleteApplication(id string) (*http.Response, error) {
	credentials := Credentials{}
	s.authSet.ApplyAPICredentials(&credentials)
	httpResponse, err := s.sling.New().
		Delete(id).
		QueryStruct(credentials).
		ReceiveSuccess(new(interface{}))
	return httpResponse, err
}

// TODO: This is vaguely useful. Work out what to do with it.
func jsonError(e interface{}) error {
	// TODO: Need to remove this function
	rep, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	return fmt.Errorf("%s", string(rep))
}
