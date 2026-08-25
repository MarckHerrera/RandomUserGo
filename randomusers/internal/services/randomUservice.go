package services

import (
	"context"
	"net/http"
	"encoding/json"

	"randomusers/internal/models"
)

type RandomUserService struct {
}

func NewRandomUserService() *RandomUserService {
	return &RandomUserService{}
}

func (s *RandomUserService) Search(ctx context.Context) (*models.RandomUser, error) {
	url := "https://randomuser.me/api/?results=500" 

	/* Hago la peticion sin mandarla aun */
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	/* Envio la peticion */
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	/* El defer cierra la respuesta */
	defer resp.Body.Close()

	/* Decodifico el json a struct */
	var result models.RandomUser

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
