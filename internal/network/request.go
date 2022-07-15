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
	req    *http.Request
	errors []types.Error
}

func (r *GOTRequest) AddError(title string, message string, detail string, errorType string) {
	gerr := types.Error{Title: title, Message: message, Detail: detail, Error_type: errorType}

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
	queryParams url.Values,
	httpMethod string,
	payload io.Reader,
) *GOTRequest {
	newRequest := GOTRequest{}

	req, err := http.NewRequest(httpMethod, endpoint+"?"+queryParams.Encode(), payload)

	if err != nil {
		newRequest.AddError(
			"request generation error",
			fmt.Sprintf(
				"error generating request for endpoint: %s | query: %s -> %s",
				endpoint,
				queryParams.Encode(),
				err.Error(),
			),
			"",
			"request")
		return &newRequest
	}

	if httpMethod == http.MethodPost {
		req.Header.Add("content-type", "application/json")
	}

	newRequest.req = req
	return &newRequest
}

func (r *GOTRequest) Authorize() {
	checkRequest(r.req)
	oauth.AuthorizeRequest(r.req)
}

func (r *GOTRequest) VerifyQueryParams(requiredParams []string) {
	checkRequest(r.req)
	queryParams := r.req.URL.Query()

	if err := utils.VerifyRequiredQueryParams(queryParams, requiredParams); err != nil {
		r.AddError("query parameter error", err.Error(), "", "request")
	}
}

func (r *GOTRequest) Execute() (data interface{}, errors []types.Error) {
	checkRequest(r.req)

	if len(r.Errors()) > 0 {
		return nil, r.Errors()
	}

	client := GetHttpClient()

	resp, err := client.Do(r.req)

	if err != nil {
		panic("query execution error")
	}

	// valid range 200-299
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		r.AddError(resp.Status, "twitter api response status", fmt.Sprint(resp.StatusCode), "query")
		return nil, r.Errors()
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		panic("error reading response body from request")
	}

	return data, nil
}

func checkRequest(req *http.Request) {
	if req == nil {
		panic("request not initialized")
	}
}
