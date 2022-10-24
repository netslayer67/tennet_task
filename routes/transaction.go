package routes

import (
	"net/http"
	"task/config"
	"task/handlers"
	"task/pkg/middlewares"
	"task/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	repo := repositories.NewRepository(config.DB)
	h := handlers.NewHandler(repo)

	r.HandleFunc("/transaction", middlewares.Transaction(h.Transaction)).Methods(http.MethodPost)
}
