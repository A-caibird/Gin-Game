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
	//r.Use(middleware.Auth)
	// remote
	r.Group("/download/")
	r.POST("/login/:method", handler.Login)
	r.POST("/signup", handler.SignUp)
	r.POST("/sms/:usage", handler.SendCode)
	r.GET("/login_histories/:id", handler.QueryLh)
	// start up
	r.Run(":8000")
}
