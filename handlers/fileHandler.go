package handlers

import (
	"demir/cloud"
	"demir/models"
	"demir/repositories"
	"demir/validations"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetFileByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.FileByUserId(key))
}

func FileAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	content := r.ContentLength
	if content >= 1024*1024*5 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Dosya boyutu 5 MB'ı geçmemelidir"})
		return
	}

	myfile, header, _ := r.FormFile("file")
	defer myfile.Close()

	file := models.File{}
	file.Id, _ = strconv.Atoi(r.FormValue("id"))
	file.UserId, _ = strconv.Atoi(r.FormValue("user_id"))
	if r.FormValue("private") == "true" {
		file.Private = true
	} else {
		file.Private = false
	}

	data, err := validations.FileValidation(file)
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
	json.NewEncoder(w).Encode(data)
}

func FileDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	repositories.FileDelete(key)
	json.NewEncoder(w).Encode("Deleted")
}
