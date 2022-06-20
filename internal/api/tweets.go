package api

import (
	"net/http"

	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools"
)

// see https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
// for complete response data/query parameters
type TweetsResponse struct {
	Data []struct {
		ID               string `json:"id"`
		Text             string `json:"text"`
		CreatedAt        string `json:"created_at"` // optional fields start here
		AuthorId         string `json:"author_id"`
		ConversationId   string `json:"conversation_id"`
		InReplyToUserId  string `json:"in_reply_to_user_id"`
		ReferencedTweets []struct {
			Type string `json:"type"`
			Id   string `json:"id"`
		}
		Attachments []struct {
			MediaKeys []interface{} `json:"media_keys"` // this needs to be updated to proper type
			PollIds   []string      `json:"poll_ids"`
		}
		Geo struct {
			GeoCoordinates struct {
				Type        string `json:"type"`
				Coordinates []int  `json:"coordinates"`
				PlaceId     string `json:"place_id"`
			} `json:"coordinates"`
		} `json:"geo"`
		ContextAnnotations []struct {
			Domain struct {
				Id          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Entity      struct {
				} `json:"entity"`
			} `json:"domain"`
		} `json:"context_annotations"`
	} `json:"data"`
}

// QUERY PARAMETERS:
// ids string      - a comma separated list of Tweet IDs. Max 100 per single request
// expansions enum - enable you to request additional data objects.
//     Available expansions:
//         attachments.poll_ids
//         attachments.media_keys
//         author_id
//         entities.mentions.username
//         geo.place_id
//         in_reply_to_user_id
//         referenced_tweets.id.author_id
// media.fields enum - select which specific media fields will be returned
// NOTE: Must pass `expansions=attachments.media_keys` in the query parameters
//     Available media.fields:
//         duration_ms
//         height
//         media_key
//         preview_image_url,
//         type
//         url
//         width
//         public_metrics
//         non_public_metrics
//         organic_metrics
//         promoted_metrics
//         alt_text
// place.fields enum - select specific place fields to return
// NOTE: Must pass `expansions=geo.place_id` in the query parameters
//    Available place.fields:
//        contained_within
//        country
//        country_code
//        full_name
//        geo
//        id
//        name
//        place_type
// poll.fields enum - select which specific poll fields will be returned
// NOTE: Must pass `expansions=attachments.poll_ids` in query parameters
//    Available poll.fields:
//        duration_minutes
//        end_datetime
//        id
//        options
//        voting_status
// tweet.fields enum - select which specific Tweet fields are returned
// NOTE: Must pass `expansions=referenced_tweets.id` in query parameters
//
func GetTweets(w http.ResponseWriter, req *http.Request) {
	ApiRoute(w, req, endpoint.Tweets, http.MethodGet, tools.EMPTY_PAYLOAD{})
}

func TweetById(w http.ResponseWriter, req *http.Request) {
	tweet_id := req.URL.Query().Get("id")

	if len(tweet_id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("`id` query parameter required for this endpoint"))
		return
	}

	q := req.URL.Query()
	q.Del("id")
	req.URL.RawQuery = q.Encode()

	ApiRoute(w, req, endpoint.TweetById(tweet_id), req.Method, tools.EMPTY_PAYLOAD{})
}

func Tweets(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fallthrough
	case http.MethodOptions:
		GetTweets(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("only [GET, OPTIONS] method allowed for this endpoint"))
	}
}

func TweetsById(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodDelete:
		fallthrough
	case http.MethodGet:
		TweetById(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("only [GET, DELETE] method allowed for this endpoint"))
	}
}
