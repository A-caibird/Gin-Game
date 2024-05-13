package handler

import "C"
import (
	"Game/mysql"
	"Game/mysql/entiy"
	redis2 "Game/redis"
	session "Game/sessions"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"sync"
)

type body struct {
	Phone    string
	Name     string
	Password string
	Email    string
	Code     string
}

func SignUp(c *gin.Context) {
	Store := session.NewSessionStore()
	Session, err := Store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var rby body
	if ok := c.BindJSON(&rby); ok != nil {
	}
	// retrieve db sql
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// check if user exists
	res := db.Where("Phone = ?", rby.Phone).First(&entiy.User{})
	if res.RowsAffected != 0 {
		c.String(http.StatusConflict, "user has exited")
		return
	}
	// verify code match check
	rdb := redis2.NewRedisClient()
	if val, err := rdb.Get(context.Background(), rby.Phone+"-"+"SignUp").Result(); errors.Is(err, redis.Nil) {
		c.JSON(http.StatusUnauthorized, struct {
			ErrorID int
			Info    string
		}{
			ErrorID: 1,
			Info:    "code has expiration",
		})
		return
	} else {
		if val != rby.Code {
			c.JSON(http.StatusUnauthorized, struct {
				ErrorID int
				Info    string
			}{
				ErrorID: 1,
				Info:    "code error",
			})
			return
		}
	}
	// write user info to database
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		res = db.Create(&entiy.User{
			Phone:    rby.Phone,
			Email:    rby.Email,
			Password: rby.Password,
			Name:     rby.Name,
		})
		if res.RowsAffected != 1 {
			c.String(http.StatusInternalServerError, "when writing user information to database throw a error")
		} else {
			c.String(http.StatusOK, "sign up successfully")
		}
		wg.Done()
	}(&wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		Session.Values["user_phone"] = rby.Phone
		Session.Save(c.Request, c.Writer)
		wg.Done()
	}(&wg)
	wg.Wait()
}
