package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/daddydemir/kirtasiye-projesi/auth"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	/* 	fmt.Println("REQUEST : ", r)
	   	w.Header().Set("Content-Type", "application/json")
	   	//w.Header().Set("Access-Control-Allow-Origin", "*")
	   	w.Header().Add("Access-Control-Allow-Headers", "Content-Type , Authorization , X-Requested-With")
	   	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") */
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

func EnableCors(w *http.ResponseWriter) {
	header := (*w).Header()
	header.Add("Access-Control-Allow-Methods", "DELETE,POST,GET,OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type , Authorization, X-Requested-With")

}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {}
