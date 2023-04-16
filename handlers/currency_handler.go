package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wchopite/api-spark-v1/services"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	currency := params.Get("id")

	result := services.GetUser(currency)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
