package interfaces

import (
	"context"

	"randomusers/internal/models"
)

type RandomUserService interface {
	Search(ctx context.Context) (*models.RandomUser, error)
}