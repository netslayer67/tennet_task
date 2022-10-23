package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	AssetRoutes(r)
	WalletRoutes(r)
	TransactionRoutes(r)
}
