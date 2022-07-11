package utils

import (
	"net/url"
	"testing"
)

func TestGetPathParameterFromQuery(t *testing.T) {
	query_params := url.Values{}
	expected_val := "test_val"
	query_params.Add("test_param", expected_val)
	query_params.Add("remaining_param", "remaining_val")

	actual_val, new_params, err := ExtractParameterFromQuery(query_params, "test_param")

	if err != nil {
		t.Errorf("\nerror getting query param: %s\n", err.Error())
	}

	if actual_val != expected_val {
		t.Errorf("\nexpected: %s, got: %s\n", expected_val, actual_val)
	}

	if new_params.Get("remaining_param") == "" {
		t.Errorf("\nexpected remaining param key `remaining_param` to be %s\n", "remaining_val")
	}
}

func TestMissingPathParameterFromQuery(t *testing.T) {
	query_params := url.Values{}
	param := "test_param"

	actual_val, new_params, err := ExtractParameterFromQuery(query_params, param)

	if err == nil {
		t.Error("expected error with missing query param but got nil")
	}

	if new_params.Encode() != "" {
		t.Error("expected empty query params")
	}

	if actual_val != "" {
		t.Errorf("\nexpected empty string for missing key but got: %s", actual_val)
	}
}
