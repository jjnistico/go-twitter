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

func GetUsersByUsername(w http.ResponseWriter, req *http.Request) {
	usernames := req.URL.Query().Get("usernames")

	if len(usernames) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("`usernames` query parameter not present or not populated"))
		return
	}

	data, err := tools.RequestData(endpoint.GetUsers, "usernames="+usernames, http.MethodGet, nil)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(w, string(data))
}
