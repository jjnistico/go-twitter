package main

import (
	"fmt"
	"log"

	got "gotwitter/internal"
	t "gotwitter/internal/types"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	got_client := got.NewClient()

	users, errors := got_client.Users.Get(t.GOTOptions{
		"ids":          {"1531882200"},
		"expansions":   {"pinned_tweet_id"},
		"tweet.fields": {"author_id", "id"},
		"user.fields":  {"created_at", "id", "description", "name", "username", "verified"},
	})

	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	fmt.Println()

	for _, error := range errors {
		fmt.Printf("%+v\n", error)
	}
}
