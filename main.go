package main

import (
	"fmt"
	gt "gotwitter/client"
	"log"
	"os"

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

	userResponse := got.Users.GetByUsername("jack", gt.GetUserByUsernameOptions{})

	fmt.Println("=====Users=====")
	fmt.Printf("%+v\n", userResponse)

	tweetResponse := got.Tweets.Get(gt.GetTweetsOptions{
		Ids:        []string{"32", "324234355", "23498585"},
		Expansions: []string{"attachments.poll_ids", "author_id", "entities.mentions.username"},
	})

	fmt.Println("=====Tweets=====")
	fmt.Printf("%+v\n", tweetResponse.Data)
	fmt.Printf("%+v\n", tweetResponse.Errors)
}
