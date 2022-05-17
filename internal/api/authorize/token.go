package authorize

import (
	"fmt"
	"gotwitter/internal/tools/oauth"
	"net/http"
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
