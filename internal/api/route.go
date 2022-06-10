package api

import (
	"fmt"
	"gotwitter/internal/tools"
	"net/http"
)

func ApiRoute(w http.ResponseWriter, req *http.Request, endpoint string, http_method string) {
	data, status_code, err := tools.RequestData(endpoint, req.URL.Query(), http_method, nil)

	w.WriteHeader(status_code)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
