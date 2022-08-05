package gotwit

import "gotwitter/internal/auth"

type client struct {
	TimelineTweets *TimelineTweets
	Tweets         *Tweets
	Users          *Users
}

func NewClient(apiKey string, apiSecret string, clientId string) *client {
	auth.InitOAuth2(apiKey, apiSecret, clientId)
	client := client{}
	return &client
}
