package handler

import (
	"Game/mysql"
	"Game/mysql/entiy"
	redis2 "Game/redis"
	session "Game/sessions"
	"Game/tools"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
)

// 登录方式:手机号+验证码,手机号+密码,邮箱+密码

// one login request body
type one struct {
	Phone string
	Code  string
}

// two login request body
type two struct {
	Phone    string
	Password string
}

// three login request body
type three struct {
	Email    string
	Password string
}

func Login(c *gin.Context) {
	// get Session store
	store := session.NewSessionStore()
	Session, err := store.Get(c.Request, "Session")
	if err != nil {
		// 400
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// repeat login check
	if !Session.IsNew {
		// 409
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	// retrieve login method
	method := c.Param("method")
	if method != "one" && method != "two" && method != "three" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	switch method {
	case "one":
		{
			var body one
			if err := c.BindJSON(&body); err != nil {
				// 400
				return
			}
			// check if user exists
			db, err := mysql.NewOrmDb()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			res := db.Where("Phone = ?", body.Phone).Find(&entiy.User{})
			if res.RowsAffected != 1 {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			rdb := redis2.NewRedisClient()
			if val, ok := rdb.Get(context.Background(), body.Phone+"-"+"LogIn").Result(); errors.Is(ok, redis.Nil) {
				c.JSON(http.StatusUnauthorized, struct {
					ID   int
					Text string
				}{
					ID:   1,
					Text: "code expiration",
				})
				return
			} else {
				if val != body.Code {
					c.JSON(http.StatusUnauthorized, struct {
						ID   int
						Text string
					}{
						ID:   2,
						Text: "code error",
					})
					return
				}
			}
			// TODO 拿到储存的验证码
		}
	case "two":
		{
			var body two
			check(body, c, Session)
			return
		}
	case "three":
		{
			var body three
			check(body, c, Session)
			return
		}
	}
}
func check(body interface{}, c *gin.Context, s *sessions.Session) bool {
	var bodyTwo two
	var bodyThr three
	sign := true
	// type assertion
	if v, o := body.(two); o {
		bodyTwo = v
	} else {
		v, _ := body.(three)
		bodyThr = v
		sign = false
	}
	// Unmarshal body
	var err error
	if sign {
		err = c.Bind(&bodyTwo)
	} else {
		err = c.Bind(&bodyThr)
	}
	if err != nil {
		return false
	}
	// sql query
	db, err := mysql.NewOrmDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false
	}
	var res *gorm.DB
	var user entiy.User
	if sign {
		res = db.Where(&entiy.User{
			Phone:    bodyTwo.Phone,
			Password: bodyTwo.Password,
		}).First(&user)
	} else {
		res = db.Where(&entiy.User{
			Email:    bodyThr.Email,
			Password: bodyTwo.Password,
		}).First(&user)
	}
	// Account validity check
	if res.RowsAffected == 1 {
		s.Values["ID"] = user.Phone
		s.Save(c.Request, c.Writer)
		tools.Red("%#v%s\n", c.Writer.Header().Get("Set-Cookie"), user.Phone)
		c.AbortWithStatus(http.StatusOK)
		// TODO 查询ip归属地,生成登录历史
		return true
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return false
	}
}