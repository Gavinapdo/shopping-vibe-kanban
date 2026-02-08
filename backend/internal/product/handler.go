package product

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/products", h.listProducts)
	group.GET("/products/:id", h.getProduct)
	group.POST("/products", h.createProduct)
	group.PUT("/products/:id", h.updateProduct)
	group.DELETE("/products/:id", h.deleteProduct)
}

func (h *Handler) listProducts(c *gin.Context) {
	products := h.service.ListProducts()
	c.JSON(http.StatusOK, gin.H{"items": products})
}

func (h *Handler) getProduct(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	product, err := h.service.GetProduct(id)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) createProduct(c *gin.Context) {
	var input CreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数格式错误"})
		return
	}

	product, err := h.service.CreateProduct(input)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *Handler) updateProduct(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var input UpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数格式错误"})
		return
	}

	product, err := h.service.UpdateProduct(id, input)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) deleteProduct(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品ID不合法"})
		return 0, false
	}
	return id, true
}

func (h *Handler) handleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrProductNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, ErrInvalidInput):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
	}
}
