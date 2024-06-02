package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func SearchUser(c *gin.Context) {
	//
	method := c.Param("method")
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	if method == "name" {
		name, ok := c.GetQuery("name")
		if !ok {
			c.AbortWithStatus(400)
			color.Red("%s", name)
			return
		}
		//
		var users []entiy.User
		db.Model(entiy.User{}).Where("name LIKE ?", "%"+name+"%").Find(&users)
		//
		var res []struct {
			ID     uint
			Name   string
			Online bool
		}
		for _, v := range users {
			res = append(res, struct {
				ID     uint
				Name   string
				Online bool
			}{
				ID:     v.ID,
				Name:   v.Name,
				Online: v.Online,
			})
		}
		c.JSON(200, res)
		return
	} else {
		phone, ok := c.GetQuery("phone")
		if !ok {
			c.AbortWithStatus(400)
			color.Red("%s", phone)
			return
		}
		//
		var users []entiy.User
		db.Model(entiy.User{}).Where("phone LIKE ?", "%"+phone+"%").Find(&users)
		//
		var res []struct {
			ID     uint
			Name   string
			Online bool
		}
		for _, v := range users {
			res = append(res, struct {
				ID     uint
				Name   string
				Online bool
			}{
				ID:     v.ID,
				Name:   v.Name,
				Online: v.Online,
			})
		}
		c.JSON(200, res)
		return
	}
}
