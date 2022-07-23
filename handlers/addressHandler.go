package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/gorilla/mux"
)

func GetAddressById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	address := repositories.GetAddressById(key)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

func GetAddressByCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	address := repositories.GetAddressByCity(key)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

func GetAddressByDistrict(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	address := repositories.GetAddressByDistrict(key)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}
