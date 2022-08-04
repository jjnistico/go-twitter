package gotwit

import (
	"fmt"
	"gotwitter/internal/network"
)

type Tweets struct{}

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
//   Parameters:
//     `ids` []string - array of tweet ids to query **REQUIRED**
//     `expansions` []string - array of expansions (see link for available expansions)
//     `media.fields` []string - array of media fields to include. Requires `attachments.media_keys` expansion
//     `place.fields` []string - array of place fields to include. Requires `geo.place_id` expansion
//     `poll.fields`  []string - array of poll fields to include. Requires `attachment.poll_ids` expansion
//     `tweet.fields` []string - array of tweet fields to include.
//     `user.fields`  []string - array of user fields to include. Requires certain expansions (see link)
//
func (*Tweets) Get(options ...getOption) tweetsResponse {
	urlVals := buildQueryParamsFromOptions(options)
	response, err := network.Get[tweetsResponse](tweetsEndpoint, urlVals)
	if err != nil {
		fmt.Println(err.Error())
	}
	return response
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
// Payload parameters:
//    `text` string - text of the tweet being created **REQUIRED** if `media.media_ids` is not present
//    `direct_message_deep_link` string - tweets a link directly to a Direct Message conversation with an account
//    `for_super_followers_only` bool - allows you to tweet exclusively to Super Followers
//    `geo.place_id` string - place id being attached to the Tweet for geo location
//    `media.media_ids` []string - a list of media ids being attached to the tweet. Required if request includes `tagged_user_ids`
//    `media.tagged_user_ids` []string - a list of user ids being tagged in the Tweet with media
//    `poll.duration_minutes` int - Duration of the poll in minutes for a Tweet with a poll
//
func (*Tweets) Create(payload CreateTweetPayload) createTweetResponse {
	response, err := network.Post[createTweetResponse](tweetsEndpoint, payload)
	if err != nil {
		fmt.Println(err.Error())
	}
	return response
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
// Parameters:
//    `id` string - the id of the tweet to be deleted
//
func (*Tweets) Delete(tweetId string) deleteTweetResponse {
	response, err := network.Delete[deleteTweetResponse](tweetByIdEndpoint(tweetId))
	if err != nil {
		fmt.Println(err.Error())
	}
	return response
}

type attachments struct {
	MediaKeys []string `json:"media_keys"`
	PollIds   []string `json:"poll_ids"`
}

type referencedTweets struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type reply struct {
	ExcludeReplyUserIds []string `json:"exclude_reply_user_ids"`
	InReplyToTweetId    string   `json:"in_reply_to_tweet_id"`
}

type CreateTweetPayload struct {
	DirectMessageDeepLink string `json:"direct_message_deep_link"`
	ForSuperFollowersOnly string `json:"for_super_followers_only"`
	GeoLocation           geo    `json:"geo"`
	MediaInformation      media  `json:"media"`
	PollOptions           poll   `json:"poll"`
	QuoteTweetId          string `json:"quote_tweet_id"`
	ReplyInformation      reply  `json:"reply"`
	ReplySettings         string `json:"reply_settings"`
	Text                  string `json:"text"`
}

type createTweetPayload struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type createTweetResponse struct {
	Data   createTweetPayload `json:"data"`
	Errors []gterror          `json:"errors"`
}

type deleteTweetPayload struct {
	Deleted bool `json:"deleted"`
}

type deleteTweetResponse struct {
	Data   deleteTweetPayload `json:"data"`
	Errors []gterror          `json:"errors"`
}

type tweetData struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	// optional fields are below
	CreatedAt          string             `json:"created_at"`
	AuthorId           string             `json:"author_id"`
	ConversationId     string             `json:"conversation_id"`
	InReplyToUserId    string             `json:"in_reply_to_user_id"`
	ReferencedTweets   []referencedTweets `json:"referenced_tweets"`
	Attachments        attachments        `json:"attachments"`
	Geo                geo                `json:"geo"`
	ContextAnnotations []struct {
		Domain struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Entity      struct {
			} `json:"entity"`
		} `json:"domain"`
	} `json:"context_annotations"`
}

type tweetsResponse struct {
	Data   []tweetData `json:"data"`
	Errors []gterror   `json:"errors"`
}
