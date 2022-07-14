package network

import (
	"bytes"
	"encoding/json"
	"gotwitter/internal/types"
	"net/http"
	"net/url"
	"strings"
)

func Get[T types.ResponseData](
	endpoint string,
	options types.GOTOptions,
	required_params []string,
) (T, []types.Error) {
	return apiRequest[T](endpoint, http.MethodGet, options, required_params, nil)
}

func Post[T types.ResponseData](endpoint string, payload types.GOTPayload) (T, []types.Error) {
	return apiRequest[T](endpoint, http.MethodPost, nil, nil, payload)
}

func Delete[T types.ResponseData](endpoint string) (T, []types.Error) {
	return apiRequest[T](endpoint, http.MethodDelete, nil, nil, nil)
}

func apiRequest[T types.ResponseData](
	endpoint string,
	method string,
	options types.GOTOptions,
	required_params []string,
	payload types.GOTPayload,
) (T, []types.Error) {
	// map options to url.Values for http request
	query_params := url.Values{}
	for k, v := range options {
		query_params.Set(k, strings.Join(v, ","))
	}

	// create buffer for payload for post requests
	payload_buf := new(bytes.Buffer)
	if err := json.NewEncoder(payload_buf).Encode(payload); err != nil {
		panic(err.Error())
	}

	// include buffer for post requests, nil otherwise (breaks if you pass *nil bytes.Buffer to NewRequest)
	var request *GOTRequest
	if method == http.MethodPost {
		request = NewRequest(endpoint, query_params, method, payload_buf)
	} else {
		request = NewRequest(endpoint, query_params, method, nil)
	}

	if required_params != nil {
		request.VerifyQueryParams(required_params)
	}

	// this creates the request signature and oauth header and adds it as the `Authorization` request header
	request.Authorize()

	data, errors := request.Execute()

	// unmarshal byte array to a GO type if no errors from execution of query, else return errors
	var structured_data T
	if len(errors) == 0 {
		if err := json.Unmarshal(data.([]byte), &structured_data); err != nil {
			panic(err.Error())
		}
	}

	return structured_data, errors
}
