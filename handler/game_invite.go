package handler

import (
	"Game/RabbitMQ"
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
	//
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
	// Notify friends if I accept game invitations
	err = ch2.Publish("", "user_"+strconv.Itoa(rby.FriendId)+"_invite_result", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsondata,
	})
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Status(200)
}
