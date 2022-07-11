package api

import (
	"net/http"

	"gotwitter/internal/endpoint"
	"gotwitter/internal/utils"
)

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func GetUsers(w http.ResponseWriter, req *http.Request) {
	response := ApiRequest(endpoint.Users, http.MethodGet, req.URL.Query(), []string{"ids"}, req.Body)
	w.WriteHeader(response.Status())
	w.Write(response.ByteData())
}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func GetUserByUsername(w http.ResponseWriter, req *http.Request) {
	user_name, query_params, err := utils.ExtractParameterFromQuery(req.URL.Query(), "user_name")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := ApiRequest(endpoint.UserByUsername(user_name), http.MethodGet, query_params, nil, nil)
	w.WriteHeader(response.Status())
	w.Write(response.ByteData())
}
