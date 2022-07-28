package network

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
func Get[T interface{}](endpoint string, opt interface{}) (T, error) {
	qs, _ := query.Values(opt)
	return apiRequest[T](endpoint, http.MethodGet, qs.Encode(), nil)
}

// POST request to twitter api
func Post[T interface{}](endpoint string, payload interface{}) (T, error) {
	return apiRequest[T](endpoint, http.MethodPost, "", payload)
}

// DELETE request to twitter api
func Delete[T interface{}](endpoint string) (T, error) {
	return apiRequest[T](endpoint, http.MethodDelete, "", nil)
}

// apiRequest makes the full roundtrip request to the twitter api and returns the type passed as a type
// parameter as well as a potential error encountered through the request/response cycle. This error
// is not the same as errors the twitter api can return, which are part of the response object.
func apiRequest[T interface{}](
	endpoint string,
	method string,
	queryString string,
	payload interface{},
) (T, error) {
	payloadBuf := new(bytes.Buffer)
	if err := json.NewEncoder(payloadBuf).Encode(payload); err != nil {
		panic(err.Error())
	}

	data, err := Execute(endpoint, queryString, method, payloadBuf)

	// unmarshal byte array to a GO type if no errors from execution of query, else return errors and nil data
	var structuredData T
	if err == nil {
		if err := json.Unmarshal(data.([]byte), &structuredData); err != nil {
			panic(err.Error())
		}
	}

	return structuredData, err
}
