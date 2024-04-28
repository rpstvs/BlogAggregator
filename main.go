package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rpstvs/BlogAggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DBURL")
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Printf("Couldnt open a connection to the database")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("POST /v1/users", apiCfg.CreateUser)
	mux.HandleFunc("GET /v1/readiness", Readiness)
	mux.HandleFunc("GET /v1/err", errorHandler)

	cors := middlewareCors(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: cors,
	}

	server.ListenAndServe()
}
