package handler

import (
	"Game/RabbitMQ"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"strconv"
)

func Invite(c *gin.Context) {
	type body struct {
		UserId   int
		FriendId int
		RoomId   int
	}
	var rby body
	if err := c.ShouldBindJSON(&rby); err != nil {
		return
	}
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
