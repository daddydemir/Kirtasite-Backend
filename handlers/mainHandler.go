package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func MainRouting() http.Handler {

	r := mux.NewRouter().StrictSlash(true)
	// users
	r.HandleFunc("/api/users", AllCustomers).Methods("GET")
	r.HandleFunc("/api/user", AddUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", GetCustomerByUserId).Methods("GET")
	r.HandleFunc("/api/user/{id}", DeleteCustomer).Methods("DELETE")
	r.HandleFunc("/api/users/{id}", UpdateCustomer).Methods("PUT")
	r.HandleFunc("/api/user/{id}", UpdateImage).Methods("PUT")
	/*
		// auth
		r.HandleFunc("/api/auth/user/login", Login).Methods("POST")

		// roles
		r.HandleFunc("/api/roles", GetAllRoles).Methods("GET")
	*/
	// stationeries
	r.HandleFunc("/api/stationeries/", GetAllStationery).Methods("GET")
	r.HandleFunc("/api/stationery/{id}", GetStationeryById).Methods("GET")
	r.HandleFunc("/api/stationery/{id}", UpdateStationery).Methods("PUT")
	//r.HandleFunc("/api/stationery/{id}", DeleteStationery).Methods("DELETE")
	r.HandleFunc("/api/stationery/", AddStationery).Methods("POST")
	/*	r.HandleFunc("/api/stationery/{id}", UpdateSImage).Methods("PUT")

		// files
		r.HandleFunc("/api/files/{id}", GetFileByUserId).Methods("GET")
		r.HandleFunc("/api/files/{id}", FileDelete).Methods("DELETE")
		r.HandleFunc("/api/files/", FileAdd).Methods("POST")
		r.HandleFunc("/api/file/{id}", FileById).Methods("GET")

		// prices
		r.HandleFunc("/api/prices/", GetAllPrices).Methods("GET")
		r.HandleFunc("/api/prices/{id}", PriceById).Methods("GET")
		r.HandleFunc("/api/prices/stationery/{id}", PriceByStationeryId).Methods("GET")
		r.HandleFunc("/api/prices/{id}", PriceDelete).Methods("DELETE")
		r.HandleFunc("/api/prices/", PriceAdd).Methods("POST")
		r.HandleFunc("/api/prices/{id}", PriceUpdate).Methods("PUT")

		// orders
		r.HandleFunc("/api/orders/user/{id}", OrderByUserId).Methods("GET")
		r.HandleFunc("/api/orders/stationery/{id}", OrderByStationerId).Methods("GET")
		r.HandleFunc("/api/orders/", OrderAdd).Methods("POST")
		r.HandleFunc("/api/order/user/{id}", OrderByIdForUser).Methods("GET")
		r.HandleFunc("/api/order/stationery/{id}", OrderByIdForStationery).Methods("GET")

		// comments
		r.HandleFunc("/api/comments/user/{id}", GetCommentByUserId).Methods("GET")
		r.HandleFunc("/api/comments/stationery/{id}", CommentByStationeryId).Methods("GET")
		r.HandleFunc("/api/comments/", CommentAdd).Methods("POST")

		// address
		r.HandleFunc("/api/locate/{id}", GetAddressById).Methods("GET")
		r.HandleFunc("/api/city/{id}", GetAddressByCity).Methods("GET")
		r.HandleFunc("/api/district/{id}", GetAddressByDistrict).Methods("GET")

		// city
		r.HandleFunc("/api/cities/{id}", GetCityById).Methods("GET")

		// district

		r.HandleFunc("/api/districts/{id}", GetDistrictById).Methods("GET")
		r.HandleFunc("/api/district-city/{id}", GetDistrictByCity).Methods("GET")*/

	handler := cors.AllowAll().Handler(r)
	return handler
}
