package handler

import (
	"github.com/gin-gonic/gin"
	"math/big"
)

var cnt *big.Int

func init() {
	cnt = new(big.Int)
	cnt.SetInt64(1)
}

func RoomId(c *gin.Context) {
	b := big.NewInt(1)
	b.SetString("1", 10)
	cnt = b.Add(cnt, b)
	c.JSON(200, cnt)
}
