package auth

import "sync"

type config struct {
	apiKey           string
	apiSecret        string
	oauthToken       string
	oauthTokenSecret string
	clientId         string
	accessToken      token
}

var credentials *config = &config{}

var mx sync.RWMutex

// before each test using credentials
func resetCredentials() {
	credentials = &config{}
}

func InitOAuth2(apiKey string, apiSecret string, clientId string) {
	mx.Lock()
	defer mx.Unlock()

	credentials = &config{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		clientId:  clientId,
	}
}

func InitOAuth1(apiKey string, apiSecret string, oauthToken string, oauthTokenSecret string) {
	mx.Lock()
	defer mx.Unlock()

	credentials = &config{
		apiKey:           apiKey,
		apiSecret:        apiSecret,
		oauthToken:       oauthToken,
		oauthTokenSecret: oauthTokenSecret,
	}
}
