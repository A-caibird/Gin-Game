package RabbitMQ

import (
	"Game/tools"
	amqp "github.com/rabbitmq/amqp091-go"
)

var ConnetString = "amqp://" + tools.Conf.RabbitMQ.User + ":" + tools.Conf.RabbitMQ.Password +
	"@" + tools.Conf.RabbitMQ.Host + ":" + tools.Conf.RabbitMQ.Port + "/"

func NewAmqp() (conn *amqp.Connection, err error) {
	conn, err = amqp.Dial(ConnetString)
	return
}
