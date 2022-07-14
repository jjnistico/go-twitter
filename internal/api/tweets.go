package api

import (
	"gotwitter/internal/endpoint"
	"gotwitter/internal/types"
	"net/http"
)

type GetTweetsOptions struct {
	Ids         []string
	Expansions  []string
	TweetFields []string
}

type ITweets interface {
	Get(request_params interface{}) (types.TweetsResponse, []types.Error)
}

type Tweets struct {
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
//   Parameters:
//     `ids` []string - array of ids to query **REQUIRED**
//     `expansions` []string - array of expansions (see link for available expansions)
//     `media.fields` []string - array of media fields to include **NOTE** requires attachments.media_keys expansion
//     `place.fields` []string - array of place fields to include **NOTE** requires geo.place_id expansion
//     `poll.fields`  []string - array of poll fields to include **NOTE** requires attachment.poll_ids expansion
//     `tweet.fields` []string - array of tweet fields to include
//     `user.fields`  []string - array of user fields to include  **NOTE** requires certain expansions (see link)
//
func (t *Tweets) Get(options types.GOTOptions) ([]types.TweetData, []types.Error) {
	response := ApiRequest[types.TweetsResponse](endpoint.Tweets, http.MethodGet, options, []string{"ids"}, nil)
	return response.Data().Data, response.Data().Errors
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
// func GetTweets(params url.Values) types.TweetsResponse {
// 	response := ApiRequest(endpoint.Tweets, http.MethodGet, params, []string{"ids"}, nil)
// 	return response.GoType.(types.TweetsResponse)
// }

// // https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
// func DeleteTweet(w http.ResponseWriter, req *http.Request) {
// 	tweet_id, query_params, err := utils.ExtractParameterFromQuery(req.URL.Query(), "id")

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	response := ApiRequest(endpoint.TweetById(tweet_id), http.MethodDelete, query_params, nil, nil)
// 	w.WriteHeader(response.Status())
// 	w.Write(response.ByteData())
// }
