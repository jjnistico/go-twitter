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

	got_client := got.New()

	// tweets, errors := got_client.Tweets.Get(t.GOTOptions{
	// 	"ids":          {"32", "12345462", "324235235", "2342"},
	// 	"expansions":   {"attachments.poll_ids", "author_id", "entities.mentions.username"},
	// 	"tweet.fields": {"author_id", "created_at", "entities"},
	// })

	// for _, terr := range errors {
	// 	fmt.Println(terr.Title)
	// 	fmt.Println()
	// }

	// fmt.Println("--------")

	// for _, tweet := range tweets {
	// 	fmt.Println(tweet.AuthorId, tweet.ID, tweet.Text)
	// }

	// tweet_post, errors := got_client.Tweets.Create(t.GOTPayload{
	// 	"text": "Test",
	// })

	// for _, terr := range errors {
	// 	fmt.Println(terr.Title)
	// 	fmt.Println()
	// }

	// fmt.Println("---------")

	// fmt.Println(tweet_post.Id, tweet_post.Text)

	tweet_del, errors := got_client.Tweets.Delete(t.GOTPayload{"id": "23"})

	for _, terr := range errors {
		fmt.Println(terr.Title)
		fmt.Println()
	}

	fmt.Println("--------")

	fmt.Println(tweet_del.Deleted)
}
