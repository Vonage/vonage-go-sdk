package nexmo

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateApplication(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/v1/applications/",
		httpmock.NewStringResponder(200, `{
			"id": "aaaaaaaa-bbbb-cccc-dddd-0123456789ab",
			"name": "My Application",
			"voice": {
			  "webhooks": [
				{
				  "endpoint_type": "answer_url",
				  "endpoint": "https://example.com/webhooks/answer",
				  "http_method": "GET"
				}
			  ]
			},
			"messages": {
			  "webhooks": [
				{
				  "endpoint_type": "status_url",
				  "endpoint": "https://example.com/webhooks/status",
				  "http_method": "POST"
				}
			  ]
			},
			"keys": {
			  "public_key": "PUBLIC_KEY",
			  "private_key": "PRIVATE_KEY"
			},
			"_links": {
			  "href": "/v1/applications/aaaaaaaa-bbbb-cccc-dddd-0123456789ab"
			}
		  }`))

	response, _, err := _client.Application.CreateApplication(CreateApplicationRequest{})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "aaaaaaaa-bbbb-cccc-dddd-0123456789ab", response.ID)
	assert.Equal(t, "My Application", response.Name)
	// TODO: Should have some more tests here.
}

func TestCreateApplicationRequest(t *testing.T) {
	b, err := json.Marshal(CreateApplicationRequest{
		Name:      "My Application Name",
		Type:      "voice",
		AnswerURL: "https://api.example.com/answer",
		EventURL:  "https://api.example.com/event",
	})
	assert.NoError(t, err)

	var j map[string]interface{}
	err = json.Unmarshal(b, &j)
	assert.NoError(t, err)
	assert.Equal(t, "My Application Name", j["name"])
	assert.Equal(t, "voice", j["type"])
	assert.Equal(t, "https://api.example.com/answer", j["answer_url"])
	assert.Equal(t, "https://api.example.com/event", j["event_url"])
	_, exists := j["answer_method"]
	assert.False(t, exists)
}
