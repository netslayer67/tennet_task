package handlers

import (
	"encoding/json"
	"net/http"
	dto "task/dto/result"
	transactiondto "task/dto/transaction"
	"task/models"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

func (h *handler) Transaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get user from token
	transactionInfo := r.Context().Value("transactionInfo").(jwt.MapClaims)
	transAddress := transactionInfo["address"]

	request := new(transactiondto.CreateTransaction)
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

	asset, err := h.repo.GetAssetByID(int(request.DestAssetID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	} else if asset.Address == transAddress {

		newTransaction := models.Transaction{
			SrcWalletID:  request.SrcWalletID,
			SrcAssetID:   request.SrcAssetID,
			DestWalletID: request.DestWalletID,
			DestAssetID:  request.DestWalletID,
			Amount:       request.Amount,
			GasFee:       request.GasFee,
			Total:        request.Amount + request.GasFee,
		}

		err := h.repo.CreateTransaction(newTransaction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: newTransaction}
		json.NewEncoder(w).Encode(response)
	}
}
