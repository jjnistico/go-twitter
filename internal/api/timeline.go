package api

import (
	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools/utils"
	"net/http"
)

// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func GetTimelineTweets(w http.ResponseWriter, req *http.Request) {
	user_id := utils.GetPathParameterFromQuery(w, req, "user_id")

	if len(user_id) == 0 {
		return
	}

	ApiRoute(w, req, endpoint.TimelineTweets(user_id), http.MethodGet, nil)
}
