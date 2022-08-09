package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"kirtasiteproject/cloudinary"
	"kirtasiteproject/models"
	"kirtasiteproject/repositories"
	"kirtasiteproject/service"

	"github.com/gorilla/mux"
)

func AllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	token := r.Header["Authorization"]
	if token == nil {
		// pass
		json.NewEncoder(w).Encode(NotLoginMessage())
	} else {
		status, message := service.AllCustomers(token[0])
		if status {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}

}

func GetCustomerByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	token := r.Header["Authorization"]
	if token == nil {
		json.NewEncoder(w).Encode(NotLoginMessage())
	} else {
		vars := mux.Vars(r)
		key := vars["id"]
		status, message := service.GetCustomerByUserId(token[0], key)
		if status {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	var user models.Customers
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	status, message := service.AddUserService(user)
	if status {
		user.UserData.ImagePath = "https://res.cloudinary.com/dpdlwo6vi/image/upload/v1647981518/1647981517.png"
		user.UserData.RoleId = 1
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(message)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
	}
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	user, _ := repositories.GetCustomerByUserId(key)

	token := r.Header["Authorization"]
	if token == nil {
		json.NewEncoder(w).Encode(NotLoginMessage())
	} else {
		status, message := service.DeleteCustomer(token[0], user.(models.Customers))
		if status {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}

		json.NewEncoder(w).Encode(message)
	}

}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	var user models.Customers
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	user.UserId, _ = strconv.Atoi(key)
	token := r.Header["Authorization"]
	if token == nil {
		json.NewEncoder(w).Encode(NotLoginMessage())
	} else {
		status, message := service.UpdateCustomer(token[0], user)
		if status {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	token := r.Header["Authorization"]
	if token == nil {
		json.NewEncoder(w).Encode(NotLoginMessage())
	} else {
		customer, _ := repositories.GetStationeryByUserId(key)
		status, message := service.UpdateImageService(token[0], customer.(models.Customers))
		if status {
			file, _, err := r.FormFile("file")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			}
			url, err := cloudinary.UploadToCloudinary(file)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			} else {
				repositories.UpdateCustomerImage(url, key)
				w.WriteHeader(http.StatusAccepted)
				json.NewEncoder(w).Encode(map[string]string{"message": "Başarıyla güncellendi"})
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(message)
		}
	}
}
