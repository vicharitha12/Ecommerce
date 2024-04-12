package models

type CartItem struct {
	ID     uint `gorm:"primary_key"`
	CartID uint `json:"cart_id"`
	ItemID uint `json:"item_id"`
	// Add additional attributes for the cart item entity
}