package upload

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pitsanujiw/go-health-check/internal/errorutil"
	"github.com/pitsanujiw/go-health-check/internal/healthcheck"
	"github.com/stretchr/testify/require"
)

type mockHealthcheck struct {
	result healthcheck.Result
}

func (m *mockHealthcheck) Ping(urls []string) healthcheck.Result {
	return m.result
}

func TestUploadHandler(t *testing.T) {
	t.Run("Should return upload bad request", func(t *testing.T) {
		handler := New(&mockHealthcheck{})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/upload", nil)
		rr := httptest.NewRecorder()

		handler.UploadFileHandler(rr, req)
		res := rr.Result()
		defer res.Body.Close()

		got := rr.Body.Bytes()
		msg := errorutil.MessageError{
			Message: "Upload bad request",
		}

		want, _ := json.Marshal(msg)

		require.Equal(t, string(got), string(want)+"\n")
	})

	t.Run("Should return file not found", func(t *testing.T) {
		body := new(bytes.Buffer)

		mw := multipart.NewWriter(body)

		mw.Close()

		handler := New(&mockHealthcheck{})

		req := httptest.NewRequest(http.MethodPost, "/api/v1/upload", body)
		req.Header.Add("Content-Type", mw.FormDataContentType())

		rr := httptest.NewRecorder()

		handler.UploadFileHandler(rr, req)
		res := rr.Result()
		defer res.Body.Close()

		got := rr.Body.Bytes()
		msg := errorutil.MessageError{
			Message: "File not found",
		}

		want, _ := json.Marshal(msg)

		require.Equal(t, string(got), string(want)+"\n")
	})

	t.Run("Should return result ping", func(t *testing.T) {
		body := new(bytes.Buffer)

		mw := multipart.NewWriter(body)

		part, _ := mw.CreateFormFile("file", "file.csv")
		part.Write([]byte(`sample`))

		mw.Close()

		handler := New(&mockHealthcheck{
			result: healthcheck.Result{},
		})

		req := httptest.NewRequest(http.MethodPost, "/api/v1/upload", body)
		req.Header.Add("Content-Type", mw.FormDataContentType())

		rr := httptest.NewRecorder()

		handler.UploadFileHandler(rr, req)

		res := rr.Result()
		defer res.Body.Close()

		got := rr.Body.Bytes()
		msg := healthcheck.Result{}

		want, _ := json.Marshal(msg)

		require.Equal(t, string(got), string(want)+"\n")
	})
}
