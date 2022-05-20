package authorize

import (
	"encoding/json"
	"fmt"
	"gotwitter/internal/tools/oauth"
	"net/http"
	"os"
)

// // // // // // // // // // // // // // // // // // // // // // // //  //
// The following are specific to the application making the request      //
// API_KEY = oauth_consumer_key                                          //
// API_SECRET = oauth_consumer_secret                                    //
//                                                                       //
// The following are tokens used to authentication on behalf of the user //
// ACCESS_TOKEN = oauth_token                                            //
// ACCESS_SECRET = oauth_secret                                          //
// // // // // // // // // // // // // // // // // // // // // // // //  //
func AuthenticateUser(w http.ResponseWriter, req *http.Request) {
	token_response, status_code, err := oauth.RequestToken()

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(status_code)
		w.Write([]byte(err.Error()))
		return
	}

	authorize_resp, status_code, err := oauth.Authorize(token_response.Token)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(status_code)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(authorize_resp)
}

func AccessToken(w http.ResponseWriter, req *http.Request) {
	oauth_token := req.URL.Query().Get("oauth_token")
	oauth_verifier := req.URL.Query().Get("oauth_verifier")

	access_token_response, status_code, err := oauth.AccessToken(oauth_token, oauth_verifier)

	if err != nil {
		fmt.Printf("Error getting access token: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	os.Setenv("OAUTH_TOKEN", access_token_response.OAuthToken)
	os.Setenv("OAUTH_TOKEN_SECRET", access_token_response.OAuthTokenSecret)

	response, err := json.Marshal(&access_token_response)

	if err != nil {
		fmt.Printf("error marshalling access token response: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(status_code)
	w.Write(response)
}
