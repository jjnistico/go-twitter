package api

import (
	"fmt"
	"net/http"
	"os"
)

func Callback(w http.ResponseWriter, req *http.Request) {
	query_params := req.URL.Query()
	fmt.Println(query_params)
	os.Setenv("oauth_token", query_params.Get("oauth_token"))
	os.Setenv("oauth_verifier", query_params.Get("oauth_verifier"))
}
