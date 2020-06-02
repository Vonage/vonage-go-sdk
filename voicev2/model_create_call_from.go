/*
 * Voice API BETA
 *
 * This is the *beta* version of the Voice API. Calls created with v2 must be managed using [v1 endpoints](/api/voice).  Voice v2 is provided to allow users to create IP calls. If you do not have this requirement we recommend that you stay on v1 for now.  > This API may break backwards compatibility at short notice (60 days) 
 *
 * API version: 2.1.1
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package voicev2
// CreateCallFrom Connect to a Phone (PSTN) number
type CreateCallFrom struct {
	// The type of connection. Must be `phone`
	Type string `json:"type"`
	// The phone number to connect to
	Number string `json:"number"`
	// Provide [DTMF digits](https://developer.nexmo.com/voice/voice-api/guides/dtmf) to send when the call is answered
	DtmfAnswer string `json:"dtmfAnswer,omitempty"`
}
