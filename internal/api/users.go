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
	ApiRoute(w, req, endpoint.Users, http.MethodGet, nil, nil)
}

func GetUserByUsername(w http.ResponseWriter, req *http.Request) {
	user_name := req.URL.Query().Get("user_name")

	if len(user_name) == 0 {
		error_msg := "`user_name` is a required query parameter but has not been supplied"
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error_msg))
		return
	}

	data, status_code, err := tools.RequestData(endpoint.UserByUsername(user_name), nil, http.MethodGet, nil)

	w.WriteHeader(status_code)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
