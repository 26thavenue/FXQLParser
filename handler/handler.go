package handler

import (
	"encoding/json"
	"net/http"

	"github.com/26thavenue/FXQLParser/repository"
)

type CreateRequestBody struct {
	Input string `json:"input"`
}


func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody CreateRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := repository.Create(requestBody.Input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Transactions created successfully"))
}

func CheckCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	sourceCurrency := r.URL.Query().Get("source")
	destinationCurrency := r.URL.Query().Get("destination")

	if sourceCurrency == "" || destinationCurrency == "" {
		http.Error(w, "Missing source or destination query parameter", http.StatusBadRequest)
		return
	}

	exists := repository.CheckCurrencyPair(sourceCurrency, destinationCurrency)

	if exists {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Currency pair exists"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Currency pair does not exist"))
	}
}


