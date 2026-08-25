package handlers

import (
	"net/http"
	"randomNasa/internal/aggregator"
	"randomNasa/internal/services"
	"randomNasa/internal/utils"
)

type RandomNasaHandler struct {
	service *services.RandomNasaService
}

func NewRandomNasaHandler()  *RandomNasaHandler{
	return &RandomNasaHandler{
		service: services.NewRandomNasaService(),
	}
}

func (h *RandomNasaHandler) GetRandomNasaUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	/* El servicio */
	randomNasa, err := aggregator.Aggregate(ctx, h.service)

	if err != nil {
		/* El utils Error */
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	/* El utils Ok */
	utils.ResponsseSuccess(w, http.StatusOK, randomNasa)
}