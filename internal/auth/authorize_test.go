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
	const expected = "Jmh6B+ICOqA8VLA97aDQ8phhwb4="

	testSig := getRequestSignature(signaturePayload, http.MethodGet, "www.example.com")

	compareResults(t, testSig, expected)
}

func TestGetRequestSignaturePost(t *testing.T) {
	const expected = "5IGS6mXEFHqkNT++zrO+wKiahFM="

	testSig := getRequestSignature(signaturePayload, http.MethodPost, "www.example.com")

	compareResults(t, testSig, expected)
}

func TestGetRequestSignatureDelete(t *testing.T) {
	const expected = "p6dPMf0E14nlV7KCOQlvBJeUj44="

	testSig := getRequestSignature(signaturePayload, http.MethodDelete, "www.example.com")

	compareResults(t, testSig, expected)
}

func TestBuildAuthorizationHeader(t *testing.T) {
	const expected = `OAuth oauth_consumer_key="apiKey", oauth_nonce="HphbgSiKgVIXyAlvLIpAOsRVBpLmPsnskdXtEKzvjx", oauth_signature="Jmh6B%2BICOqA8VLA97aDQ8phhwb4%3D", oauth_signature_method="HMAC-SHA1", oauth_timestamp="1654041600", oauth_token="oauthToken", oauth_version="1.0"`
	unixFunc = func() int64 {
		return timestamp
	}

	nonceFunc = func(len int) string {
		return mockNonce
	}

	defer resetClockImpl()
	defer resetNonceImpl()

	authHeader := BuildAuthorizationHeader(http.MethodGet, "www.example.com", url.Values{})

	compareResults(t, authHeader, expected)
}

func TestBuilderAuthorizationHeaderWithParams(t *testing.T) {
	const expected = `OAuth oauth_consumer_key="apiKey", oauth_nonce="HphbgSiKgVIXyAlvLIpAOsRVBpLmPsnskdXtEKzvjx", oauth_signature="4FlfcqvAMEm43N410Dee4yecDQQ%3D", oauth_signature_method="HMAC-SHA1", oauth_timestamp="1654041600", oauth_token="oauthToken", oauth_version="1.0"`
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

	authHeader := BuildAuthorizationHeader(http.MethodGet, "www.example.com", queryParams)

	compareResults(t, authHeader, expected)
}

func compareResults(t *testing.T, got string, expected string) {
	if got != expected {
		t.Errorf("got: %s, expected: %s", got, expected)
	}
}
