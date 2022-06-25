package api

import (
	"net/http"

	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools/utils"
)

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func GetTweets(w http.ResponseWriter, req *http.Request) {
	ApiRoute(w, req, endpoint.Tweets, http.MethodGet, []string{"ids"})
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
func DeleteTweet(w http.ResponseWriter, req *http.Request) {
	tweet_id := utils.GetPathParameterFromQuery(w, req, "id")

	if len(tweet_id) == 0 {
		return
	}

	ApiRoute(w, req, endpoint.TweetById(tweet_id), http.MethodDelete, nil)
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
func PostTweet(w http.ResponseWriter, req *http.Request) {
	ApiRoute(w, req, endpoint.Tweets, http.MethodPost, nil)
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
