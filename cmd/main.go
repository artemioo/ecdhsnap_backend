package main

import (
	"log"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/database"
	"github.com/artemioo/ecdhsnap_backend/internal/handler"
	"github.com/artemioo/ecdhsnap_backend/internal/service"
)

func main() {
	db := database.NewDatabase()
	services := service.NewService(db)       // services depends on DB
	handlers := handler.NewHandler(services) // handlers depends on services

	srv := new(ecdhsnap.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
