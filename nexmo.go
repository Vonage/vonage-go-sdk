// Simple client for using Nexmo's communication APIs. See https://nexmo.com for more information about the APIs.
package nexmo

import (
	"fmt"
	"net/http"

	"github.com/nexmo-community/nexmo-go/sling"
)

// The main client object
type Client struct {
	sling       *sling.Sling
	Insight     *InsightService
	SMS         *SMSService
	Call        *CallService
	Verify      *VerifyService
	Developer   *DeveloperService
	Application *ApplicationService
}

// Get a new Client object with the auth configured
func NewClient(httpClient *http.Client, authSet *AuthSet) *Client {
	base := sling.New().
		Client(httpClient).
		Set("User-Agent", "nexmo-go/2.0 (nexmo-community)")
	return &Client{
		sling:       base,
		Insight:     newInsightService(base.New(), authSet),
		SMS:         newSMSService(base.New(), authSet),
		Call:        newCallService(base.New(), authSet),
		Verify:      newVerifyService(base.New(), authSet),
		Developer:   newDeveloperService(base.New(), authSet),
		Application: newApplicationService(base.New(), authSet),
	}
}

type APIError struct {
	Status       int64
	ErrorMessage string
}

func (a APIError) Error() string {
	return fmt.Sprintf("%d: %s", a.Status, a.ErrorMessage)
}

// func (s *Sling) DoWithPrejudice(req *http.Request, successV, errorV interface{}) (*http.Response, error) {
// 	resp, err := s.Do(req, successV, nil)
// 	if err != nil {
// 		return resp, err
// 	}
// 	// when err is nil, resp contains a non-nil resp.Body which must be closed
// 	defer resp.Body.Close()

// 	var buf bytes.Buffer
// 	body := io.TeeReader(resp.Body, &buf)

// 	// Don't try to decode on 204s
// 	if resp.StatusCode == 204 {
// 		return resp, nil
// 	}

// 	if successV != nil || errorV != nil {
// 		err = decodeResponseJSON(resp.StatusCode, body, successV, errorV)
// 	}

// 	// This is where we will put our prejudice:
// 	if err != nil {
// 		log.Println("Could not parse response:", buf)
// 	}

// 	return resp, err
// }

// // decodeResponse decodes response Body into the value pointed to by successV
// // if the response is a success (2XX) or into the value pointed to by failureV
// // otherwise. If the successV or failureV argument to decode into is nil,
// // decoding is skipped.
// // Caller is responsible for closing the resp.Body.
// func decodeResponseJSON(statusCode int, body io.Reader, successV, failureV interface{}) error {
// 	if 200 <= statusCode && statusCode <= 299 {
// 		if successV != nil {
// 			return decodeResponseBodyJSON(body, successV)
// 		}
// 	} else {
// 		if failureV != nil {
// 			return decodeResponseBodyJSON(body, failureV)
// 		}
// 	}
// 	return nil
// }

// // decodeResponseBodyJSON JSON decodes a Response Body into the value pointed
// // to by v.
// // Caller must provide a non-nil v and close the resp.Body.
// func decodeResponseBodyJSON(body io.Reader, v interface{}) error {
// 	return json.NewDecoder(body).Decode(v)
// }
