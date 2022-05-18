package oauth

import (
	"fmt"
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
	query_params string,
) string {
	builder := strings.Builder{}
	for idx, entry := range signature_params {
		for k, v := range entry {
			builder.WriteString(fmt.Sprintf("%s=%s", PercentEncode(k), PercentEncode(v)))
		}
		if idx < len(signature_params)-1 {
			builder.WriteString("&")
		}
	}

	parameter_string := builder.String()
	builder.Reset()

	builder.WriteString(fmt.Sprintf("%s&%s&%s", method, PercentEncode(base_url), PercentEncode(parameter_string)))
	signature_base_string := builder.String()
	builder.Reset()

	oauth_consumer_secret := os.Getenv("API_SECRET")
	oauth_token_secret := os.Getenv("OAUTH_TOKEN_SECRET")

	builder.WriteString(fmt.Sprintf("%s&%s", PercentEncode(oauth_consumer_secret), PercentEncode(oauth_token_secret)))

	signing_key := builder.String()
	hash := HmacHash(signature_base_string, signing_key)

	return hash
}
