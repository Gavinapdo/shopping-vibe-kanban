package main

import (
	"log"
	"os"

	"shopping-vibe-kanban/backend/internal/server"
)

func main() {
	router := server.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
