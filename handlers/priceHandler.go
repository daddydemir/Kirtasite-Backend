package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
	"github.com/daddydemir/kirtasiye-projesi/validations"

	"github.com/gorilla/mux"
)

func GetAllPrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllPrices())

}

func PriceByStationeryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.PriceByStationeryId(key))
}

func PriceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.PriceById(key))

}

func PriceAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	var price models.Price
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &price)
	token := r.Header["Authorization"]
	status, message := service.PriceAddService(token[0], price)
	if status {
		_, err := validations.PriceValidation(price)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
			repositories.PriceAdd(price)
		}
		json.NewEncoder(w).Encode(message)
	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}

}

func PriceDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	token := r.Header["Authorization"]
	status, message := service.PriceDeleteService(token[0], repositories.PriceById(key))
	if status {
		w.WriteHeader(http.StatusOK)
		repositories.PriceDelete(key)
		json.NewEncoder(w).Encode(message)
	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func PriceUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]

	var price models.Price
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &price)
	price.Id, _ = strconv.Atoi(key)

	securityCode := repositories.PriceById(key).StationeryId
	price.StationeryId = securityCode

	token := r.Header["Authorization"]
	status, message := service.PriceUpdateService(token[0], price)
	if status {
		w.WriteHeader(http.StatusOK)
		repositories.PriceUpdate(price)
		json.NewEncoder(w).Encode(message)
	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}
