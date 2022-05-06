package api

import (
	"fmt"
	"net/http"

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

	data, err := tools.RequestData("2/tweets/", "ids="+ids, http.MethodGet, nil)

	if err != nil {
		fmt.Printf("Error getting tweets: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(data))
}
