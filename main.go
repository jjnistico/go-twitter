package main

import (
	"fmt"
	"log"

	"gotwitter/internal/client"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	got_client := client.New()

	tweets, errors := got_client.Tweets.Get(map[string][]string{
		"ids":          {"32", "12345462", "324235235", "2342"},
		"expansions":   {"attachments.poll_ids", "author_id", "entities.mentions.username"},
		"tweet.fields": {"author_id", "created_at", "entities"},
	})

	for _, terr := range errors {
		fmt.Println(terr.Title)
		fmt.Println()
	}

	fmt.Println("--------")

	for _, tweet := range tweets {
		fmt.Println(tweet.AuthorId, tweet.ID, tweet.Text)
	}
}
