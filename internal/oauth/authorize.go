package oauth

import (
	"fmt"
	"gotwitter/internal/utils"
	"net/http"
	"os"
	"strings"
	"time"
)

// see https://developer.twitter.com/en/docs/authentication/oauth-1-0a for oauth 1.0 auth flow used below
func AuthorizeRequest(req *http.Request) {
	oauth_consumer_key := os.Getenv("TWITTER_API_KEY")
	oauth_token := os.Getenv("TWITTER_OAUTH_TOKEN")
	nonce := generateNonce(14)
	timestamp := time.Now().Unix()

	query_params := req.URL.Query()

	signature_payload := []map[string]string{
		{"oauth_consumer_key": oauth_consumer_key},
		{"oauth_nonce": nonce},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": oauth_token},
		{"oauth_version": "1.0"},
	}

	for key, param := range query_params {
		signature_payload = append(signature_payload, map[string]string{key: strings.Join(param, ",")})
	}

	utils.SortByMapKey(signature_payload)

	req_url := req.URL.Scheme + "://" + req.URL.Host + req.URL.Path

	request_signature := GetRequestSignature(
		signature_payload,
		req.Method,
		req_url,
		req.URL.RawQuery,
	)

	authorization_header_payload := []map[string]string{
		{"oauth_consumer_key": oauth_consumer_key},
		{"oauth_nonce": nonce},
		{"oauth_signature": request_signature},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": oauth_token},
		{"oauth_version": "1.0"}}

	auth_header := buildAuthorizationHeader(authorization_header_payload)

	req.Header.Add("Authorization", auth_header)
}

func buildAuthorizationHeader(header_entries []map[string]string) string {
	builder := strings.Builder{}
	builder.WriteString("OAuth ")
	for idx, entry := range header_entries {
		for k, v := range entry {
			if v == "" {
				continue
			}
			builder.WriteString(fmt.Sprintf("%s=\"%s\"", percentEncode(k), percentEncode(v)))
			if idx < len(header_entries)-1 {
				builder.WriteString(", ")
			}
		}
	}

	return builder.String()
}
