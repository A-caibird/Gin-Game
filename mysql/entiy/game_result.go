package entiy

import (
	"gorm.io/gorm"
	"time"
)

type GameResult struct {
	gorm.Model    `json:"-"`
	CreatedAt     time.Time `json:"CreatedAt"`
	UserId        int
	Type          string
	Mode          string
	Count         int
	PointsChanges int
	BeansChanges  int
	Result        bool
}
