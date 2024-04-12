package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}