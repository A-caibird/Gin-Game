package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Pay CNY payment interface

func BuyDiamond() func(c *gin.Context) {
	return func(c *gin.Context) {
		type body struct {
			UserId int
			Amount int
		}
		var rby body
		if err := c.ShouldBindJSON(&rby); err != nil {
			return
		}
		//
		db, err := mysql.NewOrmDb()
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		//
		res := db.Model(entiy.UserBackpack{}).Where("user_id = ?", rby.UserId).Update("diamond", rby.Amount)
		if res.RowsAffected == 1 {
			c.Status(200)
			return
		}
		c.AbortWithStatus(http.StatusPaymentRequired)
	}
}

func BuyGameProp(c *gin.Context) {
	type body struct {
		UserId int
		Count  int
		Amount int // Unit price: 15 diamonds
	}
	var rby body
	if err := c.ShouldBindJSON(&rby); err != nil {
		return
	}
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	var userBackPack entiy.UserBackpack
	res := db.Model(&entiy.UserBackpack{}).Where("user_id = ?", rby.UserId).First(&userBackPack)
	if res.RowsAffected != 1 {
		c.AbortWithStatus(400)
		return
	}
	if userBackPack.Diamond < rby.Amount {
		c.AbortWithStatus(402)
	} else {
		db.Model(&entiy.UserBackpack{}).Where("user_id = ? ", rby.UserId).Update("diamond", userBackPack.Diamond-rby.Amount)
		db.Model(&entiy.UserBackpack{}).Where("user_id = ? ", rby.UserId).Update("card_counter", userBackPack.CardCounter+rby.Count)
		c.AbortWithStatus(200)
	}
}

func BuyBeans(c *gin.Context) {
	type body struct {
		UserId int
		Count  int
		Amount int
	}
	var rby body
	if err := c.ShouldBindJSON(&rby); err != nil {
		return
	}
	//
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	var userBackPack entiy.UserBackpack
	res := db.Model(&entiy.UserBackpack{}).Where("user_id = ?", rby.UserId).First(&userBackPack)
	if res.RowsAffected != 1 {
		c.AbortWithStatus(400)
		return
	}
	if userBackPack.Diamond < rby.Amount {
		c.AbortWithStatus(402)
	} else {
		db.Model(&entiy.UserBackpack{}).Where("user_id = ? ", rby.UserId).Update("diamond", userBackPack.Diamond-rby.Amount)
		db.Model(&entiy.UserBackpack{}).Where("user_id = ? ", rby.UserId).Update("beans", userBackPack.Beans+rby.Count)
		c.AbortWithStatus(200)
	}
}
