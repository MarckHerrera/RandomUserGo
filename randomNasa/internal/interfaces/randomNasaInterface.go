package interfaces

import (
	"context"
	"randomNasa/internal/models"
)

type RandomNasaService interface {
	Search(ctx context.Context) (*models.RandomUser, error)
}