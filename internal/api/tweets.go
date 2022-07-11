package api

import (
	"net/http"

	"gotwitter/internal/endpoint"
	"gotwitter/internal/utils"
)

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func GetTweets(w http.ResponseWriter, req *http.Request) {
	response := ApiRequest(endpoint.Tweets, http.MethodGet, req.URL.Query(), []string{"ids"}, nil)
	w.WriteHeader(response.Status())
	w.Write(response.ByteData())
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
func DeleteTweet(w http.ResponseWriter, req *http.Request) {
	tweet_id, query_params, err := utils.ExtractParameterFromQuery(req.URL.Query(), "id")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := ApiRequest(endpoint.TweetById(tweet_id), http.MethodDelete, query_params, nil, nil)
	w.WriteHeader(response.Status())
	w.Write(response.ByteData())
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
func PostTweet(w http.ResponseWriter, req *http.Request) {
	response := ApiRequest(endpoint.Tweets, http.MethodPost, req.URL.Query(), nil, nil)
	w.WriteHeader(response.Status())
	w.Write(response.ByteData())
}

func Tweets(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fallthrough
	case http.MethodOptions:
		GetTweets(w, req)
	case http.MethodDelete:
		DeleteTweet(w, req)
	case http.MethodPost:
		PostTweet(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("only [GET, OPTIONS, DELETE, POST] method allowed for this endpoint"))
	}
}
