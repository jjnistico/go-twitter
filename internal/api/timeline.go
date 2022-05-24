package api

import (
	"fmt"
	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools"
	"net/http"
)

type UserMention struct {
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ID         uint64 `json:"id"`
	IDString   string `json:"id_str"`
}

type EntitiesT struct {
	Hashtags     []interface{} `json:"hashtags"`
	Symbols      []interface{} `json:"symbols"`
	UserMentions []UserMention `json:"user_mentions"`
}

type UserTimelineResponse struct {
	Data []struct {
		CreatedAt string    `json:"created_at"`
		ID        uint64    `json:"id"`
		IDString  string    `json:"id_str"`
		Text      string    `json:"text"`
		Truncated bool      `json:"truncated"`
		Entities  EntitiesT `json:"entities"`
	} `json:"data"`
}

// AVAILABLE QUERY PARAMETERS FOR FILTERING RESPONSE - All query params are optional                 //
// user_id          number             id of the user for whome to return results
// screen_name      string             screen name of the user for whome to return results
// since_id         number             return results with more recent id than parameter
// count            number             specifies the number of tweets to try and retrieve (max 200)
// max_id           number             return results with less recent id than parameter
// trim_user        boolean | number   when set to true/1, each tweet returned includes only userid
// exclude_replies  boolean            prevents replies from appearing in returned timeline
// include_rts      boolean            when set to false, timeline strips any native retweets
// for more information, see https://developer.twitter.com/en/docs/twitter-api/v1/tweets/timelines/api-reference/get-statuses-user_timeline
//                                                                                                  //
func GetUserTimeline(w http.ResponseWriter, req *http.Request) {
	data, status_code, err := tools.RequestData(endpoint.UserTimeline, req.URL.Query(), http.MethodGet, nil)

	w.WriteHeader(status_code)

	if err != nil {
		fmt.Printf("error requesting user timeline: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}

type HomeTimelineResponse struct {
	Data []struct {
		CreatedAt string    `json:"created_at"`
		ID        uint64    `json:"id"`
		IDString  string    `json:"id_str"`
		Text      string    `json:"text"`
		Truncated bool      `json:"truncated"`
		Entities  EntitiesT `json:"entities"`
	} `json:"data"`
}

// AVAILABLE QUERY PARAMETERS FOR FILTERING RESPONSE - All query params are optional                 //
// since_id         number             return results with more recent id than parameter
// count            number             specifies the number of tweets to try and retrieve (max 200)
// max_id           number             return results with less recent id than parameter
// trim_user        boolean | number   when set to true/1, each tweet returned includes only userid
// exclude_replies  boolean            prevents replies from appearing in returned timeline
// include_entities boolean            the entities node will not be included when set to false
// for more information, see https://developer.twitter.com/en/docs/twitter-api/v1/tweets/timelines/api-reference/get-statuses-home_timeline
//                                                                                                  //
func GetHomeTimeline(w http.ResponseWriter, req *http.Request) {
	data, status_code, err := tools.RequestData(endpoint.HomeTimeline, req.URL.Query(), http.MethodGet, nil)

	w.WriteHeader(status_code)

	if err != nil {
		fmt.Printf("error requesting home timeline: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
