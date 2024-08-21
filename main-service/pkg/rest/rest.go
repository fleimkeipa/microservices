package rest

import (
	"net/http"
	"time"
)

func CreateNewRestClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
