package gotwit

import "gotwitter/internal/auth"

type client struct {
	Tweets *Tweets
	Users  *Users
}

func NewClient(apiKey string, apiSecret string, oauthToken string, oauthTokenSecret string) *client {
	auth.Init(apiKey, apiSecret, oauthToken, oauthTokenSecret)
	client := client{}
	return &client
}
