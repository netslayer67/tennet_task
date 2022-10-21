package handlers

import (
	"net/http"
	"task/repositories"
)

type handlerAsset struct {
	AssetRepository repositories.AssetRepository
}

func HandlerAsset(AssetRepository repositories.AssetRepository) *handlerAsset {
	return &handlerAsset{AssetRepository}
}

func (h *handlerAsset) CreateAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	request := assetdto.AssetRequest{
		Title:       r.FormValue("titlelink"),
		Description: r.FormValue("descriptionlink"),
		Image:       r.FormValue("file"),
		Template:    r.FormValue("template"),
	}
}
