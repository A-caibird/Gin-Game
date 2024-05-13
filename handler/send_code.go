package handler

import (
	redis2 "Game/redis"
	"Game/tools"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"
)

// TODO 发送验证码

func SendCode(c *gin.Context) {
	//  usage
	//  SignUP, LogIn
	usage := c.Param("usage")
	//
	type body struct {
		Phone string
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// generate rand code
	var max uint64 = 999999
	var min uint64 = 100000
	r := rand.New(rand.NewPCG(1, 2))
	code := strconv.FormatUint(r.Uint64N(max-min+1)+min, 10)
	// send
	res, err := tools.SendSMS(rby.Phone, code)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if *res.StatusCode != 200 {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// save to redis
	rdb := redis2.NewRedisClient()
	if _, err := rdb.Set(context.Background(), rby.Phone+"-"+usage, code, 5*time.Minute).Result(); errors.Is(err, redis.Nil) {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.AbortWithStatus(http.StatusOK)
}
