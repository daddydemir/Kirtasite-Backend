package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/daddydemir/kirtasiye-projesi/auth"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)

	resp, status := auth.LoginToSystem(user.Username, user.Password)
	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(resp)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {}
