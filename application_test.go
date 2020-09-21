package vonage

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestApplicationNewApplicationClient(*testing.T) {
	auth := CreateAuthFromKeySecret("123", "456")
	NewApplicationClient(auth)
}

func TestApplicationsList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/v2/applications/",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "page_size": 10,
  "page": 1,
  "total_items": 6,
  "total_pages": 1,
  "_embedded": {
    "applications": [
      {
        "id": "78d335fa323d01149c3dd6f0d48968cf",
        "name": "My Application",
        "capabilities": {
          "voice": {
            "webhooks": {
              "answer_url": {
                "address": "https://example.com/webhooks/answer",
                "http_method": "POST"
              },
              "fallback_answer_url": {
                "address": "https://fallback.example.com/webhooks/answer",
                "http_method": "POST"
              },
              "event_url": {
                "address": "https://example.com/webhooks/event",
                "http_method": "POST"
              }
            }
          },
          "messages": {
            "webhooks": {
              "inbound_url": {
                "address": "https://example.com/webhooks/inbound",
                "http_method": "POST"
              },
              "status_url": {
                "address": "https://example.com/webhooks/status",
                "http_method": "POST"
              }
            }
          },
          "rtc": {
            "webhooks": {
              "event_url": {
                "address": "https://example.com/webhooks/event",
                "http_method": "POST"
              }
            }
          },
          "vbc": {}
        }
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

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewApplicationClient(auth)
	response, _, _ := client.GetApplications(GetApplicationsOpts{})

	message := "Total Items: " + strconv.FormatInt(int64(response.TotalItems), 10)
	if message != "Total Items: 6" {
		t.Errorf("List applications failed")
	}
}

func TestApplicationGet(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.nexmo.com/v2/applications/abcd1234",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
{
  "id": "78d335fa323d01149c3dd6f0d48968cf",
  "name": "My Application",
  "capabilities": {
    "voice": {
      "webhooks": {
        "answer_url": {
          "address": "https://example.com/webhooks/answer",
          "http_method": "POST"
        },
        "fallback_answer_url": {
          "address": "https://fallback.example.com/webhooks/answer",
          "http_method": "POST"
        },
        "event_url": {
          "address": "https://example.com/webhooks/event",
          "http_method": "POST"
        }
      }
    },
    "messages": {
      "webhooks": {
        "inbound_url": {
          "address": "https://example.com/webhooks/inbound",
          "http_method": "POST"
        },
        "status_url": {
          "address": "https://example.com/webhooks/status",
          "http_method": "POST"
        }
      }
    },
    "rtc": {
      "webhooks": {
        "event_url": {
          "address": "https://example.com/webhooks/event",
          "http_method": "POST"
        }
      }
    },
    "vbc": {}
  }
}
	`,
			)

			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewApplicationClient(auth)
	response, _, _ := client.GetApplication("abcd1234")

	message := "App Name: " + response.Name
	if message != "App Name: My Application" {
		t.Errorf("Get an application failed")
	}
}
