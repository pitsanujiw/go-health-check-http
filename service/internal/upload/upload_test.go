package upload

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	checker_mock "github.com/pitsanujiw/go-health-check/mocks"
)

func TestUploadHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHttpClient := checker_mock.NewMockHttpClient(ctrl)

	t.Run("Should return upload bad request", func(t *testing.T) {
		uploadCtrl := uploadHandler{
			httpClientInstance: mockHttpClient,
		}
		req := httptest.NewRequest(http.MethodGet, "/api/v1/upload", nil)
		rr := httptest.NewRecorder()

		uploadCtrl.UploadFileHandler(rr, req)
		res := rr.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		require.Nil(t, deep.Equal(string(data), "Upload bad request\n"))
	})

	t.Run("Should return file not found", func(t *testing.T) {
		body := new(bytes.Buffer)

		mw := multipart.NewWriter(body)

		mw.Close()

		uploadCtrl := uploadHandler{
			httpClientInstance: mockHttpClient,
		}
		req := httptest.NewRequest(http.MethodPost, "/api/v1/upload", body)
		req.Header.Add("Content-Type", mw.FormDataContentType())

		rr := httptest.NewRecorder()

		uploadCtrl.UploadFileHandler(rr, req)
		res := rr.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		require.Nil(t, deep.Equal(string(data), "File not found\n"))
	})
}
