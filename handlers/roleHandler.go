package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/daddydemir/kirtasiye-projesi/repositories"
)

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.RoleGetAll())

}
