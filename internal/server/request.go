package server

import (
	"fmt"
	gerror "gotwitter/internal/error"
	"gotwitter/internal/oauth"
	"gotwitter/internal/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GOTRequest struct {
	endpoint     string
	query_params url.Values
	http_method  string
	payload      io.Reader
	req          *http.Request
	errors       []gerror.Error
}

func (r *GOTRequest) AddError(title string, message string, detail string, error_type string) {
	gerr := gerror.Error{Title: title, Message: message, Detail: detail, Error_type: error_type}

	if r.errors == nil {
		r.errors = []gerror.Error{}
	}

	r.errors = append(r.errors, gerr)
}

func (r *GOTRequest) Errors() []gerror.Error {
	return r.errors
}

func NewRequest(
	endpoint string,
	query_params url.Values,
	http_method string,
	payload io.Reader,
) *GOTRequest {
	new_request := GOTRequest{}

	new_request.endpoint = endpoint
	new_request.query_params = query_params
	new_request.http_method = http_method
	new_request.payload = payload

	req, err := http.NewRequest(
		new_request.http_method,
		new_request.endpoint+"?"+new_request.query_params.Encode(),
		new_request.payload,
	)

	if err != nil {
		new_request.AddError(
			"request generation error",
			fmt.Sprintf(
				"error generating request for endpoint: %s | query: %s -> %s",
				new_request.endpoint,
				new_request.query_params.Encode(),
				err.Error(),
			),
			"",
			"request")
		return &new_request
	}

	if new_request.payload != nil {
		req.Header.Add("content-type", "application/json")
	}

	new_request.req = req
	return &new_request
}

func (r *GOTRequest) Authorize() {
	if r.req == nil {
		fmt.Println("request not initialized")
		return
	}

	oauth.AuthorizeRequest(r.req)
}

func (r *GOTRequest) VerifyQueryParams(required_params []string) {
	query_params := r.req.URL.Query()

	if err := utils.VerifyRequiredQueryParams(query_params, required_params); err != nil {
		r.AddError("query parameter error", err.Error(), "", "request")
	}
}

func (r *GOTRequest) Execute() (data interface{}, status_code int, gerr []gerror.Error) {
	if r.req == nil {
		fmt.Println("request not initialized")
		return nil, http.StatusInternalServerError, nil
	}

	if len(r.Errors()) > 0 {
		return nil, http.StatusInternalServerError, r.Errors()
	}

	client := &http.Client{}

	resp, err := client.Do(r.req)

	if err != nil {
		r.AddError("query execution error", err.Error(), "", "query")
		return nil, resp.StatusCode, r.Errors()
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		r.AddError("response read error", err.Error(), "", "query")
		return nil, http.StatusInternalServerError, r.Errors()
	}

	return data, resp.StatusCode, nil
}
