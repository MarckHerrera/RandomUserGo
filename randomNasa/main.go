package main

import (
	"log"
	"net/http"
	"randomNasa/internal/config"
	"randomNasa/internal/handlers"
)

func main()  {
	configEnv := config.Load()

	Handler := handlers.NewRandomNasaHandler()
	mainRouter := goRouter(Handler)

	log.Printf("Escuchando en el puerto %+v", configEnv)
	log.Fatal(http.ListenAndServe(":8080",mainRouter))
}