package nexmo

import (
	"context"
	"encoding/json"
	"runtime"

	"github.com/antihax/optional"
	"github.com/nexmo-community/nexmo-go/voice"
)

// VoiceClient for working with the Voice API
type VoiceClient struct {
	Config *voice.Configuration
	JWT    string
}

// NewVoiceClient Creates a new Voice Client, supplying an Auth to work with
func NewVoiceClient(Auth Auth) *VoiceClient {
	client := new(VoiceClient)
	creds := Auth.GetCreds()
	client.JWT = creds[0]

	client.Config = voice.NewConfiguration()
	client.Config.UserAgent = "nexmo-go/0.15-dev Go/" + runtime.Version()
	client.Config.AddDefaultHeader("Authorization", "Bearer "+client.JWT)
	// client.Config.BasePath = "http://localhost:4010"
	return client
}

// List your calls
func (client *VoiceClient) GetCalls() (voice.GetCallsResponse, VoiceErrorResponse, error) {
	// create the client
	voiceClient := voice.NewAPIClient(client.Config)

	// set up and then parse the options
	voiceOpts := voice.GetCallsOpts{}

	ctx := context.Background()
	result, _, err := voiceClient.DefaultApi.GetCalls(ctx, &voiceOpts)

	// catch HTTP errors
	if err != nil {
		return voice.GetCallsResponse{}, VoiceErrorResponse{}, err
	}

	return result, VoiceErrorResponse{}, nil
}

// CreateCallOpts: Options for creating a call
type CreateCallOpts struct {
	From CallFrom
	To   CallTo
	Ncco Ncco
}

// CallFrom details of the caller
type CallFrom struct {
	Type       string
	Number     string
	DtmfAnswer string
}

// CallTo details of the callee
type CallTo struct {
	Type        string
	Number      string
	DtmfAnswer  string
	Uri         string
	Extension   string
	ContentType string
	Headers     map[string]string
}

// VoiceErrorResponse is a container for error types since we can get more than
// one type of error back and they have incompatible data types
type VoiceErrorResponse struct {
	Error interface{}
}

// VoiceErrorInvalidParamsResponse can come with a 400 response if
// it is caused by some invalid_parameters
type VoiceErrorInvalidParamsResponse struct {
	Type              int                 `json:"type, omitempty"`
	Title             string              `json:"title, omitempty"`
	Detail            string              `json:"detail, omitempty"`
	Instance          string              `json:"instance, omitempty"`
	InvalidParameters []map[string]string `json:"invalid_parameters, omitempty"`
}

// VoiceErrorGeneralResponse covers some common error types that come
// from the webserver/gateway rather than the API itself
type VoiceErrorGeneralResponse struct {
	Type  string `json:"type, omitempty"`
	Title string `json:"error_title, omitempty"`
}

// CreateCall Makes a phone call given the from/to details and
// either an AnswerURL or an NCCO
func (client *VoiceClient) CreateCall(opts CreateCallOpts) (voice.CreateCallResponse, VoiceErrorResponse, error) {

	voiceClient := voice.NewAPIClient(client.Config)
	voiceCallOpts := voice.CreateCallRequest{}
	// assuming phone start with
	var to []interface{}
	to = append(to, voice.EndpointPhone{Type: "phone", Number: opts.To.Number})
	voiceCallOpts.To = to
	// from has to be a phone
	voiceCallOpts.From = voice.EndpointPhone{Type: "phone", Number: opts.From.Number}

	// ncco has its own features
	voiceCallOpts.Ncco = opts.Ncco.GetActions()

	callOpts := optional.NewInterface(voiceCallOpts)

	ctx := context.Background()
	createCallOpts := &voice.CreateCallOpts{CreateCallRequest: callOpts}
	result, _, err := voiceClient.DefaultApi.CreateCall(ctx, createCallOpts)

	if err != nil {
		e := err.(voice.GenericOpenAPIError)
		errorType := e.Error()
		data := e.Body()

		// now handle the errors we know we might get
		if errorType == "401 Unauthorized" {
			var errResp VoiceErrorGeneralResponse
			json.Unmarshal(data, &errResp)
			return voice.CreateCallResponse{}, VoiceErrorResponse{Error: errResp}, err
		} else if errorType == "404 Not Found" {
			var errResp VoiceErrorInvalidParamsResponse
			json.Unmarshal(data, &errResp)
			return voice.CreateCallResponse{}, VoiceErrorResponse{Error: errResp}, err
		} else if errorType == "400 Bad Request" {
			var errResp VoiceErrorInvalidParamsResponse
			json.Unmarshal(data, &errResp)
			return voice.CreateCallResponse{}, VoiceErrorResponse{Error: errResp}, err
		} else {
			return voice.CreateCallResponse{}, VoiceErrorResponse{}, err
		}

	}
	return result, VoiceErrorResponse{}, nil
}
