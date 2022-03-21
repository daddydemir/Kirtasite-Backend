package handlers

import (
	"demir/config"
	"demir/models"
	"demir/repositories"
	"demir/validations"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	users := repositories.UserGetAll()
	json.NewEncoder(w).Encode(users)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(repositories.GetUserById(key))

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	w.Header().Set("Content-Type", "application/json")

	data, err := validations.UserValidation(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
		repositories.AddUser(user)
	}
	json.NewEncoder(w).Encode(data)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var user models.User

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	config.DB.Find(&user, "id = ?", key)

	repositories.DeleteUser(user)
	json.NewEncoder(w).Encode("Deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var user models.User
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	user.Id, _ = strconv.Atoi(key)

	repositories.UpdateUser(user)
	json.NewEncoder(w).Encode("Updated")

}
