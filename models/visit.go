package models

import (
	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model
	ID         uint32   `json:"id"`
	Location   Location `json:"location"`
	LocationID int
	User       User `json:"user"`
	UserID     int
	VisitedAt  uint64 `json:"visited_at"`
	Mark       uint8  `json:"mark"`
}
