package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FriendNotifyEmail 生成添加好友的邮件
func FriendNotifyEmail(c *gin.Context) {
	type body struct {
		UserId   uint
		FriendId uint
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		color.Red("%s", err.Error())
		return
	}
	//
	if rby.FriendId == rby.UserId {
		c.AbortWithStatus(400)
		return
	}
	//
	var sender_name string
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
	}
	var user entiy.User
	db.Where(entiy.User{
		ID: rby.UserId,
	}).First(&user)
	sender_name = user.Name
	content := "用户" + sender_name + "请求添加您为好友!"
	//
	var res = db.Create(&entiy.Emails{
		SenderId:   rby.UserId,
		ReceiverId: rby.FriendId,
		Type:       1,
		TypeName:   "加好友",
		Content:    content,
	}).RowsAffected
	if res == 1 {
		c.AbortWithStatus(200)
		return
	}
	c.Status(500)
}

func GetEmail(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	color.Red("%d", idInt)
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	var email entiy.Emails
	db.Where(entiy.Emails{
		ReceiverId: uint(idInt),
	}).First(&email)
	//
	c.JSON(200, email)
}
