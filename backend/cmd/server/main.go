package main

import (
	"log"

	"shopping-vibe-kanban/backend/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
