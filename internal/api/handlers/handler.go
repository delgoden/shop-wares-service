package handlers

import (
	"github.com/delgoden/shop-wares-service/internal/api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	{
		categories := api.Group("categories")
		{
			categories.POST("/", h.createCategory)
			categories.GET("/", h.getAllCategories)
			categories.PUT("/:id", h.updateCategory)
			categories.DELETE("/:id", h.deleteCategory)
		}

	}

	return router
}
