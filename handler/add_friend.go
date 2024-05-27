package handler

import (
	"Game/mysql"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddFriend(c *gin.Context) {
	type body struct {
		UserId   uint
		FriendId uint
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		color.Red("%s", err.Error())
	}
	//
	db, err := mysql.InitDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	defer db.Close()
	//
	smtp, err := db.Prepare("insert into friends (created_at, updated_at, user_id, user_friend_id) values (?,?,?,?)")
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	res, err := smtp.Exec(time.Now(), time.Now(), rby.UserId, rby.FriendId)
	if cnt, _ := res.RowsAffected(); cnt == 1 {
		c.AbortWithStatus(200)
		return
	} else {
		color.Red("%#v", cnt)
	}
	c.AbortWithStatus(http.StatusServiceUnavailable)
}
