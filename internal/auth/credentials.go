package auth

type config struct {
	apiKey           string
	apiSecret        string
	oauthToken       string
	oauthTokenSecret string
}

// testing default
var credentials *config = &config{"apiKey", "apiSecret", "oauthToken", "oauthTokenSecret"}

func Init(apiKey string, apiSecret string, oauthToken string, oauthTokenSecret string) {
	credentials = &config{
		apiKey:           apiKey,
		apiSecret:        apiSecret,
		oauthToken:       oauthToken,
		oauthTokenSecret: oauthTokenSecret,
	}
}
