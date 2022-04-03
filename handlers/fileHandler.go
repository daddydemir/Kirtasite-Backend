package handlers

import (
	"demir/auth"
	"demir/cloud"
	"demir/models"
	"demir/repositories"
	"demir/service"
	"demir/validations"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetFileByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, _ := strconv.Atoi(key)
	token := r.Header["Authorization"]
	status, message := service.GetFileByUserIdService(token[0], id)
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.FileByUserId(key))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
	}
}

func FileById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	token := r.Header["Authorization"]
	status, message := service.FileByIdService(token[0], key)
	if status {
		var files models.File
		var orders []models.Order
		Id := auth.TokenParser(token[0])
		stationeryId, _ := repositories.StationeryByName(Id)
		orders = repositories.OrderByStationerId(strconv.Itoa(stationeryId.Id))
		for _, s := range orders {
			if strconv.Itoa(s.FileId) == key {
				if s.StationeryId == stationeryId.Id {
					log.Println("FILE ID : ", s.FileId)
					files = repositories.FileById(strconv.Itoa(s.FileId))
				}
			}
		}

		if files.Id == 0 {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"message": "Yetkisiz kullanıcı"})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(files)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
	}
}

func FileAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file := models.File{}
	file.Id, _ = strconv.Atoi(r.FormValue("id"))
	file.UserId, _ = strconv.Atoi(r.FormValue("user_id"))
	if r.FormValue("private") == "true" {
		file.Private = true
	} else {
		file.Private = false
	}
	token := r.Header["Authorization"]
	status, message := service.FileAddService(token[0], file)
	if status {
		content := r.ContentLength
		if content >= 1024*1024*5 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Dosya boyutu 5 MB'ı geçmemelidir"})
			return
		}
		myfile, header, _ := r.FormFile("file")
		defer myfile.Close()
		_, err := validations.FileValidation(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			path, folder, err := cloud.UploadFile(myfile, header)
			if err != nil {
				fmt.Println("Servis hatası - ", err.Error())
				return
			}
			file.FilePath = path
			file.FolderId = folder
			w.WriteHeader(http.StatusCreated)
			file.CreatedDate = time.Now()
			repositories.FileAdd(file)
		}
		json.NewEncoder(w).Encode(message)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
	}
}

func FileDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//repositories.FileDelete(key)
	json.NewEncoder(w).Encode(map[string]string{"message": key + " deleted."})
}
