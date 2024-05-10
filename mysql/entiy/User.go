package entiy

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Phone    string
	Email    string
}
