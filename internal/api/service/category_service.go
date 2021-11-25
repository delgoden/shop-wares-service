package service

import (
	"context"

	"github.com/delgoden/shop-wares-service/internal/models"
	"github.com/delgoden/shop-wares-service/pkg/storage"
)

type CategoriService struct {
	storage storage.Category
}

func NewCategoriService(storage storage.Category) *CategoriService {
	return &CategoriService{storage: storage}
}

func (s *CategoriService) Create(ctx context.Context, category models.Category) (models.Category, error) {
	return s.storage.Create(ctx, category)
}

func (s *CategoriService) GetAll(ctx context.Context) ([]models.Category, error) {
	return s.storage.GetAll(ctx)
}

func (s *CategoriService) Update(ctx context.Context, category models.Category) (models.Category, error) {
	return s.storage.Update(ctx, category)
}

func (s *CategoriService) Delete(ctx context.Context, ID int) error {
	return s.storage.Delete(ctx, ID)
}
