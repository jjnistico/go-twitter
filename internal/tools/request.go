package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func RequestData(endpoint string, query_params string, method string, payload io.Reader) ([]byte, error) {
	base_url := os.Getenv("BASE_URL")
	bearer_token := os.Getenv("BEARER_TOKEN")

	req, err := http.NewRequest(method, base_url+endpoint+"?"+url.QueryEscape(query_params), payload)

	if err != nil {
		fmt.Printf("Error encountered %s %s - %s\n", method, endpoint, err)
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer_token))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error encountered %s %s - %s\n", method, endpoint, err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading response %s\n", err)
		return nil, err
	}

	return body, nil
}

func AuthorizeRequest(oauth_token string, signature string) string {
	header_keys := []string{
		"oauth_consumer_key",
		"oauth_nonce",
		"oauth_signature",
		"oauth_signature_method",
		"oauth_timestamp",
		"oauth_token",
		"oauth_version",
	}

	header_params := map[string]string{
		"oauth_consumer_key":     os.Getenv("API_KEY"),
		"oauth_nonce":            GenerateNonce(42),
		"oauth_signature":        signature,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        fmt.Sprint(time.Now().Unix()),
		"oauth_token":            url.QueryEscape(oauth_token),
		"oauth_version":          "1.0",
	}

	// header string is fields joined by ", ". All key/values are percent encoded
	header_string := "OAuth "
	for idx, key := range header_keys {
		header_string += url.QueryEscape(key)
		header_string += "="
		header_string += url.QueryEscape(header_params[key])
		if idx < len(header_keys)-1 {
			header_string += ", "
		}
	}

	return header_string
}
