package main

import (
	"log"
	"net/http"

	"gotwitter/internal/api"
	"gotwitter/internal/api/authorize"
	"gotwitter/internal/tools"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/callback", authorize.Callback)
	http.HandleFunc("/authenticate", authorize.AuthenticateUser)
	http.HandleFunc("/api/tweets", api.GetTweets)
	http.HandleFunc("/api/tweets/by/id", api.GetTweetsByIds)
	http.HandleFunc("/api/users/by/username", api.GetUsersByUsername)
	http.HandleFunc("/api/users", api.GetUsers)

	http.ListenAndServe(":8090", tools.LogRequest(http.DefaultServeMux))
}
