package mysql

import (
	"Game/mysql/entiy"
	"fmt"
	"os"
	"testing"
)

func TestOrm(t *testing.T) {
	db, err := NewOrmDb()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}
	if !db.Migrator().HasTable(&entiy.User{}) {
		db.Migrator().CreateTable(&entiy.User{})
		db.Create(&entiy.User{
			Name:     "root",
			Password: "root",
			Phone:    "1",
			Email:    "fdasf",
		})
		var user entiy.User
		db.Where("name=?", "root").First(&user)
		fmt.Print(user.Name)
	}
}
