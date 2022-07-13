package api

import (
	"net/http"
)

// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func GetTimelineTweets(w http.ResponseWriter, req *http.Request) {
	// user_id, query_params, err := utils.ExtractParameterFromQuery(req.URL.Query(), "user_id")

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// response := ApiRequest(endpoint.TimelineTweets(user_id), http.MethodGet, query_params, nil, nil)
	// w.WriteHeader(response.Status())
	// w.Write(response.Data.([]byte))
}
