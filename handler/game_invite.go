package handler

import (
	"Game/RabbitMQ"
	"Game/mysql"
	"Game/mysql/entiy"
	redis2 "Game/redis"
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"strconv"
)

// Invite friends to a game
//
// create RabbitMQ queue name: user_FriendId_ invite
func Invite(c *gin.Context) {
	type body struct {
		UserId   int
		FriendId int
		RoomId   string
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		return
	}
	color.Red("%#v\n", rby)
	jsonData, err := json.Marshal(rby)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	//
	rbmqConn, err := RabbitMQ.InitAmpq()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	ch2, err := rbmqConn.Channel()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	queue, err := ch2.QueueDeclare("user_"+strconv.Itoa(rby.FriendId)+"_invite", true, false, false, false, nil)
	if err != nil {
		c.AbortWithStatus(500)
		color.Red("%s", err.Error())
		return
	}
	color.Red("%s", queue.Name)
	err = ch2.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	})
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(200)
}

// HandleInvite Notify friends if I accept game invitations
//
// RabbitMQ queue name: user_FriendId_ invite_handle
func HandleInvite(c *gin.Context) {
	type body struct {
		UserId   int // I
		FriendId int // Friends who invited you to play the game
		RoomId   string
		Result   bool
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		return
	}
	jsondata, err := json.Marshal(rby)
	if rby.Result {
		db, err := mysql.NewOrmDb()
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		// save all user in room
		rdb := redis2.NewRedisClient()
		if _, err := rdb.RPush(context.Background(), rby.RoomId, rby.UserId).Result(); err != nil {
			c.AbortWithStatus(500)
			return
		}
		// retrieve all user in room
		var userListInRoom []struct {
			UserId   int    `json:"userId"`
			UserName string `json:"userName"`
		}
		elements, err := rdb.LRange(context.Background(), rby.RoomId, 0, -1).Result()
		if err != nil {
			color.Red("3333 %s", err.Error())
		}
		color.Red("1111 %#v", elements)
		for _, val := range elements {
			e, _ := strconv.Atoi(val)
			var user entiy.User
			db.Where(entiy.User{ID: uint(e)}).First(&user)
			userListInRoom = append(userListInRoom, struct {
				UserId   int    `json:"userId"`
				UserName string `json:"userName"`
			}{
				UserId:   e,
				UserName: user.Name,
			})
		}
		//
		c.JSON(200, userListInRoom)
	}
	// notify room owner if I accept game invitations
	conn, err := RabbitMQ.NewAmqp()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	ch2, err := conn.Channel()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	err = ch2.Publish("", "user_"+strconv.Itoa(rby.FriendId)+"_invite_result", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsondata,
	})
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.AbortWithStatus(200)
}
