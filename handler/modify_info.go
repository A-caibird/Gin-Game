package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"Game/tools"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
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

func ModifyAvatar(c *gin.Context) {
	// parse user id
	id := c.Param("id")
	v, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userId := uint(v)
	// check user if exits
	db, err := mysql.NewOrmDb()
	res := db.Model(&entiy.User{ID: userId}).First(&entiy.User{}).RowsAffected
	if res != 1 {
		color.Red("%d", userId)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	//
	file, err := c.FormFile("avatar")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	avatarPath := path.Join(tools.Conf.RootPath.Path, "/public/avatar/", id+".png")
	// check avatar if exits,if exits remove origin avatar file
	if _, err := os.Stat(avatarPath); err == nil {
		if err := os.Remove(avatarPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete existing avatar file"})
			return
		}
	}
	// save avatar file
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar file"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Avatar modified successfully"})
}
