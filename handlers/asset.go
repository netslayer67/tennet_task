package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	assetdto "task/dto/asset"
	dto "task/dto/result"
	"task/models"
	jwtToken "task/pkg/jwt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func (h *handler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(assetdto.CreateAsset)
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

	// Generate Token
	claims := jwt.MapClaims{}
	claims["address"] = request.Address
	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(err)
		fmt.Println("Address is invalid")
		return
	}

	newAsset := models.Asset{
		WalletID: request.WalletID,
		Name:     request.Name,
		Symbol:   request.Symbol,
		Network:  request.Network,
		Balance:  request.Balance,
		Address:  request.Address,
		Token:    token,
	}

	err = h.repo.CreateAsset(newAsset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: "Success Create Asset"}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) FindAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	assets, err := h.repo.FindAsset()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: assets}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")

	assetId, err := extractAssetId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	request := new(assetdto.UpdateAsset)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := h.repo.UpdateAsset(assetId, request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")

	assetId, err := extractAssetId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := h.repo.DeleteAsset(assetId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func extractAssetId(r *http.Request) (int, error) {
	urlParams := mux.Vars(r)
	assetId, ok := urlParams["id"]
	if !ok {
		return 0, errors.New("Asset ID Should be passed") //return htpt bad request
	}

	parsedAssetId, err := strconv.ParseInt(assetId, 10, 16)
	if err != nil {
		return 0, err
	}
	return int(parsedAssetId), nil
}
