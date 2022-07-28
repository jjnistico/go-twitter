package auth

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var (
	nonce            = "HphbgSiKgVIXyAlvLIpAOsRVBpLmPsnskdXtEKzvjx"
	timestamp        = time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC).Unix()
	signaturePayload = []map[string]string{
		{"oauth_consumer_key": credentials.apiKey},
		{"oauth_nonce": nonce},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": credentials.oauthToken},
		{"oauth_version": "1.0"},
	}
)

func TestGetRequestSignatureGet(t *testing.T) {
	const expected = "Jmh6B+ICOqA8VLA97aDQ8phhwb4="

	testSig := getRequestSignature(signaturePayload, http.MethodGet, "www.example.com")

	if testSig != expected {
		t.Errorf("got: %s, expected: %s", testSig, expected)
	}
}

func TestGetRequestSignaturePost(t *testing.T) {
	const expected = "5IGS6mXEFHqkNT++zrO+wKiahFM="

	testSig := getRequestSignature(signaturePayload, http.MethodPost, "www.example.com")

	if testSig != expected {
		t.Errorf("got: %s, expected: %s", testSig, expected)
	}
}

func TestGetRequestSignatureDelete(t *testing.T) {
	const expected = "p6dPMf0E14nlV7KCOQlvBJeUj44="

	testSig := getRequestSignature(signaturePayload, http.MethodDelete, "www.example.com")

	if testSig != expected {
		t.Errorf("got: %s, expected: %s", testSig, expected)
	}
}
