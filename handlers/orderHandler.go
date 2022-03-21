package handlers

import (
	"demir/models"
	"demir/repositories"
	"demir/validations"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func OrderByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.OrderByUserId(key))

}

func OrderByStationerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.OrderByStationerId(key))
}

func OrderAdd(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &order)
	w.Header().Set("Content-Type", "application/json")
	data, err := validations.OrderyValidation(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
		order.DeliveryDate = time.Now()
		repositories.OrderAdd(order)
	}
	json.NewEncoder(w).Encode(data)
}
