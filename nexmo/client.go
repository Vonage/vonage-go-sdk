package nexmo

import "net/http"

type nexmoClient struct {
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient *http.Client
}

func (client *nexmoClient) SetBaseURL(baseURL string) {
	client.baseURL = baseURL
}
