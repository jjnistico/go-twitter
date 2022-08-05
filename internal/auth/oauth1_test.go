package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var (
	mockNonce        = "HphbgSiKgVIXyAlvLIpAOsRVBpLmPsnskdXtEKzvjx"
	timestamp        = time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC).Unix()
	signaturePayload = []map[string]string{
		{"oauth_consumer_key": credentials.apiKey},
		{"oauth_nonce": mockNonce},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": credentials.oauthToken},
		{"oauth_version": "1.0"},
	}
)

func TestGetRequestSignatureGet(t *testing.T) {
	resetCredentials()
	const expected = "KNWMksJ3Rn37Ffy6xVT1f//R15w="

	testSig := getRequestSignature(signaturePayload, http.MethodGet, "www.example.com")

	compareResults(t, testSig, expected)
}

func TestGetRequestSignaturePost(t *testing.T) {
	resetCredentials()
	const expected = "yaKt59q6rV7Fjar2zLdDtzD1lV8="

	testSig := getRequestSignature(signaturePayload, http.MethodPost, "www.example.com")

	compareResults(t, testSig, expected)
}

func TestGetRequestSignatureDelete(t *testing.T) {
	resetCredentials()
	const expected = "rOx34j9xOvR+4i0f0EDwxNm0umg="

	testSig := getRequestSignature(signaturePayload, http.MethodDelete, "www.example.com")

	compareResults(t, testSig, expected)
}

func TestBuildAuthorizationHeader(t *testing.T) {
	resetCredentials()
	const expected = `OAuth oauth_nonce="HphbgSiKgVIXyAlvLIpAOsRVBpLmPsnskdXtEKzvjx", oauth_signature="KNWMksJ3Rn37Ffy6xVT1f%2F%2FR15w%3D", oauth_signature_method="HMAC-SHA1", oauth_timestamp="1654041600", oauth_version="1.0"`
	unixFunc = func() int64 {
		return timestamp
	}

	nonceFunc = func(len int) string {
		return mockNonce
	}

	defer resetClockImpl()
	defer resetNonceImpl()

	authHeader := buildAuthorizationHeader(http.MethodGet, "www.example.com", url.Values{})

	compareResults(t, authHeader, expected)
}

func TestBuilderAuthorizationHeaderWithParams(t *testing.T) {
	resetCredentials()
	const expected = `OAuth oauth_nonce="HphbgSiKgVIXyAlvLIpAOsRVBpLmPsnskdXtEKzvjx", oauth_signature="MNC1IJ2yMRzR5LKvMMg0MGa9q18%3D", oauth_signature_method="HMAC-SHA1", oauth_timestamp="1654041600", oauth_version="1.0"`
	queryParams := url.Values{}
	queryParams.Add("key1", "val1")
	queryParams.Add("key2", "val2")

	unixFunc = func() int64 {
		return timestamp
	}

	nonceFunc = func(len int) string {
		return mockNonce
	}

	defer resetClockImpl()
	defer resetNonceImpl()

	authHeader := buildAuthorizationHeader(http.MethodGet, "www.example.com", queryParams)

	compareResults(t, authHeader, expected)
}

func compareResults(t *testing.T, got string, expected string) {
	if got != expected {
		t.Errorf("got: %s, expected: %s", got, expected)
	}
}
