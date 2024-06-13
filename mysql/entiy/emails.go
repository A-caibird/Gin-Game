package entiy

import "gorm.io/gorm"

type Emails struct {
	gorm.Model `json:"-"`
	SenderId   uint
	ReceiverId uint
	Type       uint
	TypeName   string
	Content    string
}
