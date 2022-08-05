package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// see https://developer.twitter.com/en/docs/authentication/oauth-1-0a/authorizing-a-request for algorithm
func buildAuthorizationHeader(
	method string,
	url string,
	queryParams url.Values,
) string {
	mx.RLock()
	defer mx.RUnlock()

	oauthConsumerKey := credentials.apiKey
	oauthToken := credentials.oauthToken
	timestamp := unixTimestamp()
	nonce := nonce(42)

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

func OAuth1Authorize(r *http.Request) {
	reqUrl := r.URL.Scheme + "://" + r.URL.Host + r.URL.Path
	authHeader := buildAuthorizationHeader(r.Method, reqUrl, r.URL.Query())
	r.Header.Add("Authorization", authHeader)
}

// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //
// SignRequest generates a parameter string that can be included                             //
// with a request to get authorized for the twitter api.                                     //
// see: https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature //
//                                                                                           //
// Required parts of the param string (must be in alphabetical order):                       //
// `oauth_consumer_key`: string - the applications api_key                                     //
// `oauth_nonce`: string - unique string generated for every request                           //
// `oauth_signature_method`: string - HMAC-SHA1                                                //
// `oauth_timestamp`: string - unix timestamp                                                  //
// `oauth_token`: string - token returned from /authentication                                 //
// `oauth_version`: string - should always be 1.0 for the Twitter API                          //
// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //
func getRequestSignature(signatureParams []map[string]string, method string, baseUrl string) string {
	builder := strings.Builder{}

	// 1. Build the parameter string
	// The parameter string is a concatenated list of parameters, joined by '&'.
	// Each key and value are percent encoded (see encode.go) before appending to the string.
	// NOTE: The key/value pairs must be sorted lexicographically
	for idx, entry := range signatureParams {
		for k, v := range entry {
			if v == "" {
				continue
			}

			fmt.Fprintf(&builder, "%s=%s", percentEncode(k), percentEncode(v))

			if idx < len(signatureParams)-1 {
				builder.WriteString("&")
			}
		}
	}

	parameterString := builder.String()
	builder.Reset()

	// 2. Build the signature string
	// To build the signature base string:
	//    - Convert the HTTP method to uppercase and set signature string to this value
	//    - Append '&'
	//    - Percent encode the URL and append to the signature string
	//    - Append '&'
	//    - Percent encode the parameter string and append to signature string
	fmt.Fprintf(&builder, "%s&%s&%s", method, percentEncode(baseUrl), percentEncode(parameterString))
	signatureBaseString := builder.String()
	builder.Reset()

	fmt.Fprintf(&builder, "%s&%s", percentEncode(credentials.apiSecret), percentEncode(credentials.oauthTokenSecret))

	// 3. Create the signing key
	// The signing key is a hash created from two components:
	//    - The previously created signature base string
	//    - The concatenation of the percent encoded consumer secret and percent encoded oauth token secret
	// The signature base string is passed as the data string to the HMAC-SHA1 hash function and the
	// signing key is passed as the key to the HMAC-SHA1 hash function
	signingKey := builder.String()
	return hmacHash(signatureBaseString, signingKey)
}
