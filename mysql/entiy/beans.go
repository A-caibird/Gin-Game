package entiy

import "gorm.io/gorm"

type GameBeans struct {
	gorm.Model `json:"-"`
	Name       string
	Price      int
}
