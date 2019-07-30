package model

import (
	// _ mysql
	_ "github.com/go-sql-driver/mysql"
)

// Coupon ...
type Coupon struct {
	BaseModel
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
