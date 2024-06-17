package main

import (
	"os"

	"github.com/jfortez/gogagg/api"
	"github.com/jfortez/gogagg/db"
	"github.com/jfortez/gogagg/services"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	address := os.Getenv("PORT")
	storage := db.New()
	defer storage.Close()

	hub := services.NewHub()
	go hub.Run()

	service := api.NewService(address, storage.DB, hub)
	service.Run()

}
