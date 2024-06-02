package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/gin-gonic/gin"
)

func GetDiamondPrice(c *gin.Context) {
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	var diamondPrice []entiy.DiamondPrice
	db.Model(entiy.DiamondPrice{}).Find(&diamondPrice)
	c.JSON(200, diamondPrice)
}
