package authorize

import (
	"net/http"
	"os"
)

func IsAuthenticated(w http.ResponseWriter, req *http.Request) {
	oauth_token := os.Getenv("OAUTH_TOKEN")
	oauth_token_secret := os.Getenv("OAUTH_TOKEN_SECRET")

	w.WriteHeader(http.StatusOK)

	if len(oauth_token) > 0 && len(oauth_token_secret) > 0 {
		w.Write([]byte("true"))
		return
	}
	w.Write([]byte("false"))
}
