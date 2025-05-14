package types

import (
	"time"

	"gorm.io/gorm"
)

type Timelog struct {
	gorm.Model
	ID uint `json:"id"`
	EmployeeID uint `json:"employeeId"`
	IsClockIn bool `json:"isClockIn"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}