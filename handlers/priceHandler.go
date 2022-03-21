package handlers

import (
	"demir/models"
	"demir/repositories"
	"demir/validations"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllPrices())
}

func PriceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.PriceById(key))
}

func PriceAdd(w http.ResponseWriter, r *http.Request) {
	var price models.Price
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &price)
	w.Header().Set("Content-Type", "application/json")
	data, err := validations.PriceValidation(price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
		repositories.PriceAdd(price)
	}
	json.NewEncoder(w).Encode(data)
}

func PriceDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	repositories.PriceDelete(key)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func PriceUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var price models.Price
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &price)
	price.Id, _ = strconv.Atoi(key)

	repositories.PriceUpdate(price)
	json.NewEncoder(w).Encode("Updated")
}
