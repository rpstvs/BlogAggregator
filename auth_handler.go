package main

import (
	"net/http"

	"github.com/rpstvs/BlogAggregator/auth"
	"github.com/rpstvs/BlogAggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondwithError(w, http.StatusUnauthorized, "couldn't find api key")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondwithError(w, http.StatusNotFound, "user not found")
			return
		}

		handler(w, r, user)

	}
}
