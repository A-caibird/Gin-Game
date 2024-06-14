package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/gin-gonic/gin"
	"strconv"
)

func LogOut(c *gin.Context) {
	//
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		return
	}
	db.Model(entiy.User{
		ID: uint(id),
	}).Where("id = ?", id).Update("online", false)
	c.AbortWithStatus(200)
}
