package main

import "net/http"

func (cfg *apiConfig) GetUserbyKey(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")

	if apiKey == "" {
		respondwithError(w, http.StatusBadRequest, "Not authorized")
		return
	}

	apiKey = apiKey[len("ApiKey "):]

	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		respondwithError(w, http.StatusNotFound, "user not found")
	}

	respondwithJSON(w, http.StatusOK, user)
}
