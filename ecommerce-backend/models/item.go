package models

type Item struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name"`
	// Add other fields as needed
}