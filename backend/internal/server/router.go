package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-vibe-kanban/backend/internal/product"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	repo := product.NewInMemoryRepository(product.MockProducts())
	service := product.NewService(repo)
	handler := product.NewHandler(service)

	api := router.Group("/api/v1")
	handler.RegisterRoutes(api)

	return router
}
