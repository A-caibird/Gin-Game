package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetNameById(c *gin.Context) {
	//
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	var user entiy.User
	db.Where(entiy.User{
		ID: uint(Id),
	}).First(&user)
	c.JSON(200, user.Name)
}
