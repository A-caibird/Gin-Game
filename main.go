package main

import (
	"Game/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	//
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
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
