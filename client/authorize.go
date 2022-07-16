package gotwit

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"
)

// see https://developer.twitter.com/en/docs/authentication/oauth-1-0a for oauth 1.0 auth flow used below
func authorizeRequest(req *http.Request) {
	oauthConsumerKey := credentialSvc.ApiKey()
	oauthToken := credentialSvc.OAuthToken()
	nonce := generateNonce(14)
	timestamp := time.Now().Unix()
	queryParams := req.URL.Query()

	signaturePayload := []map[string]string{
		{"oauth_consumer_key": oauthConsumerKey},
		{"oauth_nonce": nonce},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": oauthToken},
		{"oauth_version": "1.0"},
	}

	for key, param := range queryParams {
		signaturePayload = append(signaturePayload, map[string]string{key: strings.Join(param, ",")})
	}

	sort.Slice(signaturePayload, func(i, j int) bool {
		var aKey, bKey string
		for key := range signaturePayload[i] {
			aKey = key
		}
		for key := range signaturePayload[j] {
			bKey = key
		}
		return aKey < bKey
	})

	reqUrl := req.URL.Scheme + "://" + req.URL.Host + req.URL.Path

	requestSignature := getRequestSignature(
		signaturePayload,
		req.Method,
		reqUrl,
		req.URL.RawQuery,
	)

	authHeaderData := []map[string]string{
		{"oauth_consumer_key": oauthConsumerKey},
		{"oauth_nonce": nonce},
		{"oauth_signature": requestSignature},
		{"oauth_signature_method": "HMAC-SHA1"},
		{"oauth_timestamp": fmt.Sprintf("%d", timestamp)},
		{"oauth_token": oauthToken},
		{"oauth_version": "1.0"}}

	builder := strings.Builder{}
	builder.WriteString("OAuth ")
	for idx, entry := range authHeaderData {
		for k, v := range entry {
			if v == "" {
				continue
			}
			builder.WriteString(fmt.Sprintf("%s=\"%s\"", percentEncode(k), percentEncode(v)))
			if idx < len(authHeaderData)-1 {
				builder.WriteString(", ")
			}
		}
	}
	authHeader := builder.String()

	req.Header.Add("Authorization", authHeader)
}
