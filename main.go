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

	// oauth related endpoints
	http.HandleFunc("/callback", authorize.Callback)
	http.HandleFunc("/authenticate", authorize.AuthenticateUser)
	http.HandleFunc("/access_token", authorize.AccessToken)
	http.HandleFunc("/is_authenticated", authorize.IsAuthenticated)

	// tweets
	http.HandleFunc("/api/tweets", api.Tweets)

	// users
	http.HandleFunc("/api/users", api.GetUsers)
	http.HandleFunc("/api/user_by_username", api.GetUserByUsername)

	// timeline
	http.HandleFunc("/api/timeline_tweets", api.GetTimelineTweets)

	http.ListenAndServe(":8090", tools.RequestHandler(http.DefaultServeMux))
}
