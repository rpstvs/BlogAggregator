package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/BlogAggregator/internal/database"
)

func (cfg *apiConfig) CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed_id string
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "couldnt decode parameters")
		return
	}

	id := uuid.MustParse(params.Feed_id)

	feedfollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    id,
	})

	if err != nil {
		respondwithError(w, http.StatusBadRequest, "couldnt create feed follow")
		return
	}

	respondwithJSON(w, http.StatusOK, feedfollow)
}
