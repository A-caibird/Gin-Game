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
		res := db.Where("name=?", "root").First(&user)
		fmt.Print(res.RowsAffected)
	}
	//var user entiy.User
	res := db.Where("name=?", "root1").First(&entiy.User{})
	fmt.Print(res.RowsAffected)
}
