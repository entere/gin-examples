package model

import (
	// _ mysql
	_ "github.com/go-sql-driver/mysql"
)

// Product ...
type Product struct {
	BaseModel
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
