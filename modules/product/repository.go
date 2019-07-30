package product

import (
	"micro-gin/model"
)

// Repository ...
type Repository interface {
	GetByID(id string) (*model.Product, error)
}
