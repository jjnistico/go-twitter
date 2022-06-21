package types

type Error struct {
	Detail  string `json:"detail"`
	Message string `json:"message"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

type ApiResponse struct {
	Errors []Error     `json:"errors"`
	Data   interface{} `json:"data"`
}
