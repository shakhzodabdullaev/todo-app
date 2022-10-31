package main

import (
	//"log"

	"log"

	"github.com/shakhzodabdullaev/todo-app"
	"github.com/shakhzodabdullaev/todo-app/pkg/handler"
	"github.com/shakhzodabdullaev/todo-app/pkg/repository"
	"github.com/shakhzodabdullaev/todo-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("Error initializing configs: ", err.Error())
	}
	//handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error : %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
