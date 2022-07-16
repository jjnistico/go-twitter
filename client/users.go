package gotwit

import "fmt"

type Users struct{}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
//   Parameters:
//     `ids` []string - array of user ids to query **REQUIRED**
//     `expansions` []string - array of expansions (see link for available expansions)
//     `tweet.fields` []string - array of tweet fields to include.
//     `user.fields`  []string - array of user fields to include.
//
func (*Users) Get(options GetUsers) usersResponse {
	response, errors := get[usersResponse](usersEndpoint, options)
	if errors != nil {
		fmt.Printf("%+v\n", errors)
	}
	return response
}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-user
//  Parameters:
//     `user_name` string - username to query **REQUIRED**
//     `expansions` []string - array of expansions (see link for available expansions)
//     `tweet.fields` []string - array of tweet fields to include.
//     `user.fields`  []string - array of user fields to include.
//
func (*Users) GetByUsername(username string, options GetUserByUsername) userResponse {
	response, errors := get[userResponse](userByUsernameEndpoint(username), options)
	if errors != nil {
		fmt.Printf("%+v\n", errors)
	}
	return response
}

type userUrls struct {
	startEnd
	Url         string `json:"url"`
	ExpandedUrl string `json:"expanded_url"`
	DisplayUrl  string `json:"display_url"`
}

type userUrl struct {
	Urls []userUrls `json:"urls"`
}

type userHashtags struct {
	startEnd
	Hashtag string `json:"hashtag"`
}

type userCashtags struct {
	startEnd
	Cashtag string `json:"cashtag"`
}

type userMentions struct {
	startEnd
	Username string `json:"username"`
}

type publicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type userDescription struct {
	Urls            []userUrls      `json:"urls"`
	Hashtags        []userHashtags  `json:"hashtags"`
	Mentions        []userMentions  `json:"mentions"`
	Cashtags        []userCashtags  `json:"cashtags"`
	ProfileImageUrl string          `json:"profile_image_url"`
	PublicMetrics   []publicMetrics `json:"public_metrics"`
	PinnedTweetId   string          `json:"pinned_tweet_id"`
	Includes        struct {
		Tweets []tweetData `json:"tweets"`
	} `json:"includes"`
	Errors []gterror `json:"errors"`
}

type userEntities struct {
	Url         []userUrl         `json:"url"`
	Description []userDescription `json:"description"`
}

type withheld struct {
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type userMention struct {
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ID         uint64 `json:"id"`
	IDString   string `json:"id_str"`
}

type userData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	// optional fields
	CreatedAt   string   `json:"created_at"`
	Protected   bool     `json:"protected"`
	Withheld    withheld `json:"withheld"`
	Location    string   `json:"location"`
	Url         string   `json:"url"`
	Description string   `json:"description"`
	Verified    bool     `json:"verified"`
}

type userResponse struct {
	Data     userData `json:"data"`
	Includes struct {
		Tweets []tweetData `json:"tweets,omitempty"`
	} `json:"includes,omitempty"`
	Errors []gterror `json:"errors"`
}

type usersResponse struct {
	Data   []userData `json:"data"`
	Errors []gterror  `json:"errors"`
}

type userTimelineResponse struct {
	Data []struct {
		CreatedAt string    `json:"created_at"`
		ID        uint64    `json:"id"`
		IDString  string    `json:"id_str"`
		Text      string    `json:"text"`
		Truncated bool      `json:"truncated"`
		Entities  entitiesT `json:"entities"`
	} `json:"data"`
}

type GetUserByUsername struct {
	Expansions  []string `url:"expansions,omitempty,comma"`
	TweetFields []string `url:"tweet.fields,omitempty,comma"`
	UserFields  []string `url:"user.fields,omitempty,comma"`
}

type GetUsers struct {
	Ids         []string `url:"ids,comma"`
	Expansions  []string `url:"expansions,omitempty,comma"`
	TweetFields []string `url:"tweet.fields,omitempty,comma"`
	UserFields  []string `url:"user.fields,omitempty,comma"`
}
