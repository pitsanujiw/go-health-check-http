package main

import (
	"net/http"
	"time"

	"github.com/pitsanujiw/go-health-check/internal/checker"
	"github.com/pitsanujiw/go-health-check/internal/upload"
)

func main() {
	httpClient := http.Client{
		Timeout: 15 * time.Second,
	}

	checkerHttpClient := checker.GetHttpClient(httpClient)

	// Upload route
	uploadInstance := upload.GetUploadHandler(checkerHttpClient)

	http.HandleFunc("/api/v1/upload", uploadInstance.UploadFileHandler)

	//Listen on port 8080
	http.ListenAndServe("0.0.0.0:9999", nil)

}
