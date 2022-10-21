package main

import (
	"encoding/json"
	"net/http"
	"task/config"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Database Init
	config.DatabaseInit()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hello World")
	})

	http.ListenAndServe(":5000", r)
}
