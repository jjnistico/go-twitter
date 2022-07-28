package gotwit

import "testing"

func compareResults(t *testing.T, got string, expected string) {
	if got != expected {
		t.Errorf("got: %s, expected: %s", got, expected)
	}
}

func TestTimelineTweetsEndpoint(t *testing.T) {
	expected := "https://api.twitter.com/2/users/abc123/tweets"

	got := timelineTweetsEndpoint("abc123")

	compareResults(t, got, expected)
}

func TestUserByUsernameEndpoint(t *testing.T) {
	expected := "https://api.twitter.com/2/users/by/username/testUser"

	got := userByUsernameEndpoint("testUser")

	compareResults(t, got, expected)
}

func TestQuoteTweetsByTweetIdEndpoint(t *testing.T) {
	expected := "https://api.twitter.com/2/tweets/123456/quote_tweets"

	got := quoteTweetsByTweetIdEndpoint("123456")

	compareResults(t, got, expected)
}

func TestTweetByIdEndpoint(t *testing.T) {
	expected := "https://api.twitter.com/2/tweets/123456"

	got := tweetByIdEndpoint("123456")

	compareResults(t, got, expected)
}

func TestFollowersByIdEndpoint(t *testing.T) {
	expected := "https://api.twitter.com/2/users/abc123/followers"

	got := followersByIdEndpoint("abc123")

	compareResults(t, got, expected)
}
