package main

import (
	"net/http"
	"os"

	"github.com/pitsanujiw/go-health-check/internal/healthcheck"
	"github.com/pitsanujiw/go-health-check/internal/httpclient"
	"github.com/pitsanujiw/go-health-check/internal/upload"
)

func main() {

	httpClient := httpclient.New()

	healthSvc := healthcheck.New(httpClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	mux := http.NewServeMux()
	// Upload route
	uploadHandler := upload.New(healthSvc)

	mux.Handle("/api/v1/upload", cors(http.HandlerFunc(uploadHandler.UploadFileHandler)))

	//Listen on port 8080
	http.ListenAndServe(port, mux)

}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if r.Method == http.MethodOptions {
			if origin != "http://localhost:3000" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Max-Age", "5") // 5 seconds
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		h.ServeHTTP(w, r)
	})
}
