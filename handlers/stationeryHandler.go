package handlers

import (
	"demir/cloudinary"
	"demir/models"
	"demir/repositories"
	"demir/validations"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllStationery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.StationeryGetAll())

}
func GetStationeryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.StationeryGetById(key))
}
func UpdateStationery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var stationery models.Stationery
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &stationery)
	stationery.Id, _ = strconv.Atoi(key)

	repositories.StationeryUpdate(stationery)
	json.NewEncoder(w).Encode("Updated")
}

func DeleteStationery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	repositories.StationeryDelete(key)
	json.NewEncoder(w).Encode("Success")
}

func AddStationery(w http.ResponseWriter, r *http.Request) {
	var stationery models.Stationery
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &stationery)
	w.Header().Set("Content-Type", "application/json")
	stationery.ImagePath = "https://res.cloudinary.com/dpdlwo6vi/image/upload/v1647981518/1647981517.png"
	data, err := validations.StationeryValidation(stationery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
		repositories.StationeryAdd(stationery)
	}
	json.NewEncoder(w).Encode(data)
}

func UpdateSImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
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
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Başarıyla güncellendi"})
}
