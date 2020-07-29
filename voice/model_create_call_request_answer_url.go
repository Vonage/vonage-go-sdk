/*
 * Voice API
 *
 * The Voice API lets you create outbound calls, control in-progress calls and get information about historical calls. More information about the Voice API can be found at <https://developer.nexmo.com/voice/voice-api/overview>.
 *
 * API version: 1.3.2
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package voice
// CreateCallRequestAnswerUrl struct for CreateCallRequestAnswerUrl
type CreateCallRequestAnswerUrl struct {
	// The webhook endpoint where you provide the [Nexmo Call Control Object](/voice/voice-api/ncco-reference) that governs this call. 
	AnswerUrl []string `json:"answer_url"`
	// The HTTP method used to send event information to answer_url.
	AnswerMethod string `json:"answer_method,omitempty"`
	To []OneOfEndpointPhoneToEndpointSipEndpointWebsocketEndpointVbcExtension `json:"to"`
	From EndpointPhoneFrom `json:"from"`
	// **Required** unless `event_url` is configured at the application level, see [Create an Application](/api/application.v2#createApplication)  The webhook endpoint where call progress events are sent to. For more information about the values sent, see [Event webhook](/voice/voice-api/webhook-reference#event-webhook). 
	EventUrl []string `json:"event_url,omitempty"`
	// The HTTP method used to send event information to event_url.
	EventMethod string `json:"event_method,omitempty"`
	// Configure the behavior when Nexmo detects that the call is answered by voicemail. If Continue Nexmo sends an HTTP request to event_url with the Call event machine. hangup  end the call
	MachineDetection string `json:"machine_detection,omitempty"`
	// Set the number of seconds that elapse before Nexmo hangs up after the call state changes to answered.
	LengthTimer int32 `json:"length_timer,omitempty"`
	// Set the number of seconds that elapse before Nexmo hangs up after the call state changes to ‘ringing’.
	RingingTimer int32 `json:"ringing_timer,omitempty"`
}