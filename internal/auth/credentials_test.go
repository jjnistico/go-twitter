package auth

import (
	"testing"
)

func TestCredentialsInit(t *testing.T) {
	Init("1", "2", "3", "4")

	if credentials.apiKey != "1" {
		t.Errorf("expected: 1, got: %s", credentials.apiKey)
	}
	if credentials.apiSecret != "2" {
		t.Errorf("expected: 2, got: %s", credentials.apiSecret)
	}
	if credentials.oauthToken != "3" {
		t.Errorf("expected: 3, got: %s", credentials.oauthToken)
	}
	if credentials.oauthTokenSecret != "4" {
		t.Errorf("expected: 4, got: %s", credentials.oauthTokenSecret)
	}
}
