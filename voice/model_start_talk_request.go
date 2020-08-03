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
// StartTalkRequest struct for StartTalkRequest
type StartTalkRequest struct {
	// The text to read
	Text string `json:"text"`
	VoiceName VoiceName `json:"voice_name,omitempty"`
	// The number of times to repeat the text the file, 0 for infinite
	Loop int32 `json:"loop,omitempty"`
	// The volume level that the speech is played. This can be any value between `-1` to `1` in `0.1` increments, with `0` being the default.
	Level string `json:"level,omitempty"`
}
