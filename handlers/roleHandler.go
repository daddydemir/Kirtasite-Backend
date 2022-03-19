package handlers

import (
	"demir/repositories"
	"encoding/json"
	"net/http"
)

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.RoleGetAll())

}
