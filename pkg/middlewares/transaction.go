package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	dto "task/dto/result"
	jwtToken "task/pkg/jwt"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Transaction(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "Transaction Failed"}
			json.NewEncoder(w).Encode(response)
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "Transaction Failed"}
			json.NewEncoder(w).Encode(response)
			return
		}

		ctx := context.WithValue(r.Context(), "transactionInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
