package gotwit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/go-querystring/query"
)

var lock = &sync.Mutex{}

var clientInstance *http.Client

func getHttpClient() *http.Client {
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
func get[T responseData](endpoint string, opt interface{}) (T, []gterror) {
	qs, _ := query.Values(opt)
	return apiRequest[T](endpoint, http.MethodGet, qs.Encode(), nil)
}

// POST request to twitter api
func post[T responseData](endpoint string, payload interface{}) (T, []gterror) {
	return apiRequest[T](endpoint, http.MethodPost, "", payload)
}

// DELETE request to twitter api
func delete[T responseData](endpoint string) (T, []gterror) {
	return apiRequest[T](endpoint, http.MethodDelete, "", nil)
}

// apiRequest makes the full roundtrip request to the twitter api and returns the type passed as a type
// parameter as well as an array of any errors encountered through the request/response cycle. These errors
// are not the same errors that the twitter api can return, which are part of the response object.
func apiRequest[T responseData](
	endpoint string,
	method string,
	queryString string,
	payload interface{},
) (T, []gterror) {
	// include buffer for post requests, nil otherwise (breaks if you pass nil *bytes.Buffer to NewRequest)
	var request *gtRequest
	if method == http.MethodPost {
		// create buffer for payload for post requests
		payloadBuf := new(bytes.Buffer)
		if err := json.NewEncoder(payloadBuf).Encode(payload); err != nil {
			panic(err.Error())
		}
		request = newRequest(endpoint, queryString, method, payloadBuf)
	} else {
		request = newRequest(endpoint, queryString, method, nil)
	}

	// this creates the request signature and oauth header and adds it as the `Authorization` request header
	request.authorize()

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
