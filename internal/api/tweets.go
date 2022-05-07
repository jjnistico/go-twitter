package api

import (
	"fmt"
	"net/http"

	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools"
)

type TweetsResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

func GetTweetsByIds(w http.ResponseWriter, req *http.Request) {
	ids := req.URL.Query().Get("ids")

	if len(ids) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("`ids` query parameter not present or not populated"))
		return
	}

	data, err := tools.RequestData(endpoint.GetTweets, "ids="+ids, http.MethodGet, nil)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(w, string(data))
}
