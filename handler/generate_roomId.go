package handler

import (
	redis2 "Game/redis"
	"context"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"math/big"
)

var cnt *big.Int

func init() {
	cnt = new(big.Int)
	cnt.SetInt64(1)
}

func RoomId(c *gin.Context) {
	//
	id := c.Param("id")
	//
	b := big.NewInt(1)
	b.SetString("1", 10)
	cnt = b.Add(cnt, b)
	c.JSON(200, cnt)
	//
	rdb := redis2.NewRedisClient()
	if err := rdb.RPush(context.Background(), cnt.String(), id).Err(); err != nil {
		color.Red("%s\n", err.Error())
	}
}
