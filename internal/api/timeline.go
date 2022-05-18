package api

import (
	"fmt"
	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools"
	"net/http"
)

func GetUserTimeline(w http.ResponseWriter, req *http.Request) {
	data, status_code, err := tools.RequestData(endpoint.UserTimeline, req.URL.Query(), http.MethodGet, nil)

	w.WriteHeader(status_code)

	if err != nil {
		fmt.Printf("error requesting user timeline: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
