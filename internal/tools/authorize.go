package tools

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func AuthorizeRequest(req *http.Request, query_params url.Values) {
	oauth_consumer_key := os.Getenv("API_KEY")
	oauth_token := os.Getenv("ACCESS_TOKEN")
	nonce := GenerateNonce(42)
	timestamp := time.Now().Unix()

	signature_payload := []map[string]string{
		{"oauth_consumer_key": oauth_consumer_key},
		{"oauth_nonce": nonce},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_version": "1.0"}}

	req_url := req.URL.Scheme + "://" + req.URL.Host + req.URL.Path

	request_signature := GetRequestSignature(
		signature_payload,
		req.Method,
		req_url,
		query_params,
	)

	authorization_header_payload := []map[string]string{
		{"oauth_consumer_key": oauth_consumer_key},
		{"oauth_nonce": nonce},
		{"oauth_signature": request_signature},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": oauth_token},
		{"oauth_version": "1.0"}}

	auth_header := BuildAuthorizationHeader(authorization_header_payload)

	req.Header.Add("Authorization", auth_header)
}
