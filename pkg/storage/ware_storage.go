package storage

import (
	"context"

	"github.com/delgoden/shop-wares-service/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type WareStorage struct {
	pool *pgxpool.Pool
}

func NewWareStorage(pool *pgxpool.Pool) *WareStorage {
	return &WareStorage{pool: pool}
}

func (s *WareStorage) Create(ctx context.Context, category models.Category, ware models.Ware) (models.Ware, error) {
	return ware, nil
}

func (s *WareStorage) GetWareByID(ctx context.Context, ID int) (models.Ware, error) {
	var ware models.Ware
	return ware, nil
}

func (s *WareStorage) GetAllWaresInCategory(ctx context.Context, category models.Category) ([]models.Ware, error) {
	var wares []models.Ware
	return wares, nil
}

func (s *WareStorage) Update(ctx context.Context, ware models.Ware) (models.Ware, error) {
	return ware, nil
}

func (s *WareStorage) Delete(ctx context.Context, ID int) error {
	return nil
}
