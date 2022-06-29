package api

import (
	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools/utils"
	"net/http"
)

func GetFollows(w http.ResponseWriter, req *http.Request) {
	user_id := utils.GetPathParameterFromQuery(w, req, "id")

	if len(user_id) == 0 {
		return
	}

	ApiRoute(w, req, endpoint.FollowersById(user_id), http.MethodGet, nil)
}
