package gotwit

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type gtRequest struct {
	req    *http.Request
	errors []gterror
}

func (r *gtRequest) addError(title string, message string, detail string, errorType string) {
	gerr := gterror{Title: title, Message: message, Detail: detail, Error_type: errorType}

	if r.errors == nil {
		r.errors = []gterror{}
	}

	r.errors = append(r.errors, gerr)
}

func (r *gtRequest) Errors() []gterror {
	return r.errors
}

func newRequest(
	endpoint string,
	queryString string,
	httpMethod string,
	payload io.Reader,
) *gtRequest {
	fmt.Println(queryString)
	req, err := http.NewRequest(httpMethod, endpoint+"?"+queryString, payload)

	if err != nil {
		panic(fmt.Sprintf("error generating request: %s", err.Error()))
	}

	if httpMethod == http.MethodPost {
		req.Header.Add("content-type", "application/json")
	}

	return &gtRequest{req, []gterror{}}
}

func (r *gtRequest) authorize() {
	authorizeRequest(r.req)
}

func (r *gtRequest) Execute() (data interface{}, errors []gterror) {
	if len(r.Errors()) > 0 {
		return nil, r.Errors()
	}

	client := getHttpClient()

	resp, err := client.Do(r.req)

	if err != nil {
		panic("query execution error")
	}

	// valid range 200-299
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		r.addError(resp.Status, "twitter api response status", fmt.Sprint(resp.StatusCode), "query")
		return nil, r.Errors()
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		panic("error reading response body from request")
	}

	return data, nil
}
