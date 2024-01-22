package main

import (
	"log"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(ecdhsnap.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
