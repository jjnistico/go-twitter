package gotwit

type client struct {
	Tweets *Tweets
	Users  *Users
}

func NewClient(apiKey string, apiSecret string, oauthToken string, oauthTokenSecret string) *client {
	initCredentials(apiKey, apiSecret, oauthToken, oauthTokenSecret)
	client := client{}
	return &client
}
