package auth

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	browser "github.com/pkg/browser"
)

func basicAuthorizationHeader() string {
	clientId := os.Getenv("TWITTER_API_CLIENT_ID")
	clientSecret := os.Getenv("TWITTER_API_CLIENT_SECRET")

	return fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString(
		[]byte(clientId+":"+clientSecret),
	))
}

// authorize is the first step in the oauth 2.0 flow. An http server is spun up to handle the callback
// from twitter and capture the code used to get an access token in the next step.
func authorize() (code string, challenge string) {
	m := http.NewServeMux()
	server := http.Server{Addr: ":8000", Handler: m}

	cChan := make(chan string, 1)
	oauthState := generateNonce(14)
	verifier, codeChallenge := generateCodeVerifier()

	m.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")

		if state != oauthState {
			log.Fatal("auth callback failed: state != " + oauthState)
		}

		cChan <- r.URL.Query().Get("code")
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	url := "https://twitter.com/i/oauth2/authorize?" + url.Values{
		"response_type":         {"code"},
		"client_id":             {credentials.clientId},
		"redirect_uri":          {"http://localhost:8000/callback"},
		"scope":                 {"users.read tweet.read tweet.write"},
		"state":                 {oauthState},
		"code_challenge":        {codeChallenge},
		"code_challenge_method": {"S256"},
	}.Encode()
	browser.OpenURL(url)

	// block until code query parameter value returned
	cCode := <-cChan
	server.Shutdown(context.Background())
	return cCode, verifier
}

func accessToken(code string, verifier string) token {
	form := url.Values{
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {"http://localhost:8000/callback"},
		"code_verifier": {verifier},
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.twitter.com/2/oauth2/token",
		strings.NewReader(form.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", basicAuthorizationHeader())

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)

	var tokenData token
	if err := json.Unmarshal(data, &tokenData); err != nil {
		log.Fatal(err)
	}
	return tokenData
}

var m sync.Mutex

func OAuth2Authorize(r *http.Request) {
	tokenChan := make(chan string)

	go func() {
		m.Lock()
		defer m.Unlock()

		if len(credentials.accessToken.AccessToken) == 0 {
			code, challenge := authorize()
			token := accessToken(code, challenge)
			credentials.accessToken = token
		}

		tokenChan <- credentials.accessToken.AccessToken
	}()

	r.Header.Add("Authorization", "Bearer "+<-tokenChan)
}

type token struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}
