package handler

import (
	"Game/mysql"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func GetFriendInfo(c *gin.Context) {
	// retrieve user id
	id := c.Param("id")
	// db connection
	db, err := mysql.InitDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	defer db.Close()
	// prepare query
	smtp, err := db.Prepare("SELECT B.id,B.name, B.online FROM  friends  A JOIN users B ON A.user_id = B.id WHERE A.user_id = ?;")
	if err != nil {
		color.Red("%s", err.Error())
		c.AbortWithStatus(500)
		return
	}
	//
	row, err := smtp.Query(id)
	defer row.Close()
	//
	var res []struct {
		FriendId uint
		Name     string
		Online   bool
	}
	for row.Next() {
		var item struct {
			FriendId uint
			Name     string
			Online   bool
		}
		err := row.Scan(&item.FriendId, &item.Name, &item.Online)
		if err != nil {
			color.Red("%s", err.Error())
			break
		}
		//color.Blue("%#v", item)
		res = append(res, item)
	}
	c.JSON(200, res)
}
