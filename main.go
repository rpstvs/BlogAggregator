package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rpstvs/BlogAggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)

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
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.GetUserbyKey))

	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.CreateFeed))
	mux.HandleFunc("GET /v1/feeds", apiCfg.GetFeeds)

	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.DeleteFeedFollow))
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.CreateFeedFollow))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.RetrieveFeedFollow))

	mux.HandleFunc("GET /v1/readiness", Readiness)
	mux.HandleFunc("GET /v1/err", errorHandler)

	cors := middlewareCors(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: cors,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	server.ListenAndServe()
}
