package handler

import (
	"Game/RabbitMQ"
	"Game/mysql"
	"Game/mysql/entiy"
	redis2 "Game/redis"
	session "Game/sessions"
	"Game/tools"
	"context"
	"encoding/json"
	"errors"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// 登录方式:手机号+验证码,手机号+密码,邮箱+密码

// one login request body
type one struct {
	Phone string
	Code  string
	Ip    string
}

// two login request body
type two struct {
	Phone    string
	Password string
	Ip       string
}

// three login request body
type three struct {
	Email    string
	Password string
	Ip       string
}

func Login(c *gin.Context) {
	// get Session store
	store := session.NewSessionStore()
	Session, err := store.Get(c.Request, "session")
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
			//
			var user entiy.User
			res := db.Where("Phone = ?", body.Phone).Find(&user)
			if res.RowsAffected != 1 {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			// check verify code
			rdb := redis2.NewRedisClient()
			if val, ok := rdb.Get(context.Background(), body.Phone+"-"+"LogIn").Result(); errors.Is(ok, redis.Nil) {
				c.JSON(http.StatusUnauthorized, struct {
					ID      int
					Content string
				}{
					ID:      1,
					Content: "code expiration",
				})
				return
			} else {
				if val != body.Code {
					c.JSON(http.StatusUnauthorized, struct {
						ID      int
						Content string
					}{
						ID:      2,
						Content: "code error",
					})
					return
				}
			}
			// login successfully
			c.AbortWithStatus(200)
			// notify friend I are online now
			NotifyFriend(body.Phone)
			// login history
			GenerateLoginHistory(body.Ip, user, db)
			//
			Session.Values["ID"] = user.ID
			Session.Save(c.Request, c.Writer)
			return
		}
	case "two":
		{
			var body two
			check(body, c, Session)
			return
		}
	default:
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
		v, o := body.(three)
		if o {
			bodyThr = v
			sign = false
		}
	}
	// Unmarshal body
	var err error
	if sign {
		err = c.BindJSON(&bodyTwo)
	} else {
		err = c.BindJSON(&bodyThr)
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
	//
	var res *gorm.DB
	var user entiy.User
	if sign {
		// check if user  exits
		res = db.Where(&entiy.User{
			Phone: bodyTwo.Phone,
		}).First(&user)
		if res.RowsAffected != 1 {
			c.AbortWithStatus(404)
			return false
		}
		// check password
		res = db.Where(&entiy.User{
			Phone:    bodyTwo.Phone,
			Password: bodyTwo.Password,
		}).First(&user)
		if res.RowsAffected != 1 {
			c.AbortWithStatus(401)
			return false
		}
	} else {
		// check if user exits
		res = db.Where(&entiy.User{
			Email: bodyThr.Email,
		}).First(&user)
		if res.RowsAffected != 1 {
			c.AbortWithStatus(404)
			return false
		}
		// check password
		res = db.Where(&entiy.User{
			Email:    bodyThr.Email,
			Password: bodyThr.Password,
		}).First(&user)
		if res.RowsAffected != 1 {
			c.AbortWithStatus(401)
			return false
		}
	}
	// Account validity check
	if res.RowsAffected == 1 {
		// write session data to cookie
		s.Values["ID"] = user.ID
		s.Save(c.Request, c.Writer)
		//
		var ip string
		if sign {
			ip = bodyTwo.Ip
		} else {
			ip = bodyThr.Ip
		}
		//Notify friends that I'm online
		ok, err := NotifyFriend(bodyTwo.Phone)
		color.Red("%#v %#v", ok, err)
		//Generate login history
		res := GenerateLoginHistory(ip, user, db)
		if res.RowsAffected == 1 {
			c.JSON(http.StatusOK, struct {
				ID      uint8
				Content entiy.User
			}{
				ID:      1,
				Content: user,
			})
			return true
		}
		c.JSON(http.StatusOK, struct {
			ID      uint8
			Content string
		}{
			ID:      1,
			Content: "save login history error",
		})
		return false
	} else {
		c.Status(401)
		return false
	}
}

func NotifyFriend(phone string) (bool, error) {
	db, err := mysql.Newdb()
	if err != nil {
		return false, err
	}
	defer db.Close()
	// query user's  the id of friend
	smtp, err := db.Prepare("SELECT id FROM GinGame.users WHERE id IN ( SELECT user_friend_id  FROM GinGame.friends  WHERE user_id = (select id from GinGame.users where phone = ?));")
	if err != nil {
		return false, err
	}
	row, err := smtp.Query(phone)
	defer row.Close()
	//
	type friendInfo struct {
		Id int
	}
	var friendList []friendInfo
	var item friendInfo
	for row.Next() {
		err := row.Scan(&item.Id)
		if err != nil {
			color.Red("%s", err.Error())
			break
		}
		friendList = append(friendList, item)
	}
	// rabbitMq
	conn, err := RabbitMQ.InitAmpq()
	if err != nil {
		return false, err
	}
	defer conn.Close()
	//
	ch2, err := conn.Channel()
	defer ch2.Close()
	// query userId
	var user struct {
		Id   int
		Name string
	}
	row, err = db.Query("SELECT id,name FROM GinGame.users WHERE phone=?;", phone)
	if err != nil {
		return false, err
	}
	for row.Next() {
		err := row.Scan(&user.Id, &user.Name)
		if err != nil {
			break
		}
	}
	jsonData, err := json.Marshal(user)
	//Production Message
	for _, v := range friendList {
		err = ch2.Publish("", "user_"+strconv.Itoa(v.Id)+"_Friend_Login_Notify", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	}
	return true, nil
}

func GenerateLoginHistory(ip string, user entiy.User, db *gorm.DB) *gorm.DB {
	return db.Create(&entiy.LoginHistory{
		UserId: user.ID,
		Ip:     ip,
		Region: tools.IpLocationQuery(ip),
	})
}
