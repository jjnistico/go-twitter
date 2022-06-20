package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools"
	"net/http"
)

func ApiRoute(w http.ResponseWriter, req *http.Request, api_endpoint string, http_method string, payload interface{}) {
	// options request used to query available query parameters for endpoint
	if req.Method == http.MethodOptions {
		options := endpoint.GetEndpointOptions(api_endpoint)
		options_json, err := json.Marshal(options)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unable to serialize endpoint options"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(options_json)
		return
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(payload)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data, status_code, err := tools.RequestData(api_endpoint, req.URL.Query(), http_method, &buf)

	w.WriteHeader(status_code)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
