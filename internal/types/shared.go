package types

type Coordinates struct {
	Type        string `json:"type"`
	Coordinates [2]int `json:"coordinates"`
}

type EntitiesT struct {
	Hashtags     []interface{} `json:"hashtags"`
	Symbols      []interface{} `json:"symbols"`
	UserMentions []UserMention `json:"user_mentions"`
}

type Error struct {
	Detail     string `json:"detail"`
	Message    string `json:"value"`
	Title      string `json:"title"`
	Error_type string `json:"type"`
}

type Geo struct {
	Coordinates Coordinates `json:"coordinates"`
	PlaceId     string      `json:"place_id"`
}

type Media struct {
	MediaIds      []string `json:"media_ids"`
	TaggedUserIds []string `json:"tagged_user_ids"`
}

type Poll struct {
	DurationMinutes uint     `json:"duration_minutes"`
	Options         []string `json:"options"`
}

type StartEnd struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type GOTOptions = map[string][]string

type GOTPayload = map[string]string
