package handlers

import (
	"demir/models"
	"demir/repositories"
	"demir/validations"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetCommentByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.CommentByUserId(key))
}

func CommentByStationeryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.CommentByStationeryId(key))
}

func CommentAdd(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &comment)
	w.Header().Set("Content-Type", "application/json")
	data, err := validations.CommentValidation(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
		comment.Date = time.Now()
		repositories.CommentAdd(comment)
	}
	fmt.Println("SCORE : ", comment.Score)
	//fmt.Println("HATA : ", err.Error())
	json.NewEncoder(w).Encode(data)
}
