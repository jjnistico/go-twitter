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

	browser "github.com/pkg/browser"
)

func basicAuthorizationHeader() string {
	clientId := os.Getenv("TWITTER_API_CLIENT_ID")
	clientSecret := os.Getenv("TWITTER_API_CLIENT_SECRET")

	return fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString(
		[]byte(clientId+":"+clientSecret),
	))
}

func authorize() string {
	m := http.NewServeMux()
	server := http.Server{Addr: ":8000", Handler: m}

	cChan := make(chan string, 1)

	m.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")

		if state != "state" {
			log.Fatal("auth callback failed: state != state")
		}

		cChan <- r.URL.Query().Get("code")
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	queryParams := url.Values{}
	queryParams.Set("response_type", "code")
	queryParams.Set("client_id", os.Getenv("TWITTER_API_CLIENT_ID"))
	queryParams.Set("redirect_uri", "http://localhost:8000/callback")
	queryParams.Set("scope", "users.read")
	queryParams.Set("state", "state")
	queryParams.Set("code_challenge", "challenge")
	queryParams.Set("code_challenge_method", "plain")

	url := "https://twitter.com/i/oauth2/authorize?" + queryParams.Encode()
	browser.OpenURL(url)

	code := <-cChan
	server.Shutdown(context.Background())
	return code
}

func accessToken(code string) token {
	form := url.Values{
		"code":       {code},
		"grant_type": {"authorization_code"},
		// "client_id":     {os.Getenv("TWITTER_API_CLIENT_ID")},
		"redirect_uri":  {"http://localhost:8000/callback"},
		"code_verifier": {"challenge"},
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

func Oauth2Authorize(r *http.Request) {
	var aToken string
	if len(credentials.accessToken) == 0 {
		code := authorize()
		token := accessToken(code)
		aToken = token.AccessToken
	} else {
		aToken = credentials.accessToken
	}
	r.Header.Add("Authorization", "Bearer "+aToken)
}

type token struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}
