package main

import (
	"randomNasa/internal/handlers"

	"github.com/gorilla/mux"
)

func goRouter(nasaHandler *handlers.RandomNasaHandler) *mux.Router{
	router := mux.NewRouter()

	/* Rutas de tal aqui */
	router.HandleFunc("/api/v1/randomusers/", nasaHandler.GetRandomNasaUser).Methods("Get")

	return router
}