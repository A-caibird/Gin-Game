package handler

import (
	redis2 "Game/redis"
	"context"
	"github.com/gin-gonic/gin"
)

func RemoveRoom(c *gin.Context) {
	//
	id := c.Param("id")
	//
	rdb := redis2.NewRedisClient()
	if err := rdb.Del(context.Background(), id).Err(); err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(200)
}
