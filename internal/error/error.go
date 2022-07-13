package gerror

type Error struct {
	Detail     string `json:"detail"`
	Message    string `json:"value"`
	Title      string `json:"title"`
	Error_type string `json:"type"`
}

type GOTError interface {
	error
	AddError(title string, message string, detail string, error_type string) []Error
	Errors() []Error
}
