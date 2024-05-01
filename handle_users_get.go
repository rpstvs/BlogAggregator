package main

import (
	"net/http"

	"github.com/rpstvs/BlogAggregator/auth"
)

func (cfg *apiConfig) GetUserbyKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)

	if err != nil {
		respondwithError(w, http.StatusBadRequest, err.Error())
	}

	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		respondwithError(w, http.StatusNotFound, "user not found")
		return
	}

	respondwithJSON(w, http.StatusOK, databaseUsertoUser(user))
}
