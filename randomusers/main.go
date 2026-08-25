package main

import (
	"log"
	"net/http"

	"randomusers/internal/config"
	"randomusers/internal/handlers"
)

func main() {
	configEnv := config.Load()

	Handler := handlers.NewRandomUserHandler()
	mainRouter := goRouter(Handler)

	log.Printf("Listening on port %+v", configEnv)

	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}