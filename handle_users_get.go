package main

import (
	"net/http"
	"strings"
)

func (cfg *apiConfig) GetUserbyKey(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")

	if apiKey == "" {
		respondwithError(w, http.StatusBadRequest, "Not authorized")
		return
	}

	tmp := strings.Split(apiKey, " ")

	apiKey = tmp[1]

	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		respondwithError(w, http.StatusNotFound, "user not found")
		return
	}

	respondwithJSON(w, http.StatusOK, user)
}
