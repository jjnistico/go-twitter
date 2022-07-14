package types

type Attachments struct {
	MediaKeys []string `json:"media_keys"`
	PollIds   []string `json:"poll_ids"`
}

type ReferencedTweetEnum = string

const (
	Retweeted ReferencedTweetEnum = "retweeted"
	Quoted    ReferencedTweetEnum = "quoted"
	RepliedTo ReferencedTweetEnum = "replied_to"
)

type ReferencedTweets struct {
	Type ReferencedTweetEnum `json:"type"`
	Id   string              `json:"id"`
}

type Reply struct {
	ExcludeReplyUserIds []string `json:"exclude_reply_user_ids"`
	InReplyToTweetId    string   `json:"in_reply_to_tweet_id"`
}

type TweetPayload struct {
	DirectMessageDeepLink string `json:"direct_message_deep_link"`
	ForSuperFollowersOnly string `json:"for_super_followers_only"`
	GeoLocation           Geo    `json:"geo"`
	MediaInformation      Media  `json:"media"`
	PollOptions           Poll   `json:"poll"`
	QuoteTweetId          string `json:"quote_tweet_id"`
	ReplyInformation      Reply  `json:"reply"`
	ReplySettings         string `json:"reply_settings"`
	Text                  string `json:"text"`
}

type TweetData struct {
	ID                 string             `json:"id"`
	Text               string             `json:"text"`
	CreatedAt          string             `json:"created_at"` // optional fields start here
	AuthorId           string             `json:"author_id"`
	ConversationId     string             `json:"conversation_id"`
	InReplyToUserId    string             `json:"in_reply_to_user_id"`
	ReferencedTweets   []ReferencedTweets `json:"referenced_tweets"`
	Attachments        Attachments        `json:"attachments"`
	Geo                Geo                `json:"geo"`
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

type TweetsResponse struct {
	Data []TweetData `json:"data"`
	ErrorResponse
}
