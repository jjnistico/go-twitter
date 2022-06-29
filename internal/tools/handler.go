package tools

import (
	"encoding/json"
	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools/utils/response"
	"log"
	"net/http"
)

func RequestHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		HandleCors(w, req)
		LogRequest(req)

		if req.Method == http.MethodOptions {
			s_fetch := req.Header.Get("Sec-Fetch-Mode")
			if s_fetch == "" {
				HandleOptionsRequest(w, req)
				return
			}
		}

		handler.ServeHTTP(w, req)
	})
}

func HandleCors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	}
}

func HandleOptionsRequest(w http.ResponseWriter, req *http.Request) {
	options := endpoint.GetEndpointOptions(req.URL.Path)

	options_json, err := json.Marshal(
		response.ApiResponseFromData(map[string][]string{"options": options}),
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json_resp, err := json.Marshal(response.OneOffErrorResponse(
			err.Error(), "unable to serialize endpoint options",
		))

		if err != nil {
			panic(err)
		}

		w.Write(json_resp)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(options_json)
}

func LogRequest(req *http.Request) {
	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
}
