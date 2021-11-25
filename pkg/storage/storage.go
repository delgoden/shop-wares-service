package storage

import (
	"context"

	"github.com/delgoden/shop-wares-service/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Category interface {
	Create(ctx context.Context, category models.Category) (models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
	Update(ctx context.Context, category models.Category) (models.Category, error)
	Delete(ctx context.Context, ID int) error
}

type Storage struct {
	Category
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{
		Category: NewCategoryStorage(pool),
	}
}
