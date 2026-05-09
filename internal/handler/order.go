package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pintureria-lincoln/internal/model"
	"pintureria-lincoln/internal/repository"
)

type OrderHandler struct {
	repo        *repository.OrderRepository
	productRepo *repository.ProductRepository
}

func NewOrderHandler(repo *repository.OrderRepository, productRepo *repository.ProductRepository) *OrderHandler {
	return &OrderHandler{repo: repo, productRepo: productRepo}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req model.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var total float64
	for i, item := range req.Items {
		product, err := h.productRepo.GetByID(item.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("producto %d no encontrado", item.ProductID)})
			return
		}
		req.Items[i].UnitPrice = product.Price
		req.Items[i].Name = product.Name
		total += product.Price * float64(item.Quantity)
	}

	order, err := h.repo.Create(req, total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}
