package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	dto "task/dto/result"
	walletdto "task/dto/wallet"
	"task/models"

	"github.com/go-playground/validator/v10"
)

func (h *handler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(walletdto.CreateWallet)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	wallet := models.Wallet{
		Name: request.Name,
	}

	newWallet, err := h.repo.CreateWallet(wallet)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: newWallet}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) FindWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wallets, err := h.repo.FindWallet()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: wallets}
	json.NewEncoder(w).Encode(response)
}
