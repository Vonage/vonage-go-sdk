package nexmo

import "testing"

func TestNewClient(t *testing.T) {
	client := NewClient("abcd", "def")
	switch client := client.(type) {
	case *nexmoClient:
		if client.apiKey != "abcd" {
			t.Errorf("client apiKey was incorrect: %s", client.apiKey)
		}
		if client.apiSecret != "def" {
			t.Errorf("client apiSecret was incorrect: %s", client.apiSecret)
		}
	default:
		t.Error("NexmoClient returned incorrect concrete type!")
	}
}

func TestSetBaseURL(t *testing.T) {
	client := NewClient("abcd", "def")
	client.SetBaseURL("a base url")
	switch client := client.(type) {
	case *nexmoClient:
		if client.baseURL != "a base url" {
			t.Errorf("client base URL was incorrect: %s", client.baseURL)
		}
	default:
		t.Error("NexmoClient returned incorrect concrete type!")
	}
}
