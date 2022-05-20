package authorize

import (
	"fmt"
	"gotwitter/internal/tools/oauth"
	"net/http"
	"os"
)

func Callback(w http.ResponseWriter, req *http.Request) {
	query_params := req.URL.Query()

	oauth_token := query_params.Get("oauth_token")
	oauth_verifier := query_params.Get("oauth_verifier")

	os.Setenv("OAUTH_TOKEN", oauth_token)
	os.Setenv("OAUTH_VERIFIER", oauth_verifier)

	access_token_response, status_code, err := oauth.AccessToken(oauth_token, oauth_verifier)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(status_code)
		w.Write([]byte(err.Error()))
		return
	}

	os.Setenv("OAUTH_TOKEN", access_token_response.OAuthToken)
	os.Setenv("OAUTH_TOKEN_SECRET", access_token_response.OAuthTokenSecret)
	os.Setenv("USER_ID", access_token_response.UserId)
	os.Setenv("SCREEN_NAME", access_token_response.ScreenName)

	w.WriteHeader(status_code)
}
