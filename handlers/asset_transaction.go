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

	//TODO get asset by id
	asset, err := h.repo.GetAssetByID(int(request.DestAssetID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Code: http.StatusNotFound, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//TODO check address
	if asset.Address != transAddress {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "invalid address"}
		json.NewEncoder(w).Encode(response)
		return
	}

	total := request.Amount + request.GasFee

	//TODO check if src asset balance is greater or equal to request amount

	asset, err = h.repo.GetAssetByID(int(request.SrcAssetID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Code: http.StatusNotFound, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if asset.Balance >= request.Amount {
		//TODO if true, update src asset balance = current balance - request amount
		assetBalance := asset.Balance - total
		err := h.repo.UpdateAssetBalance(asset, assetBalance, false)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		//Todo update destination asset balance = current balance + request total
		destAsset, err := h.repo.GetAssetByID(int(request.DestAssetID))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			response := dto.ErrorResult{Code: http.StatusNotFound, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		destAssetBalance := request.Amount

		err = h.repo.UpdateAssetBalance(destAsset, destAssetBalance, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "your balance is too low"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//Todo create new transaction
	newTransaction := models.Transaction{
		SrcWalletID:  request.SrcWalletID,
		SrcAssetID:   request.SrcAssetID,
		DestWalletID: request.DestWalletID,
		DestAssetID:  request.DestAssetID,
		Amount:       request.Amount,
		GasFee:       request.GasFee,
		Total:        total,
	}

	err = h.repo.CreateTransaction(newTransaction)
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
