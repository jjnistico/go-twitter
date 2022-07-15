package client

import (
	"gotwitter/internal/api"
	"gotwitter/internal/auth"
)

type Config struct {
	ApiKey           string
	ApiSecret        string
	OAuthToken       string
	OAuthTokenSecret string
}

type Options = map[string][]string

type Payload = map[string]string

type Client struct {
	Tweets *api.Tweets
	Users  *api.Users
}

func NewClient(apiKey string, apiSecret string, oauthToken string, oauthTokenSecret string) *Client {
	auth.Init(apiKey, apiSecret, oauthToken, oauthTokenSecret)
	client := Client{}
	return &client
}
