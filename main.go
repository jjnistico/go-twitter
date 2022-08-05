package main

import (
	"fmt"
	"log"
	"os"

	g "gotwitter/client"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecret := os.Getenv("TWITTER_API_SECRET")
	clientId := os.Getenv("TWITTER_API_CLIENT_ID")

	got := g.NewClient(apiKey, apiSecret, clientId)

	userResponse := got.Users.GetByUsername("jack")

	fmt.Println("=====Users=====")
	fmt.Printf("%+v\n", userResponse.Data)

	tweetResponse := got.TimelineTweets.Get("235235")
	// tweetResponse := got.Tweets.Get(
	// 	g.With("ids", "32", "1123346", "908934727234"),
	// 	g.With("expansions", "attachments.poll_ids", "author_id", "entities.mentions.username"))

	// tweetResponse := got.Tweets.Count(g.With("query", "ufo"))
	fmt.Println("=====Tweets=====")
	fmt.Printf("%+v\n", tweetResponse.Data)
	fmt.Printf("%+v\n", tweetResponse.Errors)
}
