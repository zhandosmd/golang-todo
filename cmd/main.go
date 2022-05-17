package main

import (
	"log"

	todo "github.com/zhandosmd/golang-todo"
	"github.com/zhandosmd/golang-todo/pkg/handler"
	"github.com/zhandosmd/golang-todo/pkg/repository"
	"github.com/zhandosmd/golang-todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8888", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
