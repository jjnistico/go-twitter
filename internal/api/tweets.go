package api

import (
	"gotwitter/internal/endpoint"
	"gotwitter/internal/network"
	"gotwitter/internal/types"
)

type GetTweetsOptions struct {
	Ids         []string
	Expansions  []string
	TweetFields []string
}

type Tweets struct {
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
//   Parameters:
//     `ids` []string - array of ids to query **REQUIRED**
//     `expansions` []string - array of expansions (see link for available expansions)
//     `media.fields` []string - array of media fields to include. Requires `attachments.media_keys` expansion
//     `place.fields` []string - array of place fields to include. Requires `geo.place_id` expansion
//     `poll.fields`  []string - array of poll fields to include. Requires attachment.poll_ids` expansion
//     `tweet.fields` []string - array of tweet fields to include.
//     `user.fields`  []string - array of user fields to include. Requires certain expansions (see link)
//
func (t *Tweets) Get(options types.GOTOptions) ([]types.TweetData, []types.Error) {
	response, errors := network.Get[types.TweetsResponse](endpoint.Tweets, options, []string{"ids"})
	if errors != nil {
		return []types.TweetData{}, errors
	}
	return response.Data, response.Errors
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
func (t *Tweets) Create(payload types.GOTPayload) (types.CreateTweet, []types.Error) {
	response, errors := network.Post[types.CreateTweetResponse](endpoint.Tweets, payload)
	if errors != nil {
		return types.CreateTweet{}, errors
	}
	return response.Data, response.Errors
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
// Payload parameters:
//    `id` string - the id of the tweet to be deleted
//
func (t *Tweets) Delete(payload types.GOTPayload) (types.DeleteTweet, []types.Error) {
	tweet_id := payload["id"]
	response, errors := network.Delete[types.DeleteTweetResponse](endpoint.TweetById(tweet_id))
	if errors != nil {
		return types.DeleteTweet{}, errors
	}
	return response.Data, response.Errors
}
