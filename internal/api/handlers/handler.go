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
		categories := api.Group("/categories")
		{
			categories.POST("/", h.createCategory)
			categories.GET("/", h.getAllCategories)
			categories.PUT("/:id", h.updateCategory)
			categories.DELETE("/:id", h.deleteCategory)
		}

		wares := api.Group("/wares")
		{
			wares.POST("/", h.createWare)
			wares.GET("/category/:id", h.getAllWaresInCategory)
			wares.GET("/:id", h.getWareByID)
			wares.PUT("/:id", h.updateWare)
			wares.DELETE("/:id", h.deleteWare)
		}
	}

	return router
}
