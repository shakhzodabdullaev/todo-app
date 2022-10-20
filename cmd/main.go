package main

import (
	//"log"

	"github.com/shakhzodabdullaev/todo-app"
	"github.com/shakhzodabdullaev/todo-app/pkg/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("localhost:8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error : %s", err.Error())
	}
}
