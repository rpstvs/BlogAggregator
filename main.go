package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("GET /v1/readiness", Readiness)
	mux.HandleFunc("GET /v1/err", errorHandler)

	cors := middlewareCors(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: cors,
	}

	server.ListenAndServe()
}
