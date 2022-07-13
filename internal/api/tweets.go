package api

import (
	"encoding/json"
	"gotwitter/internal/endpoint"
	gerror "gotwitter/internal/error"
	"gotwitter/internal/types"
	"net/http"
	"net/url"
	"strings"
)

type GetTweetsOptions struct {
	Ids         []string // required
	Expansions  []string
	TweetFields []string
}

type ITweets interface {
	Get(request_params interface{}) (types.TweetsResponse, []gerror.Error)
}

type Tweets struct {
}

func (t *Tweets) Get(options GetTweetsOptions) (types.TweetsResponse, []gerror.Error) {
	request_params := url.Values{}

	request_params.Set("ids", strings.Join(options.Ids, ","))
	request_params.Set("expansions", strings.Join(options.Expansions, ","))
	request_params.Set("tweet.fields", strings.Join(options.TweetFields, ","))

	response := ApiRequest(endpoint.Tweets, http.MethodGet, request_params, []string{"ids"}, nil)

	if len(response.Errors) > 0 {
		return types.TweetsResponse{}, response.Errors
	}

	var tweetsData types.TweetsResponse
	if err := json.Unmarshal(response.Data.([]byte), &tweetsData); err != nil {
		return types.TweetsResponse{}, []gerror.Error{{Title: "response unmarshal error", Message: "", Detail: "", Error_type: ""}}
	}

	return tweetsData, nil
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
