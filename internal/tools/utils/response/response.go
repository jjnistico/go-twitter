package response

import (
	"encoding/json"
)

type Error struct {
	Detail     string `json:"detail"`
	Message    string `json:"message"`
	Title      string `json:"title"`
	Error_type string `json:"type"`
}

type Response struct {
	Errors       []Error     `json:"errors"`
	ResponseData interface{} `json:"data"`
}

// type ApiResponse interface {
// 	AddError(e Error)
// 	ResponseData(d interface{})
// 	JSON() (json []byte, err error)
// }

func (r *Response) AddError(title string, message string, detail string, error_type string) {
	r.Errors = append(r.Errors,
		Error{Title: title, Message: message, Error_type: error_type, Detail: detail},
	)
}

func (r *Response) Data(d interface{}) {
	r.ResponseData = d
}

func (r *Response) JSON() []byte {
	json, _ := json.Marshal(Response{ResponseData: r.ResponseData, Errors: r.Errors})

	return json
}
