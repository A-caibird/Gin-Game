package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ModifyName(c *gin.Context) {
	//
	id := c.Param("id")
	var body struct{ Name string }
	if err := c.BindJSON(&body); err != nil {
		return
	}
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	user := entiy.User{ID: uint(idInt)}
	res := db.First(&user)
	if res.RowsAffected == 1 {
		db.Model(&user).Update("name", body.Name)
		c.AbortWithStatus(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusNotFound)
}
