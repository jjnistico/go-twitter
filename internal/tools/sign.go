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
func SignRequest(
	oauth_token string,
	oauth_token_secret string,
	method string,
	base_url string,
	query_params string,
) (string, error) {
	// the order of keys is important and since maps are unordered in go, must provide an array
	signature_keys := []string{
		"oauth_consumer_key",
		"oauth_nonce",
		"oauth_signature_method",
		"oauth_timestamp",
		"oauth_token",
		"oauth_version",
	}

	signature_params := map[string]string{
		"oauth_consumer_key":     os.Getenv("API_KEY"),
		"oauth_nonce":            GenerateNonce(42),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        fmt.Sprint(time.Now().Unix()),
		"oauth_token":            url.QueryEscape(oauth_token),
		"oauth_version":          "1.0",
	}

	// parameter string are fields joined by '&'. All key/values are percent encoded
	parameter_string := ""
	for idx, key := range signature_keys {
		parameter_string += url.QueryEscape(key)
		parameter_string += "="
		parameter_string += url.QueryEscape(signature_params[key])
		parameter_string += "&"
		if idx == len(signature_keys)-1 {
			parameter_string += url.QueryEscape(query_params)
		}
	}

	// signature base string is http method (to uppercase), percent encoded url and percent encoded parameter string
	// concatenated with '&'
	signature_base_string := strings.ToUpper(method) + "&" + url.QueryEscape(base_url) + url.QueryEscape(parameter_string)

	// the signing key is the concatenation of the consumer secret (API_SECRET) and the oauth_token_secret (&)
	oauth_consumer_secret := os.Getenv("API_SECRET")
	signing_key := url.QueryEscape(oauth_consumer_secret) + "&" + url.QueryEscape(oauth_token_secret)

	hash := HmacHash(signature_base_string, signing_key)

	fmt.Printf("Request signature: %s", hash)

	return hash, nil
}
