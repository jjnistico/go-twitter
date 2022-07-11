package gerror

type Error struct {
	Detail     string
	Message    string
	Title      string
	Error_type string
}

type GOTError interface {
	AddError(title string, message string, detail string, error_type string) []Error
	Errors() []Error
}
