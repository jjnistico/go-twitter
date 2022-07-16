package gotwit

type config struct {
	apiKey           string
	apiSecret        string
	oauthToken       string
	oauthTokenSecret string
}

var credentialSvc *config

func initCredentials(apiKey string, apiSecret string, oauthToken string, oauthTokenSecret string) {
	credentialSvc = &config{
		apiKey:           apiKey,
		apiSecret:        apiSecret,
		oauthToken:       oauthToken,
		oauthTokenSecret: oauthTokenSecret,
	}
}

func (as *config) ApiKey() string {
	return as.apiKey
}

func (as *config) ApiSecret() string {
	return as.apiSecret
}

func (as *config) OAuthToken() string {
	return as.oauthToken
}

func (as *config) OAuthTokenSecret() string {
	return as.oauthTokenSecret
}
