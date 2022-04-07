package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/daddydemir/kirtasiye-projesi/cloudinary"
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
	"github.com/daddydemir/kirtasiye-projesi/validations"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	token := r.Header["Authorization"]
	status, message := service.GetAllUserService(token[0])
	if status {
		w.WriteHeader(http.StatusOK)
		users := repositories.UserGetAll()
		json.NewEncoder(w).Encode(users)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	token := r.Header["Authorization"]
	status, message := service.GetUserByIdService(token[0])
	if status {
		vars := mux.Vars(r)
		key := vars["id"]
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(repositories.GetUserById(key))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
	}

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	var user models.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	status, message := service.AddUserService(user)
	if status {
		user.ImagePath = "https://res.cloudinary.com/dpdlwo6vi/image/upload/v1647981518/1647981517.png"
		data, err := validations.UserValidation(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
			repositories.AddUser(user)
		}
		json.NewEncoder(w).Encode(data)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	user := repositories.GetUserById(key)

	token := r.Header["Authorization"]
	status, message := service.DeleteUserService(token[0], user)
	if status {
		repositories.DeleteUser(user)
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(message)
	} else {
		if message["message"] == "Sadece kendiniz üzerinde işlem yapabilirsiniz." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	var user models.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	user.Id, _ = strconv.Atoi(key)
	token := r.Header["Authorization"]
	status, message := service.UpdateUserService(token[0], user)
	if status {
		w.WriteHeader(http.StatusOK)
		repositories.UpdateUser(user)
		json.NewEncoder(w).Encode(message)
	} else {
		if message["message"] == "Sadece kendi bilgilerinizi güncelleyebilirsiniz." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}

}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	vars := mux.Vars(r)
	key := vars["id"]
	token := r.Header["Authorization"]
	status, message := service.UpdateImageService(token[0], repositories.GetUserById(key))
	if status {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		}
		url, err := cloudinary.UploadToCloudinary(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		} else {
			repositories.UpdateImage(url, key)
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(map[string]string{"message": "Başarıyla güncellendi"})
		}

	} else {
		if message["message"] == "Sadece kendi profil resminizi güncelleyebilirsiniz.." {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(message)
	}

}
