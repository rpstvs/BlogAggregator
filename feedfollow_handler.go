package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/BlogAggregator/internal/database"
)

func (cfg *apiConfig) RetrieveFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := cfg.DB.GetFeedFollows(r.Context())

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldnt retrieve feed follow")
		return
	}

	respondwithJSON(w, http.StatusOK, databaseFeedsFollowToFeedsFollow(feedFollow))
}

func (cfg *apiConfig) CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed_id uuid.UUID
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "couldnt decode parameters")
		return
	}

	feedfollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.Feed_id,
	})

	if err != nil {
		respondwithError(w, http.StatusBadRequest, "couldnt create feed follow")
		return
	}

	respondwithJSON(w, http.StatusOK, databaseFollowtoFollow(feedfollow))
}

func (cfg *apiConfig) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	id := r.PathValue("feedFollowID")

	followId, err := uuid.Parse(id)

	if err != nil {
		respondwithError(w, http.StatusBadRequest, "Couldnt parse feed follow id")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     followId,
		UserID: user.ID,
	})

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "couldnt delete the feed follow")
		return
	}
	w.WriteHeader(http.StatusOK)
}
