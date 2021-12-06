package storage

import (
	"context"
	"log"

	"github.com/delgoden/shop-wares-service/internal/models"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CategoryStorage struct {
	pool *pgxpool.Pool
}

func NewCategoryStorage(pool *pgxpool.Pool) *CategoryStorage {
	return &CategoryStorage{pool: pool}
}

func (s *CategoryStorage) Create(ctx context.Context, category models.Category) (models.Category, error) {
	err := s.pool.QueryRow(ctx, `
	INSERT INTO categories (name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`, category.Title).Scan(&category.ID)
	if err == pgx.ErrNoRows {
		log.Println(err)
		return category, ErrCategoryAlreadyExists
	}
	if err != nil {
		log.Println(err)
		return category, err
	}
	return category, nil
}

func (s *CategoryStorage) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	rows, err := s.pool.Query(ctx, `SELECT id, title FROM categories`)
	if err == pgx.ErrNoRows {
		log.Println(err)
		return nil, ErrCategoriesDoesNotExist
	}

	defer rows.Close()

	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.ID, &category.Title); err != nil {
			log.Println(err)
		}

		categories = append(categories, category)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return categories, ErrInternal
	}

	return categories, nil
}

func (s *CategoryStorage) Update(ctx context.Context, category models.Category) (models.Category, error) {
	_, err := s.pool.Exec(ctx, `UPDATE categories SET name = $1 WHERE id = $2`, category.Title, category.ID)
	if err != nil {
		log.Println(err)
		return category, ErrInternal
	}
	return category, nil
}

func (s *CategoryStorage) Delete(ctx context.Context, ID int) error {
	_, err := s.pool.Exec(ctx, `DELETE FROM cart WHERE product_id = $1`, ID)
	if err != nil {
		return err
	}

	return nil
}
