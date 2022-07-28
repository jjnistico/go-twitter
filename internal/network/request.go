package network

import (
	"fmt"
	"gotwitter/internal/auth"
	"io"
	"io/ioutil"
	"net/http"
)

func Execute(
	endpoint string,
	queryString string,
	method string,
	payload io.Reader,
) (interface{}, error) {
	var (
		req *http.Request
		err error
	)

	if method == http.MethodPost {
		req, err = http.NewRequest(method, endpoint+"?"+queryString, payload)
	} else {
		req, err = http.NewRequest(method, endpoint+"?"+queryString, nil)
	}

	if err != nil {
		panic(fmt.Sprintf("error generating request: %s", err.Error()))
	}

	if method == http.MethodPost {
		req.Header.Add("content-type", "application/json")
	}

	reqUrl := req.URL.Scheme + "://" + req.URL.Host + req.URL.Path

	authHeader := auth.BuildAuthorizationHeader(req.Method, reqUrl, req.URL.Query())

	req.Header.Add("Authorization", authHeader)

	client := getHttpClient()

	resp, err := client.Do(req)

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
