package main

import (
	"log"

	"github.com/delgoden/shop-wares-service/internal/api/handlers"
	"github.com/delgoden/shop-wares-service/internal/api/server"
	"github.com/delgoden/shop-wares-service/internal/api/service"
	"github.com/delgoden/shop-wares-service/pkg/storage"
)

func main() {
	coonString := "postgres://postgres:qwery@localhost:5432/postgres"
	pool, err := storage.NewPosgresDB(coonString)
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorage(pool)
	services := service.NewServise(storage)
	h := handlers.NewHandler(services)

	srv := new(server.Server)

	if err := srv.Run(":8002", h.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
