package main

import (
	"demir/config"
	"demir/handlers"
	"demir/migrations"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server Started at http://localhost:1337/api/")
	config.GetConnect()
	migrations.MainMigration()

	server := &http.Server{
		Addr:    ":1337",
		Handler: handlers.MainRouting(),
	}
	server.ListenAndServe()
}
