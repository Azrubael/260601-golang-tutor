package main

import (
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"github.com/Azrubael/260601-golang-tutor/http_gin_rest2/controller"
	"github.com/Azrubael/260601-golang-tutor/http_gin_rest2/middleware"
	"github.com/Azrubael/260601-golang-tutor/http_gin_rest2/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Logging to a file AND console at the same time.
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//server := gin.Default()
	server := gin.New()

	server.Use(gin.Recovery())

	server.Use(middleware.Logger())

	server.Use(gindump.Dump())

	server.Use(middleware.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
