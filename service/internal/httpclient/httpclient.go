package httpclient

import (
	"net/http"
	"time"
)

type HttpClienter interface {
	Get(url string) (resp *http.Response, err error)
}

func New() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
	}
}
