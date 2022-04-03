package handlers

import (
	"demir/models"
	"demir/repositories"
	"demir/service"
	"demir/validations"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header["Authorization"]
	status, message := service.GetAllPricesService(token[0])
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.GetAllPrices())
	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func PriceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	token := r.Header["Authorization"]
	status, message := service.PriceByIdService(token[0])
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.PriceById(key))
	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}

}

func PriceAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
