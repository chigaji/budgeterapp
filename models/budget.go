package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}
