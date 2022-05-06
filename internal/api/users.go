package api

import (
	"fmt"
	"net/http"

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
		w.Write([]byte("Unable to parse query params, please key with `usernames`\n"))
		return
	}

	data, err := tools.RequestData("2/users/by", "usernames="+usernames, http.MethodGet, nil)

	if err != nil {
		fmt.Printf("Error getting users: %s", err)
	}

	fmt.Fprint(w, string(data))
}
