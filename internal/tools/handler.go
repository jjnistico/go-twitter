package tools

import (
	"log"
	"net/http"
)

func RequestHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		HandleCors(w, req)
		LogRequest(req)
		handler.ServeHTTP(w, req)
	})
}

func HandleCors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	}
}

func LogRequest(req *http.Request) {
	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
}
