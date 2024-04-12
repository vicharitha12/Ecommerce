package models

import "time"

type Order struct {
	ID      uint `gorm:"primary_key"`
	CartID  uint `json:"cart_id"`
	// Add additional attributes for the order entity
	CreatedAt time.Time `json:"created_at"`
}