package routes

import (
	api "E-commerce/api/v1"
	"E-commerce/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//用户操作
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		//轮播图
		v1.GET("/carousels", api.ListCarousel)

		//商品操作
		v1.GET("/products", api.ListProduct)
		v1.GET("/products/:id", api.ShowProduct)
		v1.GET("/imgs/:id", api.ListProductImg)
		v1.GET("/categories", api.ListCategory)

		authed := v1.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user/update", api.UserUpdate)
			authed.POST("user/avatar", api.UpLoadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			//显示金额
			authed.POST("money", api.ShowMoney)

			//商品操作
			authed.POST("product", api.CreateProduct)
			authed.POST("searchproduct", api.SearchProduct)

			//收藏夹操作
			authed.POST("favorites", api.CreateFavorite)
			authed.GET("favorites", api.ListFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)

			//地址操作
			authed.POST("address", api.CreateAddress)
			authed.GET("address/:id", api.GetAddress)
			authed.PUT("change-address/:id", api.UpdateAddress)
			authed.DELETE("delete-address/:id", api.DeleteAddress)
			authed.GET("addresses", api.ListAddresses)

			//购物车操作
			authed.POST("cart", api.CreateCart)
			authed.GET("cart/:id", api.GetCart)
			authed.PUT("change-cart/:id", api.UpdateCart)
			authed.DELETE("delete-cart/:id", api.DeleteCart)

			//订单操作
			authed.POST("order", api.CreateOrder)
			authed.GET("order/:id", api.GetOrder)
			authed.GET("show-order/:id", api.ShowOrder)
			authed.DELETE("delete-order/:id", api.DeleteOrder)
			authed.POST("pay-order", api.PayOrder)
		}
	}
	return r
}
