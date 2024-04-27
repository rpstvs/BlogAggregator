package main

import "net/http"

func errorHandler(w http.ResponseWriter, r *http.Request) {

	respondwithError(w, http.StatusInternalServerError, "Internal Server Error")
}
