package tools

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
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
	oauth_token string,
	oauth_token_secret string,
	method string,
	base_url string,
	query_params url.Values,
) string {
	// the order of keys is important and since maps are unordered in go, must provide an array
	signature_keys := []string{
		"include_entities",
		"oauth_consumer_key",
		"oauth_nonce",
		"oauth_signature_method",
		"oauth_timestamp",
		"oauth_token",
		"oauth_version",
	}

	signature_params := map[string]string{
		"include_entities":       "true",
		"oauth_consumer_key":     os.Getenv("API_KEY"),
		"oauth_nonce":            GenerateNonce(42),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        fmt.Sprint(time.Now().Unix()),
		"oauth_token":            url.QueryEscape(oauth_token),
		"oauth_version":          "1.0",
	}

	// parameter string are fields joined by '&'. All key/values are percent encoded
	builder := strings.Builder{}
	for _, key := range signature_keys {
		// if signature_params[key] == "" {
		// 	continue
		// }
		builder.WriteString(fmt.Sprintf("%s=%s&", key, signature_params[key]))
	}
	// encode current string before appending query_params
	encoded_base := url.QueryEscape(builder.String())

	builder.Reset()

	builder.WriteString(encoded_base)

	// append encoded query params to end
	builder.WriteString(query_params.Encode())
	parameter_string := builder.String()

	builder.Reset()

	// signature base string is http method (to uppercase), percent encoded url and percent encoded parameter string
	// concatenated with '&'
	builder.WriteString(fmt.Sprintf("%s&%s&%s", method, url.QueryEscape(base_url), parameter_string))

	signature_base_string := builder.String()

	fmt.Println(signature_base_string)
	fmt.Println("========")
	// the signing key is the concatenation of the consumer secret (API_SECRET) and the oauth_token_secret (&)
	oauth_consumer_secret := os.Getenv("API_SECRET")
	signing_key := url.QueryEscape(oauth_consumer_secret) + "&" + url.QueryEscape(oauth_token_secret)

	hash := HmacHash(signature_base_string, signing_key)

	fmt.Println(hash)
	fmt.Println("======")
	return hash
}
