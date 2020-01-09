package nexmo

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestStartVerify(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/verify/json",
		httpmock.NewStringResponder(200, `{
				"request_id": "123",
				"status": "0", 
				"error_text": " "
			}`))

	response, _, err := _client.Verify.Start(StartVerificationRequest{
		Number:      "447520615146",
		Brand:       "NEXMOTEST",
		PINCode:     "1234",
		RequireType: "all",
		WorkflowID:  4,
	})

	assert.Nil(t, err)
	assert.Equal(t, "0", response.Status)
}
