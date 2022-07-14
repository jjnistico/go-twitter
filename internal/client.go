package client

import "gotwitter/internal/api"

type GOTClient struct {
	Tweets *api.Tweets
}

func New( /* api keys/secrets, config, etc */ ) *GOTClient {
	tweets := &api.Tweets{}
	client := GOTClient{Tweets: tweets}
	return &client
}
