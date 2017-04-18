package nexmo

import (
	"net/http"

	"fmt"
)

func (r *BasicInsightResponse) responseError() error {
	if r.Status != 0 {
		if r.StatusMessage != "" {
			return fmt.Errorf("%d: %s", r.Status, r.StatusMessage)
		}
		return fmt.Errorf("%d: %s", r.Status, r.ErrorText)
	}
	return nil
}

func (c *InsightService) GetBasicInsight(request BasicInsightRequest) (BasicInsightResponse, *http.Response, error) {
	c.authSet.ApplyAPICredentials(&request)

	insightResponse := new(BasicInsightResponse)
	resp, err := c.sling.New().Post("basic/json").BodyJSON(request).ReceiveSuccess(insightResponse)
	if err != nil {
		return *insightResponse, resp, err
	}
	err = insightResponse.responseError()
	return *insightResponse, resp, err
}
