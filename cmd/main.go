package main

import (
	"log"
	"os"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/database"
	"github.com/artemioo/ecdhsnap_backend/internal/handler"
	"github.com/artemioo/ecdhsnap_backend/internal/service"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		log.Fatalf("error loading env variables %s", err.Error())
	}

	db_init, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize DB: %s", err.Error())
	}

	db := database.NewDatabase(db_init)
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
