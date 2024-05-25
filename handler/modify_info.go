package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	redis2 "Game/redis"
	"Game/tools"
	"context"
	"errors"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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
		c.AbortWithStatus(400)
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
		return
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

func ModifyPhone(c *gin.Context) {
	type Body struct {
		ID          int64  `json:"ID"`
		OriginPhone string `json:"OriginPhone"`
		DstPhone    string `json:"DstPhone"`
		Code        string `json:"Code"`
	}
	type resp struct {
		ID      uint
		Content string
	}
	var rby Body
	if err := c.BindJSON(&rby); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		color.Red("%#v", err.Error())
		return
	}
	rdb := redis2.NewRedisClient()
	if val, err := rdb.Get(context.Background(), rby.OriginPhone+"-"+"ModifyPhone").Result(); errors.Is(err, redis.Nil) {
		c.JSON(http.StatusUnauthorized, resp{
			ID:      0,
			Content: "code expiration",
		})
		return
	} else if val != rby.Code {
		c.JSON(http.StatusUnauthorized, resp{
			ID:      1,
			Content: "code error",
		})
		return
	}

	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	res := db.Model(&entiy.User{}).Where("id = ? ", rby.ID).Update("phone", rby.DstPhone).RowsAffected
	if res == 1 {
		c.Status(http.StatusOK)
		return
	}
	c.JSON(http.StatusInternalServerError, resp{
		ID:      1,
		Content: "update database error",
	})
}

func ModifyEmail(c *gin.Context) {
	type body struct {
		ID       int64  `json:"ID"`
		Phone    string `json:"Phone"`
		NewEmail string `json:"NewEmail"`
		Code     string `json:"Code"`
	}
	type resp struct {
		ID      uint
		Content string
	}
	//
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		color.Red("%#v\n", err.Error())
		c.AbortWithStatus(400)
		return
	}
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	color.Red("%d", rby.ID)
	res := db.First(&entiy.User{}, rby.ID).RowsAffected
	if res != 1 {
		c.AbortWithStatus(404)
		return
	}
	//
	rdb := redis2.NewRedisClient()
	if val, err := rdb.Get(context.Background(), rby.Phone+"-"+"ModifyEmail").Result(); errors.Is(err, redis.Nil) {
		c.JSON(http.StatusUnauthorized, resp{
			ID:      0,
			Content: "code expiration",
		})
		return
	} else if val != rby.Code {
		c.JSON(http.StatusUnauthorized, resp{
			ID:      1,
			Content: "code error",
		})
		return
	}
	//
	res = db.Model(&entiy.User{ID: uint(rby.ID)}).Update("email", rby.NewEmail).RowsAffected
	if res == 1 {
		c.Status(http.StatusOK)
		return
	}
	c.JSON(http.StatusInternalServerError, resp{
		ID:      1,
		Content: "update database error",
	})
}
