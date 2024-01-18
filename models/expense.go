package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Caterory    string  `json:"category"`
	Description string  `json:"description"`
}
