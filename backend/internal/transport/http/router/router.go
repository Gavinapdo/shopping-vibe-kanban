package router

import "github.com/gin-gonic/gin"

func Register(engine *gin.Engine) {
	engine.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}
