package main

import (
	"fmt"
	"net/http"

	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/handlers"
	"github.com/daddydemir/kirtasiye-projesi/migrations"
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
