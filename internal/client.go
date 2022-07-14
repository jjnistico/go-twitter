package client

import "gotwitter/internal/api"

type Options = map[string][]string

type Payload = map[string]string

type GOTClient struct {
	Tweets *api.Tweets
	Users  *api.Users
}

func NewClient( /* api keys/secrets, config, etc */ ) *GOTClient {
	tweets := &api.Tweets{}
	client := GOTClient{Tweets: tweets}
	return &client
}
