package routes

import (
	"net/http"
	"task/handlers"
	"task/repositories"

	"task/config"

	"github.com/gorilla/mux"
)

func AssetRoutes(r *mux.Router) {
	repo := repositories.NewRepository(config.DB)
	h := handlers.NewHandler(repo)

	r.HandleFunc("/asset", h.CreateAsset).Methods(http.MethodPost)
	r.HandleFunc("/assets", h.FindAsset).Methods(http.MethodGet)
	r.HandleFunc("/asset/{id}", h.UpdateAsset).Methods(http.MethodPatch)
	r.HandleFunc("/assets/{id}", h.DeleteAsset).Methods(http.MethodDelete)
}
