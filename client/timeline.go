package gotwit

import (
	"fmt"
	"gotwitter/internal/network"
)

type TimelineTweets struct{}

// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func (*TimelineTweets) Get(userId string, options ...getOption) tweetsResponse {
	urlVals := buildQueryParamsFromOptions(options)
	response, err := network.Get[tweetsResponse](timelineTweetsEndpoint(userId), urlVals)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

type TimelineTweetsData struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
