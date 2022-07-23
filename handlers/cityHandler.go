package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/gorilla/mux"
)

func GetCityById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	city := repositories.GetCityById(key)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(city)
}
