package server

import (
	gerror "gotwitter/internal/error"
	"gotwitter/internal/types"
)

type ResponseT interface {
	types.TweetsResponse | types.UsersResponse
}

type GOTResponse[DT ResponseT] struct {
	Data   DT             `json:"data"`
	Errors []gerror.Error `json:"errors"`
	status int
}

func NewResponse[T ResponseT](data T, errors []gerror.Error, status int) *GOTResponse[T] {
	resp := &GOTResponse[T]{}

	resp.Data = data
	resp.Errors = errors
	resp.status = status

	return resp
}

func (r *GOTResponse[_]) AddError(title string, message string, detail string, error_type string) {
	if r.Errors == nil {
		r.Errors = []gerror.Error{}
	}

	r.Errors = append(r.Errors,
		gerror.Error{Title: title, Message: message, Error_type: error_type, Detail: detail},
	)
}

func (r *GOTResponse[T]) SetData(d T) {
	r.Data = d
}

func (r *GOTResponse[_]) Status() int {
	return r.status
}
