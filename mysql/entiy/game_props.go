package entiy

import "gorm.io/gorm"

type GameProps struct {
	gorm.Model `json:"-"`
	Name       string
	Price      int
}
