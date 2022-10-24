package routes

import (
	"net/http"
	"task/config"
	"task/handlers"
	"task/repositories"

	"github.com/gorilla/mux"
)

func WalletRoutes(r *mux.Router) {
	repo := repositories.NewRepository(config.DB)
	h := handlers.NewHandler(repo)

	r.HandleFunc("/wallet", h.CreateWallet).Methods(http.MethodPost)
	r.HandleFunc("/wallets", h.FindWallet).Methods(http.MethodGet)
	r.HandleFunc("/wallet/{id}", h.UpdateWallet).Methods(http.MethodPatch)
	r.HandleFunc("/wallets/{id}", h.DeleteWallet).Methods(http.MethodDelete)
}
