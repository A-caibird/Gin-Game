package entiy

import (
	"gorm.io/gorm"
	"time"
)

type Emails struct {
	gorm.Model `json:"-"`
	CreatedAt  time.Time `json:"Time"`
	SenderId   uint
	ReceiverId uint
	Type       uint
	TypeName   string
	Content    string
}
