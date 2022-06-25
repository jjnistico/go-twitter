package types

type UserMention struct {
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ID         uint64 `json:"id"`
	IDString   string `json:"id_str"`
}

type UsersResponse struct {
	Data []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
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
