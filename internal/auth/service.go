package auth

type Config struct {
	apiKey           string
	apiSecret        string
	oauthToken       string
	oauthTokenSecret string
}

var AuthSvc *Config

func Init(apiKey string, apiSecret string, oauthToken string, oauthTokenSecret string) {
	AuthSvc = &Config{
		apiKey:           apiKey,
		apiSecret:        apiSecret,
		oauthToken:       oauthToken,
		oauthTokenSecret: oauthTokenSecret,
	}
}

func (as *Config) ApiKey() string {
	return as.apiKey
}

func (as *Config) ApiSecret() string {
	return as.apiSecret
}

func (as *Config) OAuthToken() string {
	return as.oauthToken
}

func (as *Config) OAuthTokenSecret() string {
	return as.oauthTokenSecret
}
