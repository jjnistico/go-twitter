package api

import (
	"gotwitter/internal/endpoint"
	"gotwitter/internal/utils"
	"net/http"
)

// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func GetTimelineTweets(w http.ResponseWriter, req *http.Request) {
	user_id, new_params, err := utils.ExtractParameterFromQuery(req.URL.Query(), "user_id")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := ApiRequest(endpoint.TimelineTweets(user_id), http.MethodGet, new_params, nil, nil)
	w.WriteHeader(response.Status())
	w.Write(response.JSON())
}
