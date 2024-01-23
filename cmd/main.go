package main

import (
	"log"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/database"
	"github.com/artemioo/ecdhsnap_backend/internal/handler"
	"github.com/artemioo/ecdhsnap_backend/internal/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db := database.NewDatabase()
	services := service.NewService(db)       // services depends on DB
	handlers := handler.NewHandler(services) // handlers depends on services

	srv := new(ecdhsnap.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs") // dir
	viper.SetConfigName("config")  // filename
	return viper.ReadInConfig()
}
