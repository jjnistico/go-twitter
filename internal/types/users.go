package types

type UserUrls struct {
	StartEnd
	Url         string `json:"url"`
	ExpandedUrl string `json:"expanded_url"`
	DisplayUrl  string `json:"display_url"`
}

type UserUrl struct {
	Urls []UserUrls `json:"urls"`
}

type UserHashtags struct {
	StartEnd
	Hashtag string `json:"hashtag"`
}

type UserMentions struct {
	StartEnd
	Username string `json:"username"`
}

type UserDescription struct {
	Urls []UserUrls `json:"urls"`
}

type UserEntities struct {
	Url         []UserUrl         `json:"url"`
	Description []UserDescription `json:"description"`
}

type Withheld struct {
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type UserMention struct {
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ID         uint64 `json:"id"`
	IDString   string `json:"id_str"`
}

type UserData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	// optional fields
	CreatedAt   string   `json:"created_at"`
	Protected   bool     `json:"protected"`
	Withheld    Withheld `json:"withheld"`
	Location    string   `json:"location"`
	Url         string   `json:"url"`
	Description string   `json:"description"`
	Verified    bool     `json:"verified"`
}

type UsersResponse struct {
	Data []UserData `json:"data"`
	ErrorResponse
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
