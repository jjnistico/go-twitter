package server

import (
	"encoding/json"
	gerror "gotwitter/internal/error"
)

type GOTResponse struct {
	Data   interface{}    `json:",omitempty"`
	Errors []gerror.Error `json:",omitempty"`
	status int
}

func NewResponse(data interface{}, errors []gerror.Error, status int) *GOTResponse {
	new_response := GOTResponse{}

	new_response.Data = data
	new_response.Errors = errors
	new_response.status = status

	return &new_response
}

func (r *GOTResponse) AddError(title string, message string, detail string, error_type string) {
	if r.Errors == nil {
		r.Errors = []gerror.Error{}
	}

	r.Errors = append(r.Errors,
		gerror.Error{Title: title, Message: message, Error_type: error_type, Detail: detail},
	)
}

func (r *GOTResponse) JSON() []byte {
	json, err := json.Marshal(GOTResponse{Data: r.Data, Errors: r.Errors})

	if err != nil {
		panic(err)
	}

	return json
}

func (r *GOTResponse) SetData(d interface{}) {
	r.Data = d
}

func (r *GOTResponse) SetStatus(s int) {
	r.status = s
}

func (r *GOTResponse) Status() int {
	return r.status
}
