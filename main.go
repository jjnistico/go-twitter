package main

import (
	"fmt"
	"log"

	"gotwitter/internal/api"
	"gotwitter/internal/client"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	got_client := client.New()

	tweets_resp := got_client.Tweets.Get(api.GetTweetsOptions{
		Ids: []string{"32", "123452352", "328428943"},
		Expansions: []string{
			"attachments.poll_ids",
			"author_id",
			"entities.mentions.username"},
		TweetFields: []string{
			"author_id",
			"created_at",
			"entities",
		},
	})

	for _, terr := range tweets_resp.Errors {
		fmt.Println(terr)
	}

	fmt.Printf("%+v\n", tweets_resp.Data)
}
