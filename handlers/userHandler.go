package handlers

import (
	"demir/models"
	"demir/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	users := repositories.UserGetAll()
	json.NewEncoder(w).Encode(users)

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	fmt.Println("DATA : ", user.Username)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	repositories.AddUser(user)
	json.NewEncoder(w).Encode("success")
}
