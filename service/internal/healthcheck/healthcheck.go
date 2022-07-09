package healthcheck

import (
	"time"

	"github.com/pitsanujiw/go-health-check/internal/httpclient"
)

type Result struct {
	TotalWebsite int           `json:"totalWebsite"`
	TotalSuccess int           `json:"totalSuccess"`
	TotalFailure int           `json:"totalFailure"`
	Milliseconds time.Duration `json:"milliseconds"`
}

type PingUrlResult struct {
	err error
}

type healthcheck struct {
	http httpclient.HttpClienter
}

type Healthchecker interface {
	Ping(urls []string) Result
}

func New(client httpclient.HttpClienter) *healthcheck {
	return &healthcheck{
		http: client,
	}
}

// Ping url with http client
func (c *healthcheck) PingUrl(url string) error {
	resp, err := c.http.Get(url)
	if err != nil {
		return err
	}

	resp.Body.Close()

	return nil
}

// Ping urls
func (c *healthcheck) Ping(urls []string) Result {
	totalWebsites := len(urls)
	start := time.Now()

	var success, failure int
	resultsChan := make(chan *PingUrlResult, totalWebsites)

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

	return Result{
		TotalWebsite: totalWebsites,
		TotalSuccess: success,
		TotalFailure: failure,
		Milliseconds: time.Duration(elapse.Milliseconds()),
	}
}
