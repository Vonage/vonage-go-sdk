package nexmo

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"encoding/json"

	"flag"

	"github.com/judy2k/go-vcr/cassette"
	"github.com/judy2k/go-vcr/recorder"
)

var _recorder *recorder.Recorder
var _client *Client

func TestMain(m *testing.M) {
	var mode string
	flag.StringVar(&mode, "mode", "", "Should be one of \"replaying\", \"recording\" or \"disabled\"")
	flag.Parse()
	fmt.Println("Mode:", mode)
	recorderMode := recorder.ModeReplaying
	switch mode {
	case "replaying":
		fmt.Println("Replaying requests")
		recorderMode = recorder.ModeReplaying
	case "recording":
		fmt.Println("Recording requests")
		recorderMode = recorder.ModeRecording
	case "disabled":
		fmt.Println("HTTP recorder disabled")
		recorderMode = recorder.ModeDisabled
	}

	fmt.Println("Running TestMain")
	os.Exit(func() int {
		// Start our recorder
		r, err := recorder.NewAsMode("fixtures", recorderMode, nil)
		r.SetFilter(filterFunc)
		defer r.Stop()
		if err != nil {
			log.Fatal(err)
		}
		_recorder = r
		_client = initClient()

		return m.Run()
	}())
}

func initClient() *Client {
	apiKey := os.Getenv("NEXMO_API_KEY")
	apiSecret := os.Getenv("NEXMO_API_SECRET")

	if _client != nil {
		return _client
	}

	// TODO: FIX ME!
	// path := os.Getenv("NEXMO_PRIVATE_KEY_PATH")
	// b, err := ioutil.ReadFile(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	auth := NewAuthSet()
	// if err := auth.SetApplicationAuth(os.Getenv("NEXMO_APPLICATION_ID"), b); err != nil {
	// 	log.Fatal(err)
	// }
	auth.SetAPISecret(apiKey, apiSecret)
	httpClient := http.Client{
		Transport: _recorder,
	}
	_client = NewClient(&httpClient, auth)
	return _client
}

func filterFunc(i *cassette.Interaction) error {
	// Purge the headers:
	i.Request.Headers.Del("Authorization")

	// Purge the query:
	qURL, err := url.Parse(i.Request.URL)
	if err != nil {
		return err
	}
	query := qURL.Query()
	query.Del("api_key")
	query.Del("api_secret")

	// Purge the body:
	if ct := i.Request.Headers.Get("Content-Type"); ct == "application/json" {
		data := map[string]interface{}{}
		err := json.Unmarshal([]byte(i.Request.Body), &data)
		if err != nil {
			return err
		}
		delete(data, "api_key")
		delete(data, "api_secret")
		bodyBytes, err := json.Marshal(&data)
		if err != nil {
			return err
		}
		i.Request.Body = string(bodyBytes)
	} else if ct == "application/x-www-form-urlencoded" {
		i.Request.Form.Del("api_key")
		i.Request.Form.Del("api_secret")
		urlValues, err := url.ParseQuery(i.Request.Body)
		if err != nil {
			return err
		}
		urlValues.Del("api_key")
		urlValues.Del("api_secret")
		i.Request.Body = urlValues.Encode()
	}

	return nil
}
