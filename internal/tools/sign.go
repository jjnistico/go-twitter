package tools

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //
//                                                                                           //
// SignRequest generates a parameter string that can be included                             //
// with a request to get authorized for the twitter api.                                     //
// see: https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature //
//                                                                                           //
// Required parts of the param string (must be in alphabetical order):                       //
// oauth_consumer_key: string - the applications api_key                                     //
// oauth_nonce: string - unique string generated for every request                           //
// oauth_signature_method: string - HMAC-SHA1                                                //
// oauth_timestamp: string - unix timestamp                                                  //
// oauth_token: string - token returned from /authentication                                 //
// oauth_version: string - should always be 1.0 for the Twitter API                          //
//                                                                                           //
// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //
func GetRequestSignature(
	signature_params []map[string]string,
	method string,
	base_url string,
	query_params url.Values,
) string {
	builder := strings.Builder{}
	// start with encoded query params
	builder.WriteString(fmt.Sprintf("%s&", query_params.Encode()))
	// append all signature params encoded
	for idx, entry := range signature_params {
		for k, v := range entry {
			builder.WriteString(fmt.Sprintf("%s=%s", PercentEncode(k), PercentEncode(v)))
		}
		if idx < len(signature_params)-1 {
			builder.WriteString("&")
		}
	}
	// parameter string for appending to method + url for signature base string
	parameter_string := builder.String()
	builder.Reset()

	// signature base string is http method (to uppercase), percent encoded url and percent encoded parameter string
	// concatenated with '&'
	builder.WriteString(fmt.Sprintf("%s&%s&%s", method, PercentEncode(base_url), PercentEncode(parameter_string)))
	signature_base_string := builder.String()

	// the signing key is the concatenation of the consumer secret (API_SECRET) and the oauth_token_secret (&)
	oauth_consumer_secret := os.Getenv("API_SECRET")
	oauth_token_secret := os.Getenv("ACCESS_TOKEN_SECRET")
	signing_key := PercentEncode(oauth_consumer_secret) + "&" + PercentEncode(oauth_token_secret)

	hash := HmacHash(signature_base_string, signing_key)

	return hash
}
