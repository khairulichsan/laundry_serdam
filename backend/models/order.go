package models

import "time"

type Order struct {
	ID           uint    `gorm:"primaryKey"`
	UserID       uint    `gorm:"not null"`
	User         User    `gorm:"foreignKey:UserID"`
	Status       string  `gorm:"size:20;not null"` // enum: baru, proses, selesai
	TotalPrice   float64 `gorm:"not null"`
	PickupDate   time.Time
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	ServiceID uint    `gorm:"not null"`
	Service   Service `gorm:"foreignKey:ServiceID"`
	Quantity  uint    `gorm:"not null"`
	Subtotal  float64 `gorm:"not null"`
}
