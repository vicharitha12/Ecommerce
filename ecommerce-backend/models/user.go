package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CartID    uint      `json:"cart_id"`
	CreatedAt time.Time `json:"created_at"`
}