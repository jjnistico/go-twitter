package main

import (
	"fmt"
	"log"
	"os"

	gt "gotwitter/client"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecret := os.Getenv("TWITTER_API_SECRET")
	oauthToken := os.Getenv("TWITTER_OAUTH_TOKEN")
	oauthTokenSecret := os.Getenv("TWITTER_OAUTH_TOKEN_SECRET")

	got := gt.NewClient(apiKey, apiSecret, oauthToken, oauthTokenSecret)
	getUserOptions := gt.GetUserByUsername{
		Expansions:  []string{"pinned_tweet_id"},
		TweetFields: []string{"created_at"},
		UserFields:  []string{"created_at", "description", "pinned_tweet_id", "verified"},
	}
	users := got.Users.GetByUsername("jack", getUserOptions)
	fmt.Printf("%+v\n", users)
}
