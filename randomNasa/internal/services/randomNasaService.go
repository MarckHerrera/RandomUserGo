package services

import (
	"context"
	"encoding/json"
	"net/http"
	"randomNasa/internal/models"
)

type RandomNasaService struct {
	
}

func NewRandomNasaService() *RandomNasaService{
	return &RandomNasaService{}
}

func (s *RandomNasaService) Search(ctx context.Context) (*models.RandomUser, error){
	url := "https://randomuser.me/api/"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result models.RandomUser

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}
	
	return &result, nil
}