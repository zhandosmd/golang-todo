package main

import (
	"log"

	todo "github.com/zhandosmd/golang-todo"
	"github.com/zhandosmd/golang-todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8888", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
