package gotwit

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type requestTokenResponse struct {
	Token             string `json:"oauth_token"`
	Secret            string `json:"oauth_token_secret"`
	CallbackConfirmed bool   `json:"oauth_callback_confirmed"`
}

type accessTokenResponse struct {
	ScreenName       string `json:"screenName"`
	UserId           string `json:"userId"`
	OAuthToken       string `json:"oauthToken"`
	OAuthTokenSecret string `json:"oauthTokenSecret"`
}

// see https://developer.twitter.com/en/docs/authentication/api-reference/request_token
func requestToken() (requestTokenResponse, int, error) {
	authCallback := os.Getenv("AUTH_CALLBACK")
	queryParams := url.Values{}
	queryParams.Set("oauth_callback", authCallback)

	tokenReq, err := http.NewRequest(http.MethodPost, oauthRequestTokenEndpoint+"?"+queryParams.Encode(), nil)

	if err != nil {
		return requestTokenResponse{}, http.StatusInternalServerError, fmt.Errorf(
			"error generating request to oauth/request_token: %s",
			err.Error(),
		)
	}

	authorizeRequest(tokenReq)

	client := &http.Client{}
	resp, err := client.Do(tokenReq)

	if resp.StatusCode != http.StatusOK {
		return requestTokenResponse{}, resp.StatusCode, fmt.Errorf("error requesting request token: %s", resp.Status)
	}

	if err != nil {
		return requestTokenResponse{}, resp.StatusCode, fmt.Errorf("error requesting request token: %s", err.Error())
	}

	defer resp.Body.Close()
	// data, _ := ioutil.ReadAll(resp.Body)

	tokenStruct := requestTokenResponse{} // parseRequestTokenStringToStruct(string(data))

	if !tokenStruct.CallbackConfirmed {
		return requestTokenResponse{}, http.StatusUnauthorized, errors.New("callback not confirmed")
	}

	return tokenStruct, resp.StatusCode, nil
}

// authenticate is an unauthorized request. It returns the twitter page for login
func authenticate(oauth_token string) ([]byte, int, error) {
	queryParams := url.Values{}
	queryParams.Set("oauth_token", oauth_token)

	authReq, err := http.NewRequest(http.MethodGet, oauthAuthenticateEndpoint+"?"+queryParams.Encode(), nil)

	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(
			"error generating request to oauth/authorize: %s",
			err.Error(),
		)
	}

	client := &http.Client{}
	resp, err := client.Do(authReq)

	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("error requesting second leg of oauth protocol: %s", err.Error())
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	return data, resp.StatusCode, nil
}

// The oauth_token here is the same oauth_token returned from `RequestToken` from step 1 of oauth flow
func accessToken(oauthToken string, oauthVerifier string) (accessTokenResponse, int, error) {
	queryParams := url.Values{}
	queryParams.Set("oauth_token", oauthToken)
	queryParams.Set("oauth_verifier", oauthVerifier)

	accessTokenReq, err := http.NewRequest(http.MethodPost, oauthAccessTokenEndpoint+"?"+queryParams.Encode(), nil)

	if err != nil {
		return accessTokenResponse{}, http.StatusInternalServerError, fmt.Errorf(
			"error creating request for oauth/access_token: %s",
			err.Error(),
		)
	}

	client := &http.Client{}
	resp, err := client.Do(accessTokenReq)

	if err != nil {
		return accessTokenResponse{}, resp.StatusCode, fmt.Errorf(
			"error requesting third leg of oauth protocol: %s",
			err.Error(),
		)
	}

	defer resp.Body.Close()
	// data, _ := ioutil.ReadAll(resp.Body)

	tokenStruct := accessTokenResponse{} // parseAccessTokenStringToStruct(string(data))

	return tokenStruct, resp.StatusCode, nil
}
