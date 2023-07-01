package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) error {
	return writeJSON(w, http.StatusOK, "Hello Home")
}
