package upload

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pitsanujiw/go-health-check/internal/checker"
	"github.com/pitsanujiw/go-health-check/internal/reader"
)

type uploadHandler struct {
	httpClientInstance checker.HttpClient
}

type UploadHandler interface {
	UploadFileHandler(w http.ResponseWriter, r *http.Request)
}

func GetUploadHandler(httpClient checker.HttpClient) *uploadHandler {
	return &uploadHandler{
		httpClientInstance: httpClient,
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (u *uploadHandler) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, "Upload bad request\n")

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
		fmt.Fprintf(w, "File not found\n")
		return
	}

	defer file.Close()

	urls, err := reader.ReadFile(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot read file\n")
		return
	}

	result := u.httpClientInstance.Ping(urls)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
