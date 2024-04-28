package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Name       string    `json:"name"`
	Created_at time.Time `json: "created_at"`
	Updated_at time.Time `json: "updated_at`
	Id         uuid.UUID `json:"id"`
}

func (cfg *apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondwithError(w, http.StatusBadRequest, "couldnt decode params")
		return
	}

}
