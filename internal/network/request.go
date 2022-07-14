package network

import (
	"fmt"
	"gotwitter/internal/oauth"
	"gotwitter/internal/types"
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
	errors       []types.Error
}

func (r *GOTRequest) AddError(title string, message string, detail string, error_type string) {
	gerr := types.Error{Title: title, Message: message, Detail: detail, Error_type: error_type}

	if r.errors == nil {
		r.errors = []types.Error{}
	}

	r.errors = append(r.errors, gerr)
}

func (r *GOTRequest) Errors() []types.Error {
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
		panic("request not initialized")
	}

	oauth.AuthorizeRequest(r.req)
}

func (r *GOTRequest) VerifyQueryParams(required_params []string) {
	query_params := r.req.URL.Query()

	if err := utils.VerifyRequiredQueryParams(query_params, required_params); err != nil {
		r.AddError("query parameter error", err.Error(), "", "request")
	}
}

func (r *GOTRequest) Execute() (data interface{}, errors []types.Error) {
	if r.req == nil {
		panic("request not initialized")
	}

	if len(r.Errors()) > 0 {
		return nil, r.Errors()
	}

	client := GetHttpClient()

	resp, err := client.Do(r.req)

	if err != nil {
		panic("query execution error")
	}

	if resp.StatusCode != http.StatusOK {
		r.AddError(resp.Status, "twitter api status response not ok", fmt.Sprint(resp.StatusCode), "query")
		return nil, r.Errors()
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		panic("error reading response body from request")
	}

	return data, nil
}
