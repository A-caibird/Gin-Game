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
	download := r.Group("/download/")
	download.GET("/avatar/:id", handler.GetAvatar)
	r.POST("/login/:method", handler.Login)
	r.POST("/signup", handler.SignUp)
	r.POST("/sms/:usage", handler.SendCode)
	r.GET("/login_histories/:id", handler.QueryLh)
	r.GET("/avatar/:id", handler.GetAvatar)
	modify := r.Group("/modify_info")
	modify.PATCH("/name/:id", handler.ModifyName)
	modify.PATCH("/avatar/:id", handler.ModifyAvatar)
	// start up
	r.Run(":8000")
}
