package tools

import (
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
			// don't handle cors preflight here
			if s_fetch != "cors" {
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

	response := response.Response{}
	response.Data(map[string][]string{"options": options})

	w.WriteHeader(http.StatusOK)
	w.Write(response.JSON())
}

func LogRequest(req *http.Request) {
	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
}
