package types

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID uint `json:"id"` 
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age uint `json:"age"`
	Salary float32 `json:"salary"`
	Position string `json:"position"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ClockKey string `json:"clockKey"`
	Timelog []Timelog `json:"timelog" gorm:"foreignKey:EmployeeID;references:ID"`
}