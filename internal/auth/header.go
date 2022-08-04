package auth

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func BuildAuthorizationHeader(
	method string,
	url string,
	queryParams url.Values,
) string {
	oauthConsumerKey := credentials.apiKey
	oauthToken := credentials.oauthToken
	timestamp := unixTime()
	nonce := getNonce(42)

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

	requestSignature := getRequestSignature(
		signaturePayload,
		method,
		url,
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
			fmt.Fprintf(&builder, "%s=\"%s\"", percentEncode(k), percentEncode(v))
			if idx < len(authHeaderData)-1 {
				builder.WriteString(", ")
			}
		}
	}

	return builder.String()
}
