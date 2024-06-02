package entiy

import "gorm.io/gorm"

type UserBackpack struct {
	gorm.Model  `json:"-"`
	Diamond     int
	Beans       int
	CardCounter int
	UserId      int
}
