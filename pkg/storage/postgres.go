package storage

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrInternal               = errors.New("internal error")
	ErrCategoryAlreadyExists  = errors.New("category already exists")
	ErrCategoryDoesNotExist   = errors.New("category does not exist")
	ErrCategoriesDoesNotExist = errors.New("categories does not exist")
	ErrWareAlreadyExists      = errors.New("ware already exists")
	ErrWareDoesNotExist       = errors.New("ware does not exist")
)

// type Config struct {
// 	Host     string
// 	Port     string
// 	Username string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

func NewPosgresDB(connString string) (*pgxpool.Pool, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	return pgxpool.Connect(ctx, connString)
}
