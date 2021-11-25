package service

import (
	"context"

	"github.com/delgoden/shop-wares-service/internal/models"
	"github.com/delgoden/shop-wares-service/pkg/storage"
)

type Category interface {
	Create(ctx context.Context, category models.Category) (models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
	Update(ctx context.Context, category models.Category) (models.Category, error)
	Delete(ctx context.Context, ID int) error
}

type Service struct {
	Category
}

func NewServise(storage *storage.Storage) *Service {
	return &Service{
		Category: NewCategoriService(storage.Category),
	}
}
