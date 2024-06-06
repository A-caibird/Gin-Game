package entiy

import (
	"gorm.io/gorm"
	"time"
)

type LoginHistory struct {
	gorm.Model `json:"-"`
	CreatedAt  time.Time `json:"time"`
	UserId     uint      `json:"-"`
	Ip         string
	Region     string
}
