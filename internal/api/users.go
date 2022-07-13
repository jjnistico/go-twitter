package api

import (
	"net/http"
)

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func GetUsers(w http.ResponseWriter, req *http.Request) {
	// response := ApiRequest[types.UsersResponse](endpoint.Users, http.MethodGet, req.URL.Query(), []string{"ids"}, req.Body)
}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func GetUserByUsername(w http.ResponseWriter, req *http.Request) {
	// user_name, query_params, err := utils.ExtractParameterFromQuery(req.URL.Query(), "user_name")

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// response := ApiRequest(endpoint.UserByUsername(user_name), http.MethodGet, query_params, nil, nil)
	// w.WriteHeader(response.Status())
	// w.Write(response.Data.([]byte))
}
