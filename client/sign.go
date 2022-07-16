package gotwit

import (
	"fmt"
	"strings"
)

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
func getRequestSignature(
	signatureParams []map[string]string,
	method string,
	baseUrl string,
	queryParams string,
) string {
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

			builder.WriteString(fmt.Sprintf("%s=%s", percentEncode(k), percentEncode(v)))

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
	builder.WriteString(fmt.Sprintf("%s&%s&%s", method, percentEncode(baseUrl), percentEncode(parameterString)))
	signatureBaseString := builder.String()
	builder.Reset()

	builder.WriteString(fmt.Sprintf("%s&%s", percentEncode(credentialSvc.ApiSecret()), percentEncode(credentialSvc.OAuthTokenSecret())))

	// 3. Create the signing key
	// The signing key is a hash created from two components:
	//    - The previously created signature base string
	//    - The concatenation of the percent encoded consumer secret and percent encoded oauth token secret
	// The signature base string is passed as the data string to the HMAC-SHA1 hash function and the
	// signing key is passed as the key to the HMAC-SHA1 hash function
	signingKey := builder.String()
	hash := hmacHash(signatureBaseString, signingKey)

	return hash
}
