package entiy

import "gorm.io/gorm"

type DiamondPrice struct {
	gorm.Model `json:"-"`
	Name       string
	Price      int
}
