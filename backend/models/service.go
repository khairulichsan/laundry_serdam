// package models

// import "time"

//	type Service struct {
//		ID          uint      `gorm:"primaryKey"`
//		Name        string    `gorm:"size:255;not null"`
//		Price       float64   `gorm:"not null"`
//		Description string    `gorm:"type:text"`
//		CreatedAt   time.Time `gorm:"autoCreateTime"`
//	}
package models

import "time"

type Service struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Price       float64   `gorm:"not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
