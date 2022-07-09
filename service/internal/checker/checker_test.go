package checker

import (
	"net/http"
	"testing"
	"time"
)

func TestPingUrl(t *testing.T) {
	t.Run("it should return true", func(t *testing.T) {
		url := "http://google.com"

		ins := httpClient{
			http: &http.Client{},
		}

		err := ins.PingUrl(url)
		if err != nil {
			t.Error("Expected no err, but got err")
		}
	})
}

func TestPing(t *testing.T) {

	t.Run("it should return true", func(t *testing.T) {
		url := "http://google.com"

		ins := httpClient{
			http: &http.Client{},
		}

		urls := []string{url}

		_ = ins.Ping(urls)
	})
}

func TestFormatPingResult(t *testing.T) {
	totalWebsites := 2
	success := 1
	failure := 1
	totalTime := time.Duration(123450000)

	t.Run("it should return formatted ping result", func(t *testing.T) {
		want := Result{
			TotalWebsite:    2,
			Success:         1,
			Failure:         1,
			TotalTimeMin:    0,
			TotalTimeSec:    0,
			TotalTimeMiniSec: 123,
		}

		got := FormatPingResult(totalWebsites, success, failure, totalTime)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
