package endpoint

import "fmt"

const (
	GetTweets         string = "https://api.twitter.com/2/tweets"
	GetUsers          string = "https://api.twitter.com/2/users"
	OauthRequestToken string = "https://api.twitter.com/oauth/request_token"
	OauthAuthorize    string = "https://api.twitter.com/oauth/authorize"
	OauthAuthenticate string = "https://api.twitter.com/oauth/authenticate"
	OauthAccessToken  string = "https://api.twitter.com/oauth/access_token"
	HomeTimeline      string = "https://api.twitter.com/1.1/statuses/home_timeline.json"
	UserTimeline      string = "https://api.twitter.com/1.1/statuses/user_timeline.json"
)

func TimelineTweets(user_id string) string {
	return fmt.Sprintf("%s/%s/tweets", GetUsers, user_id)
}

func UserByUsername(user_name string) string {
	return fmt.Sprintf("%s/by/username/%s", GetUsers, user_name)
}
