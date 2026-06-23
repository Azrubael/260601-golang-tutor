package main

import (
	"github.com/Azrubael/260601-golang-tutor/http_gin_rest1/controller"
	"github.com/Azrubael/260601-golang-tutor/http_gin_rest1/service"
	"github.com/gin-gonic/gin"
)

var(
	videoService service.VideoService = service.New()
	videoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{ "message": "Azrubael!" })
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
