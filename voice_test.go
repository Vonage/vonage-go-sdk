package vonage

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/vonage/vonage-go-sdk/ncco"
)

func TestVoiceNewClient(*testing.T) {
	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	NewVoiceClient(auth)
}

func TestVoiceGetCalls(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/v1/calls/",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "count": 100,
  "page_size": 10,
  "record_index": 0,
  "_links": {
    "self": {
      "href": "/calls?page_size=10&record_index=20&order=asc"
    }
  },
  "_embedded": {
    "calls": [
      {
        "_links": {
          "self": {
            "href": "/calls/63f61863-4a51-4f6b-86e1-46edebcf9356"
          }
        },
        "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356",
        "conversation_uuid": "CON-f972836a-550f-45fa-956c-12a2ab5b7d22",
        "to": {
          "type": "phone",
          "number": "447700900000"
        },
        "from": {
          "type": "phone",
          "number": "447700900001"
        },
        "status": "started",
        "direction": "outbound",
        "rate": "0.39",
        "price": "23.40",
        "duration": "60",
        "start_time": "2020-01-01 12:00:00",
        "end_time": "2020-01-01 12:00:00",
        "network": "65512"
      }
    ]
  }
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	response, _, _ := client.GetCalls()
	message := response.Embedded.Calls[0].Uuid + " status: " + response.Embedded.Calls[0].Status
	if message != "63f61863-4a51-4f6b-86e1-46edebcf9356 status: started" {
		t.Errorf("Voice GetCalls failed")
	}
}

func TestVoiceGetCallsNoAuth(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/v1/calls/",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(401, `
{"type":"UNAUTHORIZED","error_title":"Unauthorized"}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	_, _, http_error := client.GetCalls()
	if http_error == nil {
		t.Errorf("Voice GetCalls with faily Auth failed")
	}
}
func TestVoiceGetCall(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/v1/calls/1234567890",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "_links": {
    "self": {
      "href": "/calls/63f61863-4a51-4f6b-86e1-46edebcf9356"
    }
  },
  "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356",
  "conversation_uuid": "CON-f972836a-550f-45fa-956c-12a2ab5b7d22",
  "to": {
    "type": "phone",
    "number": "447700900000"
  },
  "from": {
    "type": "phone",
    "number": "447700900001"
  },
  "status": "started",
  "direction": "outbound",
  "rate": "0.39",
  "price": "23.40",
  "duration": "60",
  "start_time": "2020-01-01 12:00:00",
  "end_time": "2020-01-01 12:00:00",
  "network": "65512"
}`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	response, _, _ := client.GetCall("1234567890")
	message := response.Uuid + " status: " + response.Status
	if message != "63f61863-4a51-4f6b-86e1-46edebcf9356 status: started" {
		t.Errorf("Voice GetCall (singular) failed")
	}
}

func TestVoiceMakeCallWithNcco(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/v1/calls/",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, `
{
  "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356",
  "status": "started",
  "direction": "outbound",
  "conversation_uuid": "CON-f972836a-550f-45fa-956c-12a2ab5b7d22"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	from := CallFrom{Type: "phone", Number: "447770007777"}
	to := CallTo{Type: "phone", Number: "447770007788"}

	MyNcco := ncco.Ncco{}
	talk := ncco.TalkAction{Text: "This is the golang library, calling to say hello", VoiceName: "Nicole"}
	MyNcco.AddAction(talk)

	result, _, _ := client.CreateCall(CreateCallOpts{From: from, To: to, Ncco: MyNcco})
	message := result.Uuid + " <-- call ID started"
	if message != "63f61863-4a51-4f6b-86e1-46edebcf9356 <-- call ID started" {
		t.Errorf("Voice create call with Ncco failed")
	}
}

func TestVoiceMakeCallWithAnswerUrl(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/v1/calls/",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, `
{
  "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356",
  "status": "started",
  "direction": "outbound",
  "conversation_uuid": "CON-f972836a-550f-45fa-956c-12a2ab5b7d22"
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	from := CallFrom{Type: "phone", Number: "447770007777"}
	to := CallTo{Type: "phone", Number: "447770007788"}
	answer := []string{"https://example.com/answer"}

	result, _, _ := client.CreateCall(CreateCallOpts{From: from, To: to, AnswerUrl: answer})
	message := result.Uuid + " <-- call ID started"
	if message != "63f61863-4a51-4f6b-86e1-46edebcf9356 <-- call ID started" {
		t.Errorf("Voice create call with AnswerUrl failed")
	}
}

func TestVoiceTransferCallWithAnswerUrl(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	url := []string{"https://example.com/answer"}

	result, _, _ := client.TransferCall(TransferCallOpts{Uuid: "abcdef01-2222-3333-4444-9876543210ab", AnswerUrl: url})
	message := "Status: " + result.Status
	if message != "Status: 0" {
		t.Errorf("Voice transfer call with AnswerUrl failed")
	}
}

func TestVoiceTransferCallWithNcco(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	MyNcco := ncco.Ncco{}
	talk := ncco.TalkAction{Text: "This is the golang library, calling to interrupt the other call", VoiceName: "Nicole"}
	MyNcco.AddAction(talk)

	result, _, _ := client.TransferCall(TransferCallOpts{Uuid: "abcdef01-2222-3333-4444-9876543210ab", Ncco: MyNcco})
	message := "Status: " + result.Status
	if message != "Status: 0" {
		t.Errorf("Voice transfer call with AnswerUrl failed")
	}
}

func TestVoiceHangup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.Hangup("abcdef01-2222-3333-4444-9876543210ab")
	message := "Status: " + result.Status
	if message != "Status: 0" {
		t.Errorf("Voice hangup failed")
	}
}

func TestVoiceMute(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.Mute("abcdef01-2222-3333-4444-9876543210ab")
	message := "Status: " + result.Status
	if message != "Status: 0" {
		t.Errorf("Voice Mute failed")
	}
}

func TestVoiceUnmute(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.Unmute("abcdef01-2222-3333-4444-9876543210ab")
	message := "Status: " + result.Status
	if message != "Status: 0" {
		t.Errorf("Voice Unmute failed")
	}
}

func TestVoiceEarmuff(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.Earmuff("abcdef01-2222-3333-4444-9876543210ab")
	message := "Status: " + result.Status
	if message != "Status: 0" {
		t.Errorf("Voice Earmuff failed")
	}
}

func TestVoiceUnearmuff(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.Unearmuff("abcdef01-2222-3333-4444-9876543210ab")
	message := "Status: " + result.Status
	if message != "Status: 0" {
		t.Errorf("Voice Unearmuff failed")
	}
}

func TestVoiceStartStream(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab/stream",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "message": "Stream started",
  "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356"
}
	`,
			)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.PlayAudioStream("abcdef01-2222-3333-4444-9876543210ab", "https://example.com/music.mp3", PlayAudioOpts{})
	if result.Message != "Stream started" {
		t.Errorf("Voice Start Audio Stream failed")
	}
}

func TestVoiceStopStream(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab/stream",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "message": "Stream stopped",
  "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356"
}
	`,
			)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.StopAudioStream("abcdef01-2222-3333-4444-9876543210ab")
	if result.Message != "Stream stopped" {
		t.Errorf("Voice Stop Audio Stream failed")
	}
}

func TestVoiceStartTts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab/talk",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "message": "Talk started",
  "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356"
}
	`,
			)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.PlayTts("abcdef01-2222-3333-4444-9876543210ab", "Hello world", PlayTtsOpts{})
	if result.Message != "Talk started" {
		t.Errorf("Voice Start Talk failed")
	}
}

func TestVoiceStopTts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab/talk",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "message": "Talk stopped",
  "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356"
}
	`,
			)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.StopTts("abcdef01-2222-3333-4444-9876543210ab")
	if result.Message != "Talk stopped" {
		t.Errorf("Voice Stop TTS failed")
	}
}

func TestVoiceDtmf(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.nexmo.com/v1/calls/abcdef01-2222-3333-4444-9876543210ab/dtmf",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "message": "DTMF sent",
  "uuid": "63f61863-4a51-4f6b-86e1-46edebcf9356"
}
	`,
			)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth, _ := CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", []byte("imagine this is a private key"))
	client := NewVoiceClient(auth)

	result, _, _ := client.PlayDtmf("abcdef01-2222-3333-4444-9876543210ab", "752")
	if result.Message != "DTMF sent" {
		t.Errorf("Voice DTMF send failed")
	}
}
