package network

import (
	"net/http"
	"sync"
)

var lock sync.Mutex

var clientInstance *http.Client

func GetHttpClient() *http.Client {
	lock.Lock()
	defer lock.Unlock()

	if clientInstance == nil {
		clientInstance = &http.Client{}
	}

	return clientInstance
}
