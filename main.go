package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// apiKey := os.Getenv("TWITTER_API_KEY")
	// apiSecret := os.Getenv("TWITTER_API_SECRET")
	// oauthToken := os.Getenv("TWITTER_OAUTH_TOKEN")
	// oauthTokenSecret := os.Getenv("TWITTER_OAUTH_TOKEN_SECRET")

	// got := gt.NewClient(apiKey, apiSecret, oauthToken, oauthTokenSecret)
}
