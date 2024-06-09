package handler

import (
	"Game/mysql"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func SearchUser(c *gin.Context) {
	//
	patter := c.Query("pattern")
	//
	db, err := mysql.Newdb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	smtp, err := db.Prepare("SELECT DISTINCT u.id,u.name,u.online\n FROM GinGame.users u\n WHERE u.phone LIKE ? \n   OR u.name LIKE ? \n   AND u.deleted_at IS NULL\n GROUP BY u.id;")
	if err != nil {
		color.Red("%s", err.Error())
		c.AbortWithStatus(500)
		return
	}
	//
	rows, err := smtp.Query("%"+patter+"%", "%"+patter+"%")
	defer rows.Close()
	//
	var res []struct {
		Id     int
		Name   string
		Online bool
	}
	for rows.Next() {
		var item struct {
			Id     int
			Name   string
			Online bool
		}
		err := rows.Scan(&item.Id, &item.Name, &item.Online)
		if err != nil {
			color.Red("%s", err.Error())
			break
		}
		res = append(res, item)
	}
	c.JSON(200, res)
	return
}
