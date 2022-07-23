package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
	"github.com/daddydemir/kirtasiye-projesi/validations"

	"github.com/gorilla/mux"
)

func GetCommentByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	id, _ := strconv.Atoi(key)
	token := r.Header["Authorization"]
	status, message := service.GetCommentByUserIdService(token[0], id)
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.CommentByUserId(key))
	} else {
		if message["message"] == "Yetksisiz kullanıcı." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func CommentByStationeryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.CommentByStationeryId(key))

}

func CommentAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	var comment models.Comment
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &comment)

	token := r.Header["Authorization"]
	status, message := service.CommentAddService(token[0], comment)
	if status {
		_, err := validations.CommentValidation(comment)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
			comment.Date = time.Now()
			repositories.CommentAdd(comment)
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
