package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHttpClient struct {
	err  error
	resp *http.Response
}

func (h *mockHttpClient) Get(url string) (resp *http.Response, err error) {
	return h.resp, h.err
}

func TestPingUrl(t *testing.T) {
	t.Run("it should return true", func(t *testing.T) {
		url := "http://__URL__.com"
		rr := httptest.NewRecorder()

		rr.WriteHeader(http.StatusOK)

		client := &mockHttpClient{
			err:  nil,
			resp: rr.Result(),
		}

		healthSrv := New(client)

		err := healthSrv.PingUrl(url)
		if err != nil {
			t.Error("Expected no err, but got err")
		}
	})
}

func TestPing(t *testing.T) {
	t.Run("it should return true", func(t *testing.T) {
		url := "http://__URL__.com"

		rr := httptest.NewRecorder()

		rr.WriteHeader(http.StatusOK)

		client := &mockHttpClient{
			err:  nil,
			resp: rr.Result(),
		}

		healthcheckSrv := New(client)
		urls := []string{url}

		result := healthcheckSrv.Ping(urls)
		if result.TotalWebsite == 0 {
			t.Error("Expected no err, but got err")
		}
	})
}
