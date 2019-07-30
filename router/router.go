package router

import (
	coupon_delivery "micro-gin/modules/coupon/delivery"
	product_delivery "micro-gin/modules/product/delivery"
	"micro-gin/modules/sd"
	user_delivery "micro-gin/modules/user/delivery"
	wechat_delivery "micro-gin/modules/wechat/delivery"

	"net/http"

	"micro-gin/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/v1/user")

	u.Use(middleware.JWTAuthMiddleware())
	{
		// u.GET("/", user.Index)
		u.GET("/:id", user_delivery.UserDelivery.Show)
	}

	// 注册
	signup := g.Group("/v1/signup")
	{

		signup.POST("", user_delivery.UserDelivery.Signup)

	}

	// 登录
	signin := g.Group("/v1/signin")
	{
		signin.POST("", user_delivery.UserDelivery.Signin)
	}

	// 产品
	ps := g.Group("/v1/product")
	{
		ps.GET("/:id", product_delivery.ProductDelivery.Show)
	}

	// 产品
	coupon := g.Group("/v1/coupon")
	{
		coupon.GET("/:id", coupon_delivery.CouponDelivery.Show)
	}

	// 产品
	wechat := g.Group("/v1/wechat")
	{
		wechat.GET("/oauth", wechat_delivery.WechatDelivery.OauthHandler)
		wechat.GET("/oauth/redirect", wechat_delivery.WechatDelivery.OauthRedirectHandler)
		wechat.Any("/server", wechat_delivery.WechatDelivery.ServerHandler)
		wechat.GET("/template", wechat_delivery.WechatDelivery.TemplateHandler)
		wechat.GET("/qrcode/tmp", wechat_delivery.WechatDelivery.QRCodeTmpHandler)     // 临时二维码
		wechat.GET("/qrcode/limit", wechat_delivery.WechatDelivery.QRCodeLimitHandler) // 永久二维码
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
