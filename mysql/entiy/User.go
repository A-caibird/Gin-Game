package entiy

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint `gorm:"primarykey"`
	Name       string
	Password   string `json:"-"`
	Phone      string
	Email      string
}
