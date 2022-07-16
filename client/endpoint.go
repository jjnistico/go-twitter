package gotwit

import "fmt"

const (
	tweetsEndpoint            string = "https://api.twitter.com/2/tweets"
	usersEndpoint             string = "https://api.twitter.com/2/users"
	oauthRequestTokenEndpoint string = "https://api.twitter.com/oauth/request_token"
	oauthAuthorizeEndpoint    string = "https://api.twitter.com/oauth/authorize"
	oauthAuthenticateEndpoint string = "https://api.twitter.com/oauth/authenticate"
	oauthAccessTokenEndpoint  string = "https://api.twitter.com/oauth/access_token"
)

func timelineTweetsEndpoint(user_id string) string {
	return fmt.Sprintf("%s/%s/tweets", usersEndpoint, user_id)
}

func userByUsernameEndpoint(user_name string) string {
	return fmt.Sprintf("%s/by/username/%s", usersEndpoint, user_name)
}

// see https://developer.twitter.com/en/docs/twitter-api/tweets/quote-tweets/api-reference/get-tweets-id-quote_tweets
func quoteTweetsByTweetIdEndpoint(tweet_id string) string {
	return fmt.Sprintf("%s/%s/quote_tweets", tweetsEndpoint, tweet_id)
}

// see https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id#tab0
func tweetByIdEndpoint(tweet_id string) string {
	return fmt.Sprintf("%s/%s", tweetsEndpoint, tweet_id)
}

func followersByIdEndpoint(user_id string) string {
	return fmt.Sprintf("%s/%s/followers", usersEndpoint, user_id)
}
