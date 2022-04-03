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
	"time"

	"github.com/gorilla/mux"
)

func OrderByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, _ := strconv.Atoi(key)
	token := r.Header["Authorization"]
	status, message := service.OrderByUserIdService(token[0], id)
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.OrderByUserId(key))
	} else {
		if message["message"] == "Yetkisiz kullancı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func OrderByIdForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	order := repositories.OrderById(key)
	token := r.Header["Authorization"]
	status, message := service.OrderByIdServiceForUser(token[0], order)
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	} else {
		if message["message"] == "Yetkisiz kullancı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func OrderByIdForStationery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	order := repositories.OrderById(key)
	token := r.Header["Authorization"]
	status, message := service.OrderByIdServiceForStationery(token[0], order)
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	} else {
		if message["message"] == "Yetkisiz kullancı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func OrderByStationerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, _ := strconv.Atoi(key)
	token := r.Header["Authorization"]
	status, message := service.OrderByStationerIdService(token[0], id)
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.OrderByStationerId(key))
	} else {
		if message["message"] == "Yetkisiz kullancı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func OrderAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &order)
	token := r.Header["Authorization"]
	status, message := service.OrderAddService(token[0], order)
	if status {
		_, err := validations.OrderyValidation(order)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
			order.DeliveryDate = time.Now()
			repositories.OrderAdd(order)
		}
		json.NewEncoder(w).Encode(message)
	} else {
		if message["message"] == "Yetkisiz kullancı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}
