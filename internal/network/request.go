package network

import (
	"fmt"
	"gotwitter/internal/auth"
	"io"
	"io/ioutil"
	"net/http"
)

type gtRequest struct {
	req *http.Request
}

func newRequest(endpoint string, queryString string, method string, payload io.Reader) *gtRequest {
	var (
		req *http.Request
		err error
	)

	if method == http.MethodPost {
		req, err = http.NewRequest(method, endpoint+"?"+queryString, payload)
		req.Header.Add("content-type", "application/json")
	} else {
		req, err = http.NewRequest(method, endpoint+"?"+queryString, nil)
	}

	if err != nil {
		panic(fmt.Sprintf("error generating request: %s", err.Error()))
	}
	r := gtRequest{req}
	return &r
}

func (r *gtRequest) authorize() *gtRequest {
	auth.OAuth2Authorize(r.req)
	return r
}

func (r *gtRequest) execute() (any, error) {
	client := GetHttpClient()

	resp, err := client.Do(r.req)

	if err != nil {
		panic("query execution error")
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("twitter api response %d: %s", resp.StatusCode, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic("error reading response body from request")
	}

	return data, nil
}
