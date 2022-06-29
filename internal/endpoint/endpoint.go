package endpoint

import "fmt"

const (
	Tweets            string = "https://api.twitter.com/2/tweets"
	Users             string = "https://api.twitter.com/2/users"
	OauthRequestToken string = "https://api.twitter.com/oauth/request_token"
	OauthAuthorize    string = "https://api.twitter.com/oauth/authorize"
	OauthAuthenticate string = "https://api.twitter.com/oauth/authenticate"
	OauthAccessToken  string = "https://api.twitter.com/oauth/access_token"
)

func TimelineTweets(user_id string) string {
	return fmt.Sprintf("%s/%s/tweets", Users, user_id)
}

func UserByUsername(user_name string) string {
	return fmt.Sprintf("%s/by/username/%s", Users, user_name)
}

// see https://developer.twitter.com/en/docs/twitter-api/tweets/quote-tweets/api-reference/get-tweets-id-quote_tweets
func QuoteTweetsByTweetId(tweet_id string) string {
	return fmt.Sprintf("%s/%s/quote_tweets", Tweets, tweet_id)
}

// see https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id#tab0
func TweetById(tweet_id string) string {
	return fmt.Sprintf("%s/%s", Tweets, tweet_id)
}

func FollowersById(user_id string) string {
	return fmt.Sprintf("%s/%s/followers", Users, user_id)
}
