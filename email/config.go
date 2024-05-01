package email

import (
	. "Game/tools"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func NewDialer() *gomail.Dialer {
	d := gomail.NewDialer(Conf.Email.Host, Conf.Email.Port, Conf.Email.User, Conf.Email.AuthPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d
}
