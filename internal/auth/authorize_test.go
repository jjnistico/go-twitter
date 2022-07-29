package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

func TestBuildAuthorizationHeader(t *testing.T) {
	authHeader := BuildAuthorizationHeader(http.MethodGet, "www.example.com", url.Values{})

	if authHeader[:5] != "OAuth" {
		t.Error("Missing `OAuth` prefix")
	}

	if !strings.Contains(authHeader, `oauth_consumer_key="apiKey"`) {
		t.Error("missing `oauth_consumer_key` in auth header")
	}

	if !strings.Contains(authHeader, "oauth_nonce") {
		t.Error("missing `oauth_nonce` in auth header")
	}

	if !strings.Contains(authHeader, "oauth_signature") {
		t.Error("missing `oauth_signature` in auth header")
	}

	if !strings.Contains(authHeader, `oauth_signature_method="HMAC-SHA1"`) {
		t.Error("missing `oauth_signature_method` in auth header")
	}

	if !strings.Contains(authHeader, "oauth_timestamp") {
		t.Error("missing `oauth_timestamp` in auth header")
	}

	if !strings.Contains(authHeader, `oauth_token="oauthToken"`) {
		t.Error("missing `oauth_token` in auth header")
	}

	if !strings.Contains(authHeader, `oauth_version="1.0"`) {
		t.Error("missing `oauth_version` in auth header")
	}
}

func TestBuilderAuthorizationHeaderWithParams(t *testing.T) {
	queryParams := url.Values{}
	queryParams.Add("key1", "val1")
	queryParams.Add("key2", "val2")

	authHeader := BuildAuthorizationHeader(http.MethodGet, "www.example.com", queryParams)

	if authHeader != "" {
		t.Error(authHeader)
	}
}
