// package models

//	type User struct {
//		ID       uint   `json:"id" gorm:"primaryKey"`
//		Name     string `json:"name"`
//		Email    string `json:"email" gorm:"unique"`
//		Password string `json:"password"`
//	}
package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	Phone     string    `gorm:"size:15"`
	Role      string    `gorm:"size:10;not null"` // enum: admin, karyawan, pelanggan
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
