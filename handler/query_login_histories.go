package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/gin-gonic/gin"
	"net/http"
)

// QueryLh query user's  login histories
func QueryLh(c *gin.Context) {
	//
	id := c.Param("id")
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var lh []entiy.LoginHistory
	db.Where("user_id = ?", id).Find(&lh)
	c.JSON(http.StatusOK, lh)
}
