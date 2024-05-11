package main

import (
	"Game/handler"
	"Game/middleware"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r := gin.New()
	// log
	logFilePath := "./log/gin.log"
	file, err := os.Create(logFilePath)
	if err != nil {
		color.Red("%#v", err)
	}
	defer file.Close()
	// writer
	multiWriter := io.MultiWriter(file, os.Stdout)
	// Register a global middleware
	r.Use(gin.LoggerWithWriter(multiWriter))
	r.Use(gin.Recovery())
	r.Use(middleware.Cors)
	r.Use(middleware.Auth)
	//download
	r.Group("/download/")
	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "%#v", 20+10)
		context.Next()
	}, func(context *gin.Context) {
		context.String(200, "%s", "fasfasfdas")
	})
	r.POST("/login", handler.Login)
	r.Run(":8000")
}
