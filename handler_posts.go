package main

import (
	"net/http"
	"strconv"

	"github.com/rpstvs/BlogAggregator/internal/database"
)

func (cfg *apiConfig) GetPostsbyUser(w http.ResponseWriter, r *http.Request, user database.User) {

	limitPostsStr := r.URL.Query().Get("limit")

	limit := 10

	if specifiedLimit, err := strconv.Atoi(limitPostsStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{

		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldnt retrieve posts for this user")
		return
	}

	respondwithJSON(w, http.StatusOK, sliceDatabasePoststoSlicePosts(posts))
}
