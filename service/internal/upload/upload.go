package upload

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pitsanujiw/go-health-check/internal/errorutil"
	"github.com/pitsanujiw/go-health-check/internal/healthcheck"
	"github.com/pitsanujiw/go-health-check/internal/reader"
)

type uploadHandler struct {
	healthcheckSrv healthcheck.Healthchecker
}

type UploadHandler interface {
	UploadFileHandler(w http.ResponseWriter, r *http.Request)
}

func New(checkerSvc healthcheck.Healthchecker) *uploadHandler {
	return &uploadHandler{
		healthcheckSrv: checkerSvc,
	}
}

func (u *uploadHandler) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(errorutil.MessageError{
			Message: "Upload bad request",
		})

		return
	}

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorutil.MessageError{
			Message: "File not found",
		})

		return
	}

	defer file.Close()

	urls, err := reader.ReadFile(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorutil.MessageError{
			Message: "Cannot read file",
		})

		return
	}

	result := u.healthcheckSrv.Ping(urls)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
