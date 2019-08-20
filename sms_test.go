package nexmo

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSendSMS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://rest.nexmo.com/sms/json",
		httpmock.NewStringResponder(200, `{
				"message-count": "1",
				"messages": [{
					"to": "447520615146",
					"message-id": "140000005494DDEB",
					"status": "0",
					"remaining-balance": "54.42941782",
					"message-price": "0.03330000",
					"network": "23409"
				}]
			}`))

	response, _, err := _client.SMS.SendSMS(SendSMSRequest{
		To:   "447520615146",
		From: "NEXMOTEST",
		Text: "Nêxmö Tėšt",
		Type: "unicode",
	})

	assert.Nil(t, err)
	assert.Equal(t, "1", response.MessageCount)
}

func TestSMSRequest(t *testing.T) {
	b, err := json.Marshal(SendSMSRequest{
		To:   "447520615146",
		From: "NEXMOTEST",
		Text: "Nêxmö Tėšt",
		Type: MessageTypeUnicode,
	})
	assert.NoError(t, err)

	var j map[string]interface{}
	err = json.Unmarshal(b, &j)
	assert.NoError(t, err)
	assert.Equal(t, "unicode", j["type"])
	assert.Equal(t, "Nêxmö Tėšt", j["text"])
}
