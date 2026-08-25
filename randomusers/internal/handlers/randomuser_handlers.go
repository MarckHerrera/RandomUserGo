package handlers

import (
	"net/http"
	"context"
	"time"

	"randomusers/internal/services"
	"randomusers/internal/utils"
	"randomusers/internal/aggregator"
)

type RandomUserHandler struct {
	service *services.RandomUserService
}

func NewRandomUserHandler() *RandomUserHandler {
	return &RandomUserHandler{
		service: services.NewRandomUserService(),
	}
}

func (h *RandomUserHandler) GetRandomUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()
	
	randomUser, err := aggregator.Aggregate(ctx, h.service)


	if err != nil {
		utils.ResponseHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseHttpSuccess(w, http.StatusOK, randomUser)
}

