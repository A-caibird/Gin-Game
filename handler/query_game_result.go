package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/gin-gonic/gin"
)

func QueryGameResult(c *gin.Context) {
	// id
	id := c.Param("id")
	//
	db, err := mysql.InitOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	var gameResult entiy.GameResult
	db.Where("user_id = ?", id).First(&gameResult)
	c.JSON(200, gameResult)
}
