package delivery

import (
	"micro-gin/handler"
	"micro-gin/modules/product/repository"
	"micro-gin/modules/product/service"

	"micro-gin/pkg/errno"

	"github.com/gin-gonic/gin"
)

// ProductDelivery ...
var ProductDelivery = newProductDelivery()

type productDelivery struct {
}

func newProductDelivery() *productDelivery {
	return &productDelivery{}
}

func (pd *productDelivery) Show(c *gin.Context) {
	productID := c.Param("id")
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	product, err := productService.GetByID(productID)
	// userService := services.UserService{ID: userID}
	// product, err := productService.GetByID(productID)
	if err != nil {

		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return

	}

	handler.SendResponse(c, nil, product)

}
