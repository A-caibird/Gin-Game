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
			// TODO 拿到储存的验证码
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
			Password: bodyThr.Password,
		}).First(&user)
	}
	// Account validity check
	if res.RowsAffected == 1 {
		s.Values["ID"] = user.Phone
		s.Save(c.Request, c.Writer)
		var region, ip string
		if sign {
			ip = bodyTwo.Ip
			region = tools.IpLocationQuery(ip)
		} else {
			ip = bodyThr.Ip
			region = tools.IpLocationQuery(ip)
		}
		//pattern := `region: ([^,]+),`
		//re := regexp.MustCompile(pattern)
		//match := re.FindStringSubmatch(region)
		//if len(match) >= 1 {
		//	region := match[1]
		//	fmt.Println(region)
		//} else {
		//	fmt.Println("No match found")
		//}
		res := db.Create(&entiy.LoginHistory{
			UserId: user.ID,
			Ip:     ip,
			Region: region,
		})
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
		c.JSON(http.StatusUnauthorized, struct {
			ID      int
			Content string
		}{
			ID:      3,
			Content: "user has not exits",
		})
		return false
	}
}
