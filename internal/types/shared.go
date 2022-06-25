package types

type ApiResponse struct {
	Errors []Error     `json:"errors"`
	Data   interface{} `json:"data"`
}

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
	Detail  string `json:"detail"`
	Message string `json:"message"`
	Title   string `json:"title"`
	Type    string `json:"type"`
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
