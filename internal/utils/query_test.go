package utils

import (
	"net/url"
	"testing"
)

func TestGetPathParameterFromQuery(t *testing.T) {
	queryParams := url.Values{}
	expectedVal := "test_val"
	queryParams.Add("test_param", expectedVal)
	queryParams.Add("remaining_param", "remaining_val")

	actualVal, newParams, err := ExtractParameterFromQuery(queryParams, "test_param")

	if err != nil {
		t.Errorf("\nerror getting query param: %s\n", err.Error())
	}

	if actualVal != expectedVal {
		t.Errorf("\nexpected: %s, got: %s\n", expectedVal, actualVal)
	}

	if newParams.Get("remaining_param") == "" {
		t.Errorf("\nexpected remaining param key `remaining_param` to be %s\n", "remaining_val")
	}
}

func TestMissingPathParameterFromQuery(t *testing.T) {
	queryParams := url.Values{}
	param := "test_param"

	actualVal, newParams, err := ExtractParameterFromQuery(queryParams, param)

	if err == nil {
		t.Error("expected error with missing query param but got nil")
	}

	if newParams.Encode() != "" {
		t.Error("expected empty query params")
	}

	if actualVal != "" {
		t.Errorf("\nexpected empty string for missing key but got: %s", actualVal)
	}
}
