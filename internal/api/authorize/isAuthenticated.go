package authorize

import (
	"encoding/json"
	"gotwitter/internal/tools/utils/response"
	"net/http"
	"os"
)

// this doesn't verify the oauth token/secret are valid. Only that they exist in the application context
func IsAuthenticated(w http.ResponseWriter, req *http.Request) {
	oauth_token := os.Getenv("OAUTH_TOKEN")
	oauth_token_secret := os.Getenv("OAUTH_TOKEN_SECRET")

	w.WriteHeader(http.StatusOK)

	if len(oauth_token) > 0 && len(oauth_token_secret) > 0 {
		resp_json, _ := json.Marshal(
			response.ApiResponseFromData(map[string]string{"authenticated": "true"}),
		)
		w.Write(resp_json)
		return
	}

	resp_json, _ := json.Marshal(
		response.ApiResponseFromData(map[string]string{"authenticated": "false"}),
	)

	w.Write(resp_json)
}
