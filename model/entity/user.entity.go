package entity

import (
	"time"

	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model // ID, CreatedAt, UpdatedAt, tidak perlu dimasukin lagi
// 	Name string
// }

type User struct {
	ID        	uint `json:"id" gorm:"primaryKey"`
	Name 		string `json:"name" validate:"required,min=2"`
	Email 		string `json:"email" validate:"required,email"`
	Password 	string `json:"password" validate:"required"`
	Address 	string `json:"address"`
	Phone 		string `json:"phone"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `json:"-" gorm:"index, column: deleted_at"`
}