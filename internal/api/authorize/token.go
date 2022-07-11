package authorize

import (
	"gotwitter/internal/oauth"
	"gotwitter/internal/server"
	"net/http"
	"os"
)

func AuthenticateUser(w http.ResponseWriter, req *http.Request) {
	token_response, status_code, err := oauth.RequestToken()

	response := server.GOTResponse{}

	if err != nil {
		w.WriteHeader(status_code)
		response.AddError("error requesting token", err.Error(), "", "oauth")
		w.Write(response.JSON())
		return
	}

	authorize_resp, status_code, err := oauth.Authenticate(token_response.Token)

	w.WriteHeader(status_code)

	if err != nil {
		response.AddError("error authenticating user", err.Error(), "", "oauth")
		w.Write(response.JSON())
		return
	}

	response.SetData(authorize_resp)
	w.Write(response.JSON())
}

func AccessToken(w http.ResponseWriter, req *http.Request) {
	oauth_token := req.URL.Query().Get("oauth_token")
	oauth_verifier := req.URL.Query().Get("oauth_verifier")

	access_token_response, status_code, err := oauth.AccessToken(oauth_token, oauth_verifier)

	response := server.GOTResponse{}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.AddError("error obtaining access token", err.Error(), "", "oauth")
		w.Write(response.JSON())
		return
	}

	os.Setenv("OAUTH_TOKEN", access_token_response.OAuthToken)
	os.Setenv("OAUTH_TOKEN_SECRET", access_token_response.OAuthTokenSecret)

	response.SetData(access_token_response)
	w.WriteHeader(status_code)
	w.Write(response.JSON())
}
