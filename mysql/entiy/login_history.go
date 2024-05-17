package entiy

import "gorm.io/gorm"

type LoginHistory struct {
	gorm.Model
	UserId uint
	Ip     string
	Region string
}
