package auth

import (
	"testing"
)

func TestCredentialsInit(t *testing.T) {
	Init("1", "2", "3", "4")

	compareResults(t, credentials.apiKey, "1")
	compareResults(t, credentials.apiSecret, "2")
	compareResults(t, credentials.oauthToken, "3")
	compareResults(t, credentials.oauthTokenSecret, "4")
}
