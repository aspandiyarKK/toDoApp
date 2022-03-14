package main

import (
	todo_app "github.com/aspandiyarK"
	"github.com/aspandiyarK/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("erroer occured :%s", err.Error())
	}
}
