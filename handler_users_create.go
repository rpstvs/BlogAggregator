package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/BlogAggregator/internal/database"
)

func (cfg *apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondwithError(w, http.StatusBadRequest, "couldnt decode params")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Body,
	})

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "couldnt create user")
	}

	respondwithJSON(w, http.StatusOK, user)

}
