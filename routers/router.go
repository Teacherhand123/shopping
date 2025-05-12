package routers

import (
	"net/http"
	api "shopping/api/v1"
	"shopping/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"res": "pong",
			})
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 轮播图
		v1.POST("carousels", api.ListCarousel)

		// 商品操作
		v1.GET("products", api.ListProduct)
		v1.GET("products/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)

		authed := v1.Group("/") // 需要登陆保护的
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.PUT("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			// 显示金额
			authed.POST("money", api.ShowMoney)

			// 商品操作
			authed.POST("product", api.CreateProduct)
			authed.POST("products", api.SearchProduct)

			// 收藏夹操作
			authed.GET("favorites", api.ShowFavorite)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)

			// 地址操作
			authed.GET("address", api.GetAddresses)
			authed.POST("address", api.CreateAddress)
			authed.GET("address/:id", api.ShowAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.DELETE("address/:id", api.DeleteAddress)

			// 购物车操作
			authed.POST("cart", api.CreateCart)
			authed.GET("cart", api.GetCarts)
			authed.PUT("cart/:id", api.UpdateCart)
			authed.DELETE("cart/:id", api.DeleteCart)

			// 订单操作
			authed.POST("order", api.CreateOrder)
			authed.GET("order", api.GetOrders)
			authed.PUT("order/:id", api.UpdateOrder)
			authed.DELETE("order/:id", api.DeleteOrder)
		}
	}
	return r
}
