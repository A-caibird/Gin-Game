package main

import (
	"Game/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// Register a global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Auth)
	r.Use(middleware.Cors)
	//download
	r.Group("/download/")
	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "%#v", 20+10)
		context.Next()
	}, func(context *gin.Context) {
		context.String(200, "%s", "fasfasfdas")
	})
	r.Run(":8080")
}
