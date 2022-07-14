package main

import (
	"fmt"
	"log"

	got "gotwitter/internal"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	got_client := got.NewClient()

	user, errors := got_client.Users.GetByUsername("jack", got.Options{
		"expansions":   {"pinned_tweet_id"},
		"user.fields":  {"created_at", "withheld", "verified"},
		"tweet.fields": {"created_at"},
	})

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", errors)
}
