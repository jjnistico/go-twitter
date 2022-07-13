package server

import (
	gerror "gotwitter/internal/error"
	"gotwitter/internal/types"
)

type ResponseT interface {
	types.TweetsResponse | types.UsersResponse
}

type GOTResponse[T ResponseT] struct {
	data   T
	errors []gerror.Error
	status int
}

func NewResponse[T ResponseT](data T, errors []gerror.Error, status int) *GOTResponse[T] {
	resp := &GOTResponse[T]{}

	resp.data = data
	resp.errors = errors
	resp.status = status

	return resp
}

func (r *GOTResponse[_]) AddError(title string, message string, detail string, error_type string) {
	if r.errors == nil {
		r.errors = []gerror.Error{}
	}

	r.errors = append(r.errors,
		gerror.Error{Title: title, Message: message, Error_type: error_type, Detail: detail},
	)
}

func (r *GOTResponse[T]) Data() T {
	return r.data
}

func (r *GOTResponse[_]) Errors() []gerror.Error {
	return r.errors
}

func (r *GOTResponse[_]) Status() int {
	return r.status
}
