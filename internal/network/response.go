package network

import (
	"gotwitter/internal/types"
)

type ResponseData interface {
	types.TweetsResponse | types.UsersResponse
}

type GOTResponse[T ResponseData] struct {
	data   T
	errors []types.Error
}

func NewResponse[T ResponseData](data T, errors []types.Error) *GOTResponse[T] {
	resp := &GOTResponse[T]{}

	resp.data = data
	resp.errors = errors

	return resp
}

func NewError[T ResponseData](error_data []types.Error) *GOTResponse[T] {
	resp := &GOTResponse[T]{}

	resp.errors = error_data

	return resp
}

func (r *GOTResponse[_]) AddError(title string, message string, detail string, error_type string) {
	if r.errors == nil {
		r.errors = []types.Error{}
	}

	r.errors = append(r.errors,
		types.Error{Title: title, Message: message, Error_type: error_type, Detail: detail},
	)
}

func (r *GOTResponse[T]) Data() T {
	return r.data
}

func (r *GOTResponse[_]) Errors() []types.Error {
	return r.errors
}
