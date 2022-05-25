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

func GetTweets(w http.ResponseWriter, req *http.Request) {
	data, status_code, err := tools.RequestData(endpoint.GetTweets, req.URL.Query(), http.MethodGet, nil)

	w.WriteHeader(status_code)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}

func GetTweetsByIds(w http.ResponseWriter, req *http.Request) {
	ids := req.URL.Query().Get("ids")

	if len(ids) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("`ids` query parameter not present or not populated"))
		return
	}

	data, status_code, err := tools.RequestData(endpoint.GetTweets, req.URL.Query(), http.MethodGet, nil)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(status_code)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
