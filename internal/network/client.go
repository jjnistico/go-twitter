package network

import (
	"net/http"
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
