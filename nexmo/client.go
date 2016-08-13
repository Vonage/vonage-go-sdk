package nexmo

import "time"

type nexmoClient struct {
	apiKey            string
	apiSecret         string
	baseURL           string
	connectionTimeout time.Duration
	soTimeout         time.Duration
}

func (client *nexmoClient) SetBaseURL(baseURL string) {
	client.baseURL = baseURL
}

func (client *nexmoClient) SetConnectionTimeout(timeout time.Duration) {
	client.connectionTimeout = timeout
}

func (client *nexmoClient) SetSoTimeout(timeout time.Duration) {
	client.soTimeout = timeout
}
