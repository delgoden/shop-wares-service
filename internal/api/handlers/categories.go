package handlers

import (
	"net/http"

	"github.com/delgoden/shop-wares-service/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.String(http.StatusBadRequest, "bad request : %v", err)
	}

	category, err := h.services.Category.Create(c, category)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.services.GetAll(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
	}
	c.JSON(http.StatusOK, categories)
}

func (h *Handler) updateCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.String(http.StatusBadRequest, "bad request : %v", err)
	}

	category, err := h.services.Category.Update(c, category)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) deleteCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.String(http.StatusBadRequest, "bad request : %v", err)
	}

	err := h.services.Category.Delete(c, category.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
	}

	c.JSON(http.StatusOK, category)
}
