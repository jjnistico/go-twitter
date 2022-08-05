package auth

import (
	"testing"
)

func TestCredentialsInitOauth1(t *testing.T) {
	resetCredentials()
	InitOAuth1("1", "2", "3", "4")

	compareResults(t, credentials.apiKey, "1")
	compareResults(t, credentials.apiSecret, "2")
	compareResults(t, credentials.oauthToken, "3")
	compareResults(t, credentials.oauthTokenSecret, "4")
	compareResults(t, credentials.clientId, "")
}

func TestCredentialsInitOauth2(t *testing.T) {
	resetCredentials()
	InitOAuth2("1", "2", "3")

	compareResults(t, credentials.apiKey, "1")
	compareResults(t, credentials.apiSecret, "2")
	compareResults(t, credentials.clientId, "3")
	compareResults(t, credentials.oauthToken, "")
	compareResults(t, credentials.oauthTokenSecret, "")
}
