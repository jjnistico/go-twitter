package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func RequestToken(w http.ResponseWriter) (TokenResponse, error) {
	base_url := os.Getenv("BASE_URL")
	api_key := os.Getenv("API_KEY")
	bearer_token := os.Getenv("BEARER_TOKEN")
	auth_callback := os.Getenv("APP_CALLBACK")

	// URL encode callback
	encodedAuthCallback := url.PathEscape(auth_callback)

	// Initial post request requires encoded url for auth callback and api key
	query_params := url.Values{}
	query_params.Set("oauth_callback", encodedAuthCallback)
	query_params.Add("oauth_consumer_key", api_key)

	token_req, err := http.NewRequest(http.MethodPost, base_url+"oauth/request_token?"+query_params.Encode(), nil)

	if err != nil {
		fmt.Printf("Error requesting token: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating request to oauth/request_token"))
		return TokenResponse{}, err
	}

	token_req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer_token))

	///////////
	client := &http.Client{}
	resp, err := client.Do(token_req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error requesting first leg of oauth protocol"))
		return TokenResponse{}, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	token_struct := ParseStringToStruct(string(body))

	if !token_struct.CallbackConfirmed {
		return TokenResponse{}, errors.New("callback not confirmed")
	}

	return token_struct, nil
}

func Authorize(w http.ResponseWriter, token_response TokenResponse) {
	base_url := os.Getenv("BASE_URL")
	bearer_token := os.Getenv("BEARER_TOKEN")

	query_params := url.Values{}
	query_params.Set("oauth_token", token_response.Token)
	// query_params.Set("force_login", "true")

	auth_req, err := http.NewRequest(http.MethodGet, base_url+"oauth/authorize?"+query_params.Encode(), nil)

	if err != nil {
		fmt.Printf("Error authorizing with request token: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating request to oauth/authorize"))
		return
	}

	auth_req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer_token))

	client := &http.Client{}
	resp, err := client.Do(auth_req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error requesting first leg of oauth protocol"))
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
