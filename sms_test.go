package nexmo

import "testing"

func TestSendSMS(t *testing.T) {
	_client.SMS.SendSMS(SendSMSRequest{
		To:   "447520615146",
		From: "NEXMOTEST",
		Text: "Nêxmö Tėšt",
		Type: Unicode,
	})
}
