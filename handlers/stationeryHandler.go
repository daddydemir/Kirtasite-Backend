package handlers

import (
	"demir/cloudinary"
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

func GetAllStationery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header["Authorization"]
	status, message := service.GetAllStationeryService(token[0])
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.StationeryGetAll())
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
	}

}
func GetStationeryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header["Authorization"]
	status, message := service.GetStationeryByIdService(token[0])
	if status {
		vars := mux.Vars(r)
		key := vars["id"]
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.StationeryGetById(key))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
	}

}
func UpdateStationery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	var stationery models.Stationery
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &stationery)
	stationery.Id, _ = strconv.Atoi(key)
	token := r.Header["Authorization"]
	status, message := service.UpdateStationeryService(token[0], stationery)
	if status {
		repositories.StationeryUpdate(stationery)
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

func DeleteStationery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	token := r.Header["Authorization"]
	status, message := service.DeleteStationeryService(token[0], repositories.StationeryGetById(key))
	if status {
		w.WriteHeader(http.StatusNoContent)
		repositories.StationeryDelete(key)
	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func AddStationery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stationery models.Stationery
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &stationery)
	stationery.ImagePath = "https://res.cloudinary.com/dpdlwo6vi/image/upload/v1647981518/1647981517.png"
	status, message := service.AddStationeryService(stationery)
	if status {
		data, err := validations.StationeryValidation(stationery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
			repositories.StationeryAdd(stationery)
		}
		json.NewEncoder(w).Encode(data)
	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func UpdateSImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	token := r.Header["Authorization"]
	status, message := service.UpdateSImageService(token[0], repositories.StationeryGetById(key))
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
		}
		repositories.UpdateSImage(url, key)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Başarıyla güncellendi."})

	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}

}
