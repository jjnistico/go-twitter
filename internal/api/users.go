package api

import (
	"fmt"
	"net/http"

	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools"
)

type UsersResponse struct {
	Data []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	data, status_code, err := tools.RequestData(endpoint.GetUsers, req.URL.Query(), http.MethodGet, nil)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(status_code)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(w, string(data))
}

func GetUsersByUsername(w http.ResponseWriter, req *http.Request) {
	data, status_code, err := tools.RequestData(endpoint.GetUsers, req.URL.Query(), http.MethodGet, nil)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(status_code)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(w, string(data))
}
