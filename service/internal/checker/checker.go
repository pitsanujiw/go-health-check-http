package checker

import (
	"net/http"
	"time"
)

type Result struct {
	TotalWebsite    int           `json:"totalWebsite"`
	Success         int           `json:"success"`
	Failure         int           `json:"failure"`
	TotalTimeMin    time.Duration `json:"totalTimeMin"`
	TotalTimeSec    time.Duration `json:"totalTimeSec"`
	TotalTimeMiniSec time.Duration `json:"totalTimeMiniSec"`
}

type PingUrlResult struct {
	err error
}

type httpClient struct {
	http *http.Client
}

type HttpClient interface {
	PingUrl(url string) error
	Ping(urls []string) Result
}

// Get http client
func GetHttpClient(client http.Client) *httpClient {
	return &httpClient{
		http: &client,
	}

}

// Ping url with http client
func (c *httpClient) PingUrl(url string) error {
	resp, err := c.http.Get(url)
	if err != nil {
		return err
	}

	resp.Body.Close()

	return nil
}

// Format ping result
func FormatPingResult(totalWebsites int, success int, failure int, elapse time.Duration) Result {
	return Result{
		TotalWebsite:    totalWebsites,
		Success:         success,
		Failure:         failure,
		TotalTimeMin:    time.Duration(elapse.Minutes()),
		TotalTimeSec:    time.Duration(elapse.Seconds()),
		TotalTimeMiniSec: time.Duration(elapse.Milliseconds()),
	}

}

// Ping urls
func (c *httpClient) Ping(urls []string) Result {
	totalWebsites := len(urls)
	start := time.Now()

	var success, failure int
	resultsChan := make(chan *PingUrlResult)

	defer close(resultsChan)

	for _, url := range urls {
		go func(url string) {
			err := c.PingUrl(url)
			result := &PingUrlResult{err}
			resultsChan <- result
		}(url)
	}

	var results []PingUrlResult
	for result := range resultsChan {
		results = append(results, *result)
		if result.err == nil {
			success += 1
		} else {
			failure += 1
		}

		if len(results) == totalWebsites {
			break
		}
	}

	elapse := time.Since(start)

	return FormatPingResult(totalWebsites, success, failure, elapse)
}
