package utils

import (
	"gotwitter/internal/tools/utils/test"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetPathParameterFromQuery(t *testing.T) {
	query_params := url.Values{}
	expected_val := "test_val"
	query_params.Add("test_param", expected_val)
	query_params.Add("remaining_param", "remaining_val")

	mock_req := httptest.NewRequest(http.MethodGet, "localhost:8090", nil)
	mock_req.URL.RawQuery = query_params.Encode()

	mock_writer := httptest.ResponseRecorder{}

	actual_val := GetPathParameterFromQuery(&mock_writer, mock_req, "test_param")

	if actual_val != expected_val {
		t.Errorf("\nexpected: %s, got: %s\n", expected_val, actual_val)
	}

	if mock_req.URL.Query().Get("remaining_param") == "" {
		t.Errorf("\nexpected remaining param key `remaining_param` to be %s\n", "remaining_val")
	}
}

func TestMissingPathParameterFromQuery(t *testing.T) {
	query_params := url.Values{}
	param := "test_param"

	mock_req := httptest.NewRequest(http.MethodGet, "localhost:8090", nil)
	mock_req.URL.RawQuery = query_params.Encode()

	mock_writer := httptest.NewRecorder()

	actual_val := GetPathParameterFromQuery(mock_writer, mock_req, param)

	if actual_val != "" {
		t.Errorf("\nexpected empty string for missing key but got: %s", actual_val)
	}

	if mock_writer.Result().StatusCode != http.StatusBadRequest {
		t.Errorf(
			"\nexpected bad request status code on missing key but got: %d",
			mock_writer.Result().StatusCode,
		)
	}

	expected_json_resp := test.MockResponseFromError(param)

	resp := mock_writer.Result()
	body, _ := io.ReadAll(resp.Body)

	if string(expected_json_resp) != string(body) {
		t.Error("response writer not given correct response json")
	}
}
