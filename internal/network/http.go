package network

import (
	"bytes"
	"encoding/json"
	"gotwitter/internal/types"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var lock = &sync.Mutex{}

var clientInstance *http.Client

func GetHttpClient() *http.Client {
	if clientInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		// prevent concurrency issues
		if clientInstance == nil {
			clientInstance = &http.Client{}
		}
	}

	return clientInstance
}

// GET request to twitter api
func Get[T types.ResponseData](
	endpoint string,
	options types.GOTOptions,
	required_params []string,
) (T, []types.Error) {
	return apiRequest[T](endpoint, http.MethodGet, options, required_params, nil)
}

// POST request to twitter api
func Post[T types.ResponseData](endpoint string, payload types.GOTPayload) (T, []types.Error) {
	return apiRequest[T](endpoint, http.MethodPost, nil, nil, payload)
}

// DELETE request to twitter api
func Delete[T types.ResponseData](endpoint string) (T, []types.Error) {
	return apiRequest[T](endpoint, http.MethodDelete, nil, nil, nil)
}

// apiRequest makes the full roundtrip request to the twitter api and returns the type passed as a type
// parameter as well as an array of any errors encountered through the request/response cycle. These errors
// are not the same errors that the twitter api can return, which are part of the response object.
func apiRequest[T types.ResponseData](
	endpoint string,
	method string,
	options types.GOTOptions,
	requiredParams []string,
	payload types.GOTPayload,
) (T, []types.Error) {
	// map options to url.Values for http request
	query_params := url.Values{}
	for k, v := range options {
		query_params.Set(k, strings.Join(v, ","))
	}

	// include buffer for post requests, nil otherwise (breaks if you pass nil *bytes.Buffer to NewRequest)
	var request *GOTRequest
	if method == http.MethodPost {
		// create buffer for payload for post requests
		payload_buf := new(bytes.Buffer)
		if err := json.NewEncoder(payload_buf).Encode(payload); err != nil {
			panic(err.Error())
		}
		request = NewRequest(endpoint, query_params, method, payload_buf)
	} else {
		request = NewRequest(endpoint, query_params, method, nil)
	}

	if requiredParams != nil {
		request.VerifyQueryParams(requiredParams)
	}

	// this creates the request signature and oauth header and adds it as the `Authorization` request header
	request.Authorize()

	data, errors := request.Execute()

	// unmarshal byte array to a GO type if no errors from execution of query, else return errors and nil data
	var structuredData T
	if len(errors) == 0 {
		if err := json.Unmarshal(data.([]byte), &structuredData); err != nil {
			panic(err.Error())
		}
	}

	return structuredData, errors
}
