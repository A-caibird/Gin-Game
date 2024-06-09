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
		return
	}
	if rby.FriendId == rby.UserId {
		c.AbortWithStatus(400)
		return
	}
	//
	db, err := mysql.InitDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	defer db.Close()
	//
	tx, err := db.Begin()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	// exits
	var cnt int
	err = tx.QueryRow("SELECT COUNT(*) from friends WHERE user_id = ? AND  user_friend_id = ?;", rby.UserId, rby.FriendId).Scan(&cnt)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	if cnt > 0 {
		tx.Rollback()
		c.AbortWithStatus(409)
		return
	} else {
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
	}
	c.AbortWithStatus(http.StatusServiceUnavailable)
}
