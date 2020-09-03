package vonage

import (
	"context"
	"encoding/json"
	"errors"
	"runtime"

	"github.com/antihax/optional"
	"github.com/vonage/vonage-go-sdk/voice"
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
	client.Config.UserAgent = "vonage-go/0.15-dev Go/" + runtime.Version()
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
	result, _, err := voiceClient.CallsApi.GetCalls(ctx, &voiceOpts)

	// catch HTTP errors
	if err != nil {
		return voice.GetCallsResponse{}, VoiceErrorResponse{}, err
	}

	return result, VoiceErrorResponse{}, nil
}

// GetCall for the details of a specific call
func (client *VoiceClient) GetCall(uuid string) (voice.GetCallResponse, VoiceErrorResponse, error) {
	// create the client
	voiceClient := voice.NewAPIClient(client.Config)

	ctx := context.Background()
	result, _, err := voiceClient.CallsApi.GetCall(ctx, uuid)

	// catch HTTP errors
	if err != nil {
		return voice.GetCallResponse{}, VoiceErrorResponse{}, err
	}

	return result, VoiceErrorResponse{}, nil
}

// CreateCallOpts: Options for creating a call
type CreateCallOpts struct {
	From             CallFrom
	To               CallTo
	Ncco             Ncco
	AnswerUrl        []string
	AnswerMethod     string
	EventUrl         []string
	EventMethod      string
	MachineDetection string
	LengthTimer      int32
	RingingTimer     int32
}

// CallFrom details of the caller
type CallFrom struct {
	Type   string
	Number string
}

// CallTo details of the callee
type CallTo struct {
	Type       string
	Number     string
	DtmfAnswer string
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

func (client *VoiceClient) createCallCommon(opts CreateCallOpts) voice.CreateCallRequestBase {

	var target voice.CreateCallRequestBase

	// assuming phone to start with, this needs other endpoints added later
	var to []voice.EndpointPhoneTo
	to_phone := voice.EndpointPhoneTo{Type: "phone", Number: opts.To.Number}
	if opts.To.DtmfAnswer != "" {
		to_phone.DtmfAnswer = opts.To.DtmfAnswer
	}
	to = append(to, to_phone)
	target.To = to

	// from has to be a phone
	target.From = voice.EndpointPhoneFrom{Type: "phone", Number: opts.From.Number}

	// event settings
	if len(opts.EventUrl) > 0 {
		target.EventUrl = opts.EventUrl
		if opts.EventMethod != "" {
			target.EventMethod = opts.EventMethod
		}
	}

	// other fields
	if opts.MachineDetection != "" {
		target.MachineDetection = opts.MachineDetection
	}

	if opts.RingingTimer != 0 {
		target.RingingTimer = opts.RingingTimer
	}

	if opts.LengthTimer != 0 {
		target.LengthTimer = opts.LengthTimer
	}

	return target
}

// CreateCall Makes a phone call given the from/to details and an NCCO or an Answer URL
func (client *VoiceClient) CreateCall(opts CreateCallOpts) (voice.CreateCallResponse, VoiceErrorResponse, error) {
	voiceClient := voice.NewAPIClient(client.Config)
	// use the same validation regardless of which type of call this is
	commonFields := client.createCallCommon(opts)

	// ncco has its own features
	if len(opts.Ncco.GetActions()) > 0 {
		// copy the common fields into the appropriate struct
		voiceCallOpts := voice.CreateCallRequestNcco{}
		voiceCallOpts.To = commonFields.To
		voiceCallOpts.From = commonFields.From
		voiceCallOpts.EventUrl = commonFields.EventUrl
		voiceCallOpts.EventMethod = commonFields.EventMethod
		voiceCallOpts.MachineDetection = commonFields.MachineDetection
		voiceCallOpts.LengthTimer = commonFields.LengthTimer
		voiceCallOpts.RingingTimer = commonFields.RingingTimer

		// add NCCO
		voiceCallOpts.Ncco = opts.Ncco.GetActions()

		callOpts := optional.NewInterface(voiceCallOpts)

		ctx := context.Background()
		createCallOpts := &voice.CreateCallOpts{Opts: callOpts}
		NccoResult, _, NccoErr := voiceClient.CallsApi.CreateCall(ctx, createCallOpts)
		return client.handleCreateCallErrors(NccoResult, NccoErr)
	} else if len(opts.AnswerUrl) > 0 {
		voiceCallOpts := voice.CreateCallRequestAnswerUrl{}
		// copy the common fields into the appropriate struct
		voiceCallOpts.To = commonFields.To
		voiceCallOpts.From = commonFields.From
		voiceCallOpts.EventUrl = commonFields.EventUrl
		voiceCallOpts.EventMethod = commonFields.EventMethod
		voiceCallOpts.MachineDetection = commonFields.MachineDetection
		voiceCallOpts.LengthTimer = commonFields.LengthTimer
		voiceCallOpts.RingingTimer = commonFields.RingingTimer

		// answer details
		voiceCallOpts.AnswerUrl = opts.AnswerUrl
		if opts.AnswerMethod != "" {
			voiceCallOpts.AnswerMethod = opts.AnswerMethod
		}

		callOpts := optional.NewInterface(voiceCallOpts)

		ctx := context.Background()
		createCallOpts := &voice.CreateCallOpts{Opts: callOpts}
		AnswerResult, _, AnswerErr := voiceClient.CallsApi.CreateCall(ctx, createCallOpts)
		return client.handleCreateCallErrors(AnswerResult, AnswerErr)
	}

	// this is a backstop, we shouldn't end up here
	return voice.CreateCallResponse{}, VoiceErrorResponse{}, errors.New("Unsupported combination of parameters, supply an answer URL or valid NCCO")
}

func (client *VoiceClient) handleCreateCallErrors(result voice.CreateCallResponse, err error) (voice.CreateCallResponse, VoiceErrorResponse, error) {
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

// TransferCallOpts: Options for transferring a call
type TransferCallOpts struct {
	Uuid      string
	Ncco      Ncco
	AnswerUrl []string
}

type ModifyCallResponse struct {
	Status string
}

type TransferDestinationUrl struct {
	Type string   `json:"type"`
	Url  []string `json:"url"`
}

type TransferWithUrlOpts struct {
	Action      string                 `json:"action"`
	Destination TransferDestinationUrl `json:"destination"`
}

type TransferDestinationNcco struct {
	Type string        `json:"type"`
	Ncco []interface{} `json:"ncco"`
}

type TransferWithNccoOpts struct {
	Action      string                  `json:"action"`
	Destination TransferDestinationNcco `json:"destination"`
}

// TransferCall wraps the Modify Call API endpoint
func (client *VoiceClient) TransferCall(opts TransferCallOpts) (ModifyCallResponse, VoiceErrorResponse, error) {
	// create the client
	voiceClient := voice.NewAPIClient(client.Config)

	if len(opts.AnswerUrl) > 0 {
		destination := TransferDestinationUrl{Type: "ncco", Url: opts.AnswerUrl}
		transfer := TransferWithUrlOpts{Action: "transfer", Destination: destination}
		modifyCallOpts := voice.ModifyCallOpts{Opts: optional.NewInterface(transfer)}
		ctx := context.Background()
		response, err := voiceClient.CallsApi.UpdateCall(ctx, opts.Uuid, &modifyCallOpts)
		if err != nil {
			e := err.(voice.GenericOpenAPIError)
			data := e.Body()
			errorType := e.Error()
			if errorType == "400 Bad Request" {
				var errResp VoiceErrorInvalidParamsResponse
				json.Unmarshal(data, &errResp)
				return ModifyCallResponse{}, VoiceErrorResponse{Error: errResp}, err
			}
			return ModifyCallResponse{}, VoiceErrorResponse{Error: response}, err
		} else {
			// not a whole lot to return as it's a 204, this branch is success
			return ModifyCallResponse{Status: "0"}, VoiceErrorResponse{}, nil
		}
	} else if len(opts.Ncco.GetActions()) > 0 {
		destination := TransferDestinationNcco{Type: "ncco", Ncco: opts.Ncco.GetActions()}
		transfer := TransferWithNccoOpts{Action: "transfer", Destination: destination}
		modifyCallOpts := voice.ModifyCallOpts{Opts: optional.NewInterface(transfer)}
		ctx := context.Background()
		response, err := voiceClient.CallsApi.UpdateCall(ctx, opts.Uuid, &modifyCallOpts)
		if err != nil {
			e := err.(voice.GenericOpenAPIError)
			data := e.Body()
			errorType := e.Error()
			if errorType == "400 Bad Request" {
				var errResp VoiceErrorInvalidParamsResponse
				json.Unmarshal(data, &errResp)
				return ModifyCallResponse{}, VoiceErrorResponse{Error: errResp}, err
			}
			return ModifyCallResponse{}, VoiceErrorResponse{Error: response}, err
		} else {
			// not a whole lot to return as it's a 204, this branch is success
			return ModifyCallResponse{Status: "0"}, VoiceErrorResponse{}, nil
		}
	}

	// this is a backstop, we shouldn't end up here
	return ModifyCallResponse{}, VoiceErrorResponse{}, errors.New("Unsupported combination of parameters, supply an answer URL or valid NCCO")
}

type ModifyCallOpts struct {
	Action string `json:"action"`
}

// Hangup wraps the Modify Call API endpoint
func (client *VoiceClient) Hangup(uuid string) (ModifyCallResponse, VoiceErrorResponse, error) {
	return client.voiceAction("hangup", uuid)
}

// Mute wraps the Modify Call API endpoint
func (client *VoiceClient) Mute(uuid string) (ModifyCallResponse, VoiceErrorResponse, error) {
	return client.voiceAction("mute", uuid)
}

// Unmute wraps the Modify Call API endpoint
func (client *VoiceClient) Unmute(uuid string) (ModifyCallResponse, VoiceErrorResponse, error) {
	return client.voiceAction("unmute", uuid)
}

// Earmuff wraps the Modify Call API endpoint
func (client *VoiceClient) Earmuff(uuid string) (ModifyCallResponse, VoiceErrorResponse, error) {
	return client.voiceAction("earmuff", uuid)
}

// Unearmuff wraps the Modify Call API endpoint
func (client *VoiceClient) Unearmuff(uuid string) (ModifyCallResponse, VoiceErrorResponse, error) {
	return client.voiceAction("unearmuff", uuid)
}

// voiceAction holds the code for the actions that have no extra params
func (client *VoiceClient) voiceAction(action string, uuid string) (ModifyCallResponse, VoiceErrorResponse, error) {
	// create the client
	voiceClient := voice.NewAPIClient(client.Config)
	modifyCallOpts := voice.ModifyCallOpts{Opts: optional.NewInterface(ModifyCallOpts{Action: action})}
	ctx := context.Background()

	response, err := voiceClient.CallsApi.UpdateCall(ctx, uuid, &modifyCallOpts)
	if err != nil {
		e := err.(voice.GenericOpenAPIError)
		data := e.Body()
		errorType := e.Error()
		if errorType == "400 Bad Request" {
			var errResp VoiceErrorInvalidParamsResponse
			json.Unmarshal(data, &errResp)
			return ModifyCallResponse{}, VoiceErrorResponse{Error: errResp}, err
		}
		return ModifyCallResponse{}, VoiceErrorResponse{Error: response}, err
	} else {
		// not a whole lot to return as it's a 204, this branch is success
		return ModifyCallResponse{Status: "0"}, VoiceErrorResponse{}, nil
	}

	// this is a backstop, we shouldn't end up here
	return ModifyCallResponse{}, VoiceErrorResponse{}, errors.New("Unsupported combination of parameters, supply an answer URL or valid NCCO")
}
