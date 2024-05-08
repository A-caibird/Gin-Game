//go:build wireinject
// +build wireinject

package RabbitMQ

import (
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitAmpq() (conn *amqp.Connection, err error) {
	wire.Build(NewAmqp)
	return nil, nil
}
