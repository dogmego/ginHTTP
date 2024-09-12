package main

import (
	"log"
	"maxHTTP"
	"maxHTTP/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(maxHTTP.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("failed to start server, error %s", err.Error())
	}
}
