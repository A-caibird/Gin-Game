package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	Name     string
	Password string
	ID       uint
	Email    string
	Active   int
}

func TestOrm(t *testing.T) {
	dsn := "root:775028@tcp(127.0.0.1:3306)/WebMusic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	var user User
	res := db.Where("name=?", "root").First(&user).RowsAffected
	println(res, user.Password)
}
