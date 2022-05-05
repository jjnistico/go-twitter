package api

import (
	"fmt"
	"gotwitter/internal/tools"
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
func GetOAuthToken(w http.ResponseWriter, req *http.Request) {
	token_response, err := tools.RequestToken(w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(token_response)
}
