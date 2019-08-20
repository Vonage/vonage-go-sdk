package nexmo

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestBasicInsight(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/ni/basic/json",
		httpmock.NewStringResponder(200, `{
				"status": 0,
				"status_message": "Success",
				"request_id": "261050f4-5126-43ab-b6cf-7cb2ce341c8b",
				"international_format_number": "447520615146",
				"national_format_number": "07520 615146",
				"country_code": "GB",
				"country_code_iso3": "GBR",
				"country_name": "United Kingdom",
				"country_prefix": "44"
			}`))

	ar, _, err := _client.Insight.GetBasicInsight(BasicInsightRequest{
		Number: "447520615146",
	})
	if err != nil {
		t.Error(err)
	}
	if ar.CountryCode != "GB" {
		t.Errorf("Country code should have been \"GB\", instead was %v", ar.CountryCode)
	}
}

func TestStandardInsight(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/ni/standard/json",
		httpmock.NewStringResponder(200, `{
				"status": 0,
				"status_message": "Success",
				"request_id": "e983e5c2-03d3-4432-8487-b14834eda8c9",
				"international_format_number": "447520615146",
				"national_format_number": "07520 615146",
				"country_code": "GB",
				"country_code_iso3": "GBR",
				"country_name": "United Kingdom",
				"country_prefix": "44",
				"request_price": "0.00500000",
				"remaining_balance": "54.49271782",
				"current_carrier": {
					"network_code": "23409",
					"name": "Tismi BV",
					"country": "GB",
					"network_type": "mobile"
				},
				"original_carrier": {
					"network_code": "23409",
					"name": "Tismi BV",
					"country": "GB",
					"network_type": "mobile"
				},
				"ported": "assumed_not_ported",
				"roaming": {"status": "unknown"}
			}`))

	ar, _, err := _client.Insight.GetStandardInsight(StandardInsightRequest{
		Number: "447520615146",
		CNAM:   false,
	})
	if err != nil {
		t.Error(err)
	}
	if ar.CountryCode != "GB" {
		t.Errorf("Country code should have been \"GB\", instead was %v", ar.CountryCode)
	}
}

func TestAdvancedInsight(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.nexmo.com/ni/advanced/json",
		httpmock.NewStringResponder(200, `{
				"status": 0,
				"status_message": "Success",
				"lookup_outcome": 0,
				"lookup_outcome_message": "Success",
				"request_id": "62242e49-a824-42aa-8d04-3772eb9ae315",
				"international_format_number": "447520615146",
				"national_format_number": "07520 615146",
				"country_code": "GB",
				"country_code_iso3": "GBR",
				"country_name": "United Kingdom",
				"country_prefix": "44",
				"request_price": "0.03000000",
				"remaining_balance": "54.46271782",
				"current_carrier": {
					"network_code": "23409",
					"name": "Tismi BV",
					"country": "GB",
					"network_type": "mobile"
				},
				"original_carrier": {
					"network_code": "23409",
					"name": "Tismi BV",
					"country": "GB",
					"network_type": "mobile"
				},
				"valid_number": "valid",
				"reachable": "reachable",
				"ported": "not_ported",
				"roaming": {"status": "not_roaming"},
				"ip_warnings": "unknown"
			}`))

	ar, _, err := _client.Insight.GetAdvancedInsight(AdvancedInsightRequest{
		Number: "447520615146",
		CNAM:   false,
	})
	if err != nil {
		t.Error(err)
	}
	if ar.CountryCode != "GB" {
		t.Errorf("Country code should have been \"GB\", instead was %v", ar.CountryCode)
	}
}

func ExampleInsightService_GetBasicInsight() {
	insightResponse, _, err := _client.Insight.GetBasicInsight(BasicInsightRequest{
		Number: "447520615146",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("The number is from: %s\n", insightResponse.CountryCode)
	// Output: The number is from: GB
}
