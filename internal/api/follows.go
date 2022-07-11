package api

import (
	"gotwitter/internal/endpoint"
	"gotwitter/internal/utils"
	"net/http"
)

func GetFollows(w http.ResponseWriter, req *http.Request) {
	user_id, new_params, err := utils.ExtractParameterFromQuery(req.URL.Query(), "id")

	if err != nil {
		return
	}

	response := ApiRequest(endpoint.FollowersById(user_id), http.MethodGet, new_params, nil, nil)
	w.WriteHeader(response.Status())
	w.Write(response.JSON())
}
