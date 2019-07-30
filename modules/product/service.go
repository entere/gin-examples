package product

import "micro-gin/model"

// Service ...
type Service interface {
	GetByID(id string) (*model.Product, error)
}
