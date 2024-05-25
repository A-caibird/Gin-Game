package handler

import (
	"Game/email"
	redis2 "Game/redis"
	"Game/tools"
	"context"
	"errors"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"math/rand/v2"
	"strconv"
	"time"
)

func GetEmailCode(c *gin.Context) {
	//
	usage := c.Param("usage")
	//
	var body struct {
		Email string
		Phone string
	}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}
	r := rand.New(rand.NewPCG(1, 1))
	code := strconv.FormatUint(r.Uint64N(900000)+100000, 10)
	// send code
	d := email.InitMailDialer()
	m := gomail.NewMessage()
	m.SetHeader("From", tools.Conf.Email.User)
	m.SetHeader("To", body.Email)
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Verify Code")
	m.SetBody("text/html", "<b>Code</b>:  <i>"+code+"</i>,  5分钟过后过期!")
	if err := d.DialAndSend(m); err != nil {
		color.Red("fasdfasdfs")
		c.AbortWithStatus(500)
	}
	//
	rdb := redis2.NewRedisClient()
	if _, err := rdb.Set(context.Background(), body.Phone+"-"+usage, code, 5*time.Minute).Result(); errors.Is(err, redis.Nil) {
		c.AbortWithStatus(500)
		return
	}
}
