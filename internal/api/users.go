package api

import (
	"net/http"

	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools/utils"
)

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func GetUsers(w http.ResponseWriter, req *http.Request) {
	ApiRoute(w, req, endpoint.Users, http.MethodGet, nil)
}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func GetUserByUsername(w http.ResponseWriter, req *http.Request) {
	user_name := utils.GetPathParameterFromQuery(w, req, "user_name")

	if len(user_name) == 0 {
		return
	}

	ApiRoute(w, req, endpoint.UserByUsername(user_name), http.MethodGet, nil)
}
