package main

import (
	"log"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
)

func main() {
	srv := new(ecdhsnap.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
