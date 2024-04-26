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
	cors := middlewareCors(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: cors,
	}

	server.ListenAndServe()
}
