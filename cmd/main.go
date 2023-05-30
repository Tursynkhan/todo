package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/tursynkhan/todo-app"
	"github.com/tursynkhan/todo-app/pkg/handler"
	"github.com/tursynkhan/todo-app/pkg/repository"
	"github.com/tursynkhan/todo-app/pkg/service"
)

func main() {

	if err := intiConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func intiConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
