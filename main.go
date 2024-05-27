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
	var logFile *os.File
	if _, err := os.Stat(logFilePath); os.IsExist(err) {
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			os.Exit(2)
		} else {
			logFile = file
		}
	} else {
		color.Red("%#v", "fasfasdf")
		file, err := os.Create(logFilePath)
		if err != nil {
			color.Red("%#v", err)
			os.Exit(2)
		}
		logFile = file
	}
	defer logFile.Close()
	// writer
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	// Register a global middleware
	r.Use(gin.LoggerWithWriter(multiWriter))
	r.Use(gin.Recovery())
	r.Use(middleware.Cors)
	//r.Use(middleware.Auth)
	// remote
	download := r.Group("/download/")
	download.GET("/avatar/:id", handler.GetAvatar)
	//
	r.POST("/login/:method", handler.Login)
	r.POST("/signup", handler.SignUp)
	r.POST("/sms/:usage", handler.SendCode)
	r.POST("/email_code/:usage", handler.GetEmailCode)
	r.POST("/add_friend", handler.AddFriend)
	//
	r.GET("/login_histories/:id", handler.QueryLh)
	r.GET("/avatar/:id", handler.GetAvatar)
	r.GET("/friends/:id", handler.GetFriendInfo)
	//
	modify := r.Group("/modify_info")
	modify.PATCH("/name/:id", handler.ModifyName)
	modify.PATCH("/avatar/:id", handler.ModifyAvatar)
	modify.PATCH("/phone", handler.ModifyPhone)
	modify.PATCH("/email", handler.ModifyEmail)
	// start up
	r.Run(":8000")
}
