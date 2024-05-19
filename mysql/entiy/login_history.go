package entiy

import "gorm.io/gorm"

type LoginHistory struct {
	gorm.Model `json:"-"`
	UserId     uint `json:"-"`
	Ip         string
	Region     string
}
