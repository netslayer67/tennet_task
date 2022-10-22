package main

import (
	"encoding/json"
	"net/http"
	"task/config"
	"task/migration"
	"task/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Database Init
	config.DatabaseInit()
	migration.RunMigration()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
	})

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	http.ListenAndServe(":5000", r)
}
