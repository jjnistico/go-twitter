package oauth

import (
	"fmt"
	"gotwitter/internal/auth"
	"gotwitter/internal/utils"
	"net/http"
	"strings"
	"time"
)

// see https://developer.twitter.com/en/docs/authentication/oauth-1-0a for oauth 1.0 auth flow used below
func AuthorizeRequest(req *http.Request) {
	oauthConsumerKey := auth.AuthSvc.ApiKey()
	oauthToken := auth.AuthSvc.OAuthToken()
	nonce := generateNonce(14)
	timestamp := time.Now().Unix()

	query_params := req.URL.Query()

	signature_payload := []map[string]string{
		{"oauth_consumer_key": oauthConsumerKey},
		{"oauth_nonce": nonce},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": oauthToken},
		{"oauth_version": "1.0"},
	}

	for key, param := range query_params {
		signature_payload = append(signature_payload, map[string]string{key: strings.Join(param, ",")})
	}

	utils.SortByMapKey(signature_payload)

	req_url := req.URL.Scheme + "://" + req.URL.Host + req.URL.Path

	requestSignature := GetRequestSignature(
		signature_payload,
		req.Method,
		req_url,
		req.URL.RawQuery,
	)

	authorization_header_payload := []map[string]string{
		{"oauth_consumer_key": oauthConsumerKey},
		{"oauth_nonce": nonce},
		{"oauth_signature": requestSignature},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": oauthToken},
		{"oauth_version": "1.0"}}

	authHeader := buildAuthorizationHeader(authorization_header_payload)

	req.Header.Add("Authorization", authHeader)
}

func buildAuthorizationHeader(headerEntries []map[string]string) string {
	builder := strings.Builder{}
	builder.WriteString("OAuth ")
	for idx, entry := range headerEntries {
		for k, v := range entry {
			if v == "" {
				continue
			}
			builder.WriteString(fmt.Sprintf("%s=\"%s\"", percentEncode(k), percentEncode(v)))
			if idx < len(headerEntries)-1 {
				builder.WriteString(", ")
			}
		}
	}

	return builder.String()
}
