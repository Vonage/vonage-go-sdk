package nexmo

import (
	"os"
	"testing"
)

var _client *Client

func TestMain(m *testing.M) {
	os.Exit(func() int {
		_client = initClient()

		return m.Run()
	}())
}

func initClient() *Client {
	apiKey := os.Getenv("NEXMO_API_KEY")
	apiSecret := os.Getenv("NEXMO_API_SECRET")

	if _client != nil {
		return _client
	}

	auth := NewAuthSet()
	auth.SetAPISecret(apiKey, apiSecret)
	_client = New(nil, auth)

	return _client
}
