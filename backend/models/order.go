package models

type Order struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	UserID uint    `json:"user_id"`
	Status string  `json:"status"`
	Total  float64 `json:"total"`
}
