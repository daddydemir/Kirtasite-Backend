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

func GetFileByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.FileByUserId(key))
}

func FileAdd(w http.ResponseWriter, r *http.Request) {
	var file models.File
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &file)
	w.Header().Set("Content-Type", "application/json")
	data, err := validations.FileValidation(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
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
