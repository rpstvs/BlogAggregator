package main

import "net/http"

func Readiness(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondwithJSON(w, http.StatusOK, response{Status: "ok"})

}
