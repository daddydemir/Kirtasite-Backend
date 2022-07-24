package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/daddydemir/kirtasiye-projesi/auth"
	"github.com/daddydemir/kirtasiye-projesi/cloud"
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
	"github.com/daddydemir/kirtasiye-projesi/validations"

	"github.com/gorilla/mux"
)

func GetFileByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	id, _ := strconv.Atoi(key)
	token := r.Header["Authorization"]
	if token == nil {
		json.NewEncoder(w).Encode(NotLoginMessage())
	} else {
		status, message := service.GetFileByUserIdService(token[0], id)
		if status {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(repositories.FileByUserId(key))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(message)
		}
	}
}

func FileById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]

	token := r.Header["Authorization"]
	if token == nil {
		json.NewEncoder(w).Encode(NotLoginMessage())
	} else {
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

}

func FileAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	file := models.File{}
	file.Id, _ = strconv.Atoi(r.FormValue("id"))
	file.UserId, _ = strconv.Atoi(r.FormValue("user_id"))
	if r.FormValue("private") == "true" {
		file.Private = true
	} else {
		file.Private = false
	}
	token := r.Header["Authorization"]
	if token == nil {
		json.NewEncoder(w).Encode(NotLoginMessage())
	} else {
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

}

// bunu kaldırdık sanki
func FileDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	//repositories.FileDelete(key)
	json.NewEncoder(w).Encode(map[string]string{"message": key + " deleted."})
}
