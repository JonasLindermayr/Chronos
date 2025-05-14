package types

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email string `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}