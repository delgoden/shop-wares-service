package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
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
