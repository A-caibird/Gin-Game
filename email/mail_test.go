package email

import (
	"Game/tools"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestInitMailDialer(t *testing.T) {
	d := InitMailDialer()
	m := gomail.NewMessage()
	m.SetHeader("From", tools.Conf.Email.User)
	m.SetHeader("To", "newcoder@icloud.conm", "lian04201@outlook.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
