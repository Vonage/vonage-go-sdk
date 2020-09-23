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

func TestApplicationCreate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/v2/applications/",
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
  },
  "keys": {
    "public_key": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCA\nKOxjsU4pf/sMFi9N0jqcSLcjxu33G\nd/vynKnlw9SENi+UZR44GdjGdmfm1\ntL1eA7IBh2HNnkYXnAwYzKJoa4eO3\n0kYWekeIZawIwe/g9faFgkev+1xsO\nOUNhPx2LhuLmgwWSRS4L5W851Xe3f\nUQIDAQAB\n-----END PUBLIC KEY-----\n",
    "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFA\nASCBKcwggSjAgEAAoIBAQDEPpvi+3\nRH1efQ\\nkveWzZDrNNoEXmBw61w+O\n0u/N36tJnN5XnYecU64yHzu2ByEr0\n7iIvYbavFnADwl\\nHMTJwqDQakpa3\n8/SFRnTDq3zronvNZ6nOp7S6K7pcZ\nrw/CvrL6hXT1x7cGBZ4jPx\\nqhjqY\nuJPgZD7OVB69oYOV92vIIJ7JLYwqb\n-----END PRIVATE KEY-----\n"
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
	mesg1 := ApplicationMessages{
		Webhooks: ApplicationMessagesWebhooks{
			StatusUrl:  ApplicationUrl{Address: "https://ljnexmo.eu.ngrok.io/status", HttpMethod: "POST"},
			InboundUrl: ApplicationUrl{Address: "https://ljnexmo.eu.ngrok.io/inbound", HttpMethod: "POST"},
		},
	}

	opts := CreateApplicationOpts{Capabilities: ApplicationCapabilities{Messages: &mesg1}}
	response, _, _ := client.CreateApplication("MyNewApp", opts)

	message := "App Name: " + response.Name
	if message != "App Name: My Application" {
		t.Errorf("Create an application failed")
	}
}

func TestApplicationDelete(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.nexmo.com/v2/applications/abc12345",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			return resp, nil
		},
	)

	auth := CreateAuthFromKeySecret("12345678", "456")
	client := NewApplicationClient(auth)
	response, _, _ := client.DeleteApplication("abc12345")

	if !response {
		t.Errorf("Delete an application failed")
	}
}
