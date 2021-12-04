package service

import (
	"context"

	"github.com/delgoden/shop-wares-service/internal/models"
	"github.com/delgoden/shop-wares-service/pkg/storage"
)

type WareService struct {
	storage storage.Ware
}

func NewWareStorage(storage storage.Ware) *WareService {
	return &WareService{storage: storage}
}

func (s *WareService) Create(ctx context.Context, category models.Category, ware models.Ware) (models.Ware, error) {
	return s.storage.Create(ctx, category, ware)
}

func (s *WareService) GetWareByID(ctx context.Context, ID int) (models.Ware, error) {
	return s.storage.GetWareByID(ctx, ID)
}

func (s *WareService) GetAllWaresInCategory(ctx context.Context, category models.Category) ([]models.Ware, error) {
	return s.storage.GetAllWaresInCategory(ctx, category)
}

func (s *WareService) Update(ctx context.Context, ware models.Ware) (models.Ware, error) {
	return s.storage.Update(ctx, ware)
}

func (s *WareService) Delete(ctx context.Context, ID int) error {
	return s.storage.Delete(ctx, ID)
}
