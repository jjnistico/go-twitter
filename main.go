package main

import (
	"log"
	"net/http"

	"gotwitter/internal/api"
	"gotwitter/internal/tools"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/callback", api.Callback)
	http.HandleFunc("/authenticate", api.GetOAuthToken)
	http.HandleFunc("/tweets/by/id", api.GetTweetsByIds)
	http.HandleFunc("/users/by/username", api.GetUsersByUsername)

	http.ListenAndServe(":8090", tools.LogRequest(http.DefaultServeMux))
}
