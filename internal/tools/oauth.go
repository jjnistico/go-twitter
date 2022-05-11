package tools

import (
	"errors"
	"fmt"
	"gotwitter/internal/endpoint"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// see https://developer.twitter.com/en/docs/authentication/api-reference/request_token
func RequestToken() (RequestTokenResponse, error) {
	auth_callback := os.Getenv("AUTH_CALLBACK")

	query_params := url.Values{}
	query_params.Set("oauth_callback", auth_callback)

	token_req, err := http.NewRequest(http.MethodPost, endpoint.OauthRequestToken+"?"+query_params.Encode(), nil)

	if err != nil {
		return RequestTokenResponse{}, fmt.Errorf("error generating request to oauth/request_token: %s", err.Error())
	}

	AuthorizeRequest(token_req, query_params)
	fmt.Printf("%#v", token_req)
	client := &http.Client{}
	resp, err := client.Do(token_req)

	if resp.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
		return RequestTokenResponse{}, fmt.Errorf("error requesting request token: %s", resp.Status)
	}

	if err != nil {
		return RequestTokenResponse{}, fmt.Errorf("error requesting request token: %s", err.Error())
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	token_struct := ParseRequestTokenStringToStruct(string(data))

	if !token_struct.CallbackConfirmed {
		return RequestTokenResponse{}, errors.New("callback not confirmed")
	}

	return token_struct, nil
}

// Authorize is an unauthorized request. It returns the twitter page for login
func Authorize(oauth_token string) ([]byte, error) {
	query_params := url.Values{}
	query_params.Set("oauth_token", oauth_token)

	auth_req, err := http.NewRequest(http.MethodGet, endpoint.OauthAuthorize+"?"+query_params.Encode(), nil)

	if err != nil {
		return nil, fmt.Errorf("error generating request to oauth/authorize: %s", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(auth_req)

	if err != nil {
		return nil, fmt.Errorf("error requesting second leg of oauth protocol: %s", err.Error())
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	return data, nil
}

// The oauth_token here is the same oauth_token returned from `RequestToken` from step 1 of oauth flow
func AccessToken(oauth_token string, oauth_verifier string) (AccessTokenResponse, error) {
	query_params := url.Values{}
	query_params.Set("oauth_token", oauth_token)
	query_params.Set("oauth_verifier", oauth_verifier)

	access_token_req, err := http.NewRequest(http.MethodPost, endpoint.OauthAccessToken+"?"+query_params.Encode(), nil)

	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("error creating request for oauth/access_token: %s", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(access_token_req)

	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("error requesting third leg of oauth protocol: %s", err.Error())
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	token_struct := ParseAccessTokenStringToStruct(string(data))

	return token_struct, nil
}
