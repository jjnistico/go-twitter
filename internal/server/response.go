package server

import (
	"encoding/json"
	gerror "gotwitter/internal/error"
	"net/http"
)

type GOTResponse struct {
	Data   interface{}    `json:"data"`
	Errors []gerror.Error `json:"errors"`
	status int
}

func NewResponse(data interface{}, errors []gerror.Error, status int) *GOTResponse {
	new_response := GOTResponse{}

	var resp GOTResponse
	if err := json.Unmarshal(data.([]byte), &resp); err != nil {
		new_response.Data = nil
		new_response.AddError("error unmarshalling response", err.Error(), "", "response")
		new_response.SetStatus(http.StatusInternalServerError)
		return &new_response
	}

	new_response.Data = resp.Data
	new_response.Errors = resp.Errors
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

func (r *GOTResponse) ByteData() []byte {
	return r.Data.([]byte)
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
