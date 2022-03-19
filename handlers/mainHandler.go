package handlers

import "github.com/gorilla/mux"

func MainRouting() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	// users
	r.HandleFunc("/api/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/api/users/add", AddUser).Methods("POST")

	// roles
	r.HandleFunc("/api/roles", GetAllRoles).Methods("GET")

	// stationeries
	r.HandleFunc("/api/stationery/", GetAllStationery).Methods("GET")

	// files
	r.HandleFunc("/api/files/", GetFileByUserId).Methods("GET")

	// prices
	r.HandleFunc("/api/prices/", GetAllPrices).Methods("GET")

	// orders
	r.HandleFunc("/api/orders/", GetOrderByUserId).Methods("GET")

	// comments
	r.HandleFunc("/api/comments/", GetCommentByUserId).Methods("GET")

	return r
}
