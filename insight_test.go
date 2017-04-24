package nexmo

import "testing"

func TestBasicInsight(t *testing.T) {
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
