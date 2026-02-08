package app

import (
	"github.com/gin-gonic/gin"

	"shopping-vibe-kanban/backend/internal/transport/http/router"
)

func Run() error {
	engine := gin.Default()
	router.Register(engine)

	return engine.Run(":8080")
}
