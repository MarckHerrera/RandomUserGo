package main

import (
	"github.com/gorilla/mux"
	"randomusers/internal/handlers"
)

func goRouter(randomUserHandler *handlers.RandomUserHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/randomusers/", randomUserHandler.GetRandomUser).Methods("GET")

	return router
}