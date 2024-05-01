//go:build wireinject
// +build wireinject

package email

import (
	"github.com/google/wire"
	"gopkg.in/gomail.v2"
)

func InitMailDialer() *gomail.Dialer {
	wire.Build(NewDialer)
	return nil
}
