package handler

import (
	"Game/RabbitMQ"
	"encoding/json"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"strconv"
)

func SendMessage(c *gin.Context) {
	type body struct {
		UserId         uint
		RoomId         uint
		MessageContent string
	}
	var rby body
	if err := c.BindJSON(&rby); err != nil {
		return
	}
	jsonData, err := json.Marshal(rby)
	//
	rbmqConn, err := RabbitMQ.InitAmpq()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	ch2, err := rbmqConn.Channel()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	err = ch2.Publish("", "RoomId_"+strconv.Itoa(int(rby.RoomId)), false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	})
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	//
	c.Status(200)
}
