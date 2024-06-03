package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func GetDiamondPrice(c *gin.Context) {
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	var diamondPrice []entiy.DiamondPrice
	db.Model(entiy.DiamondPrice{}).Find(&diamondPrice)
	c.JSON(200, diamondPrice)
}

func GetBeansPrice(c *gin.Context) {
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	var gameBeans []entiy.GameBeans
	db.Model(entiy.GameBeans{}).Find(&gameBeans)
	c.JSON(200, gameBeans)
}

func GetGameProps(c *gin.Context) {
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	var gameProps []entiy.GameProps
	db.Model(entiy.GameProps{}).Find(&gameProps)
	c.JSON(200, gameProps)
}

func GetUserBackpack(c *gin.Context) {
	//
	id := c.Param("id")
	//
	db, err := mysql.InitDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	defer db.Close()
	//
	smtp, err := db.Prepare("SELECT * from GinGame.user_backpacks Where user_id = ?;")
	if err != nil {
		c.AbortWithStatus(500)
		color.Red("%s", err.Error())

		return
	}
	//
	rows, err := smtp.Query(id)
	defer rows.Close()
	var res []entiy.UserBackpack
	for rows.Next() {
		var item entiy.UserBackpack
		err := rows.Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt, &item.Diamond, &item.Beans, &item.CardCounter, &item.UserId)
		if err != nil {
			color.Red("%#v", err.Error())
			break
		}
		res = append(res, item)
	}
	//
	c.JSON(200, res)
}
