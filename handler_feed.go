package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/BlogAggregator/internal/database"
)

func (cfg *apiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string
		Url  string
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondwithError(w, http.StatusNotFound, "couldnt decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldnt create feed")
		return
	}

	_, err = cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Feed follow not created")
		return
	}

	respondwithJSON(w, http.StatusOK, databaseFeedtoFeed(feed))
}

func (cfg *apiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := cfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondwithError(w, http.StatusBadRequest, "couldnt retrieve feeds")
		return
	}

	respondwithJSON(w, http.StatusOK, feed)
}
