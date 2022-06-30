package authorize

import (
	"gotwitter/internal/tools/utils/response"
	"net/http"
	"os"
)

// this doesn't verify the oauth token/secret are valid. Only that they exist in the application context
func IsAuthenticated(w http.ResponseWriter, req *http.Request) {
	oauth_token := os.Getenv("OAUTH_TOKEN")
	oauth_token_secret := os.Getenv("OAUTH_TOKEN_SECRET")

	w.WriteHeader(http.StatusOK)

	var response response.Response

	if len(oauth_token) > 0 && len(oauth_token_secret) > 0 {
		response.Data(map[string]string{"authenticated": "true"})
		w.Write(response.JSON())
		return
	}

	response.Data(map[string]string{"authenticated": "false"})
	w.Write(response.JSON())
}
