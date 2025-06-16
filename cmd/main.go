package main

import (
	"github.com/spf13/viper"
	"log"
	"todo_api/pkg/handlers"
	"todo_api/pkg/repositories"
	"todo_api/pkg/services"
	"todo_api/server"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal("Config wasn't loaded", err.Error())
	}

	repo := repositories.NewRepository()
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	// Initialize routes
	router := handler.InitRoutes()

	// Create and start server
	srv := server.New()
	log.Printf("Starting server on %s:%s...", viper.GetString("host"), viper.GetString("port"))

	if err := srv.Run(viper.GetString("port"), router); err != nil {
		log.Fatal("Server failed to start:", err.Error())
	}
}

func initConfig() error {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
