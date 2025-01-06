package main

import (
	"log"

	todo "github.com/ggleym/rest-golang"
	"github.com/ggleym/rest-golang/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running the server: %s", err.Error())
	}

}
