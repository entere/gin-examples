package delivery

import (
	"micro-gin/modules/coupon/repository"
	"micro-gin/modules/coupon/service"

	"micro-gin/handler"

	"micro-gin/pkg/errno"

	"github.com/gin-gonic/gin"
)

// CouponDelivery ...
var CouponDelivery = newCouponDelivery()

type couponDelivery struct {
}

func newCouponDelivery() *couponDelivery {
	return &couponDelivery{}
}

func (cd *couponDelivery) Show(c *gin.Context) {
	couponID := c.Param("id")
	couponRepository := repository.NewCouponRepository()
	couponService := service.NewCouponService(couponRepository)
	coupon, err := couponService.GetByID(couponID)
	// userService := services.UserService{ID: userID}
	// coupon, err := couponService.GetByID(couponID)
	if err != nil {

		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return

	}

	handler.SendResponse(c, nil, coupon)

}
