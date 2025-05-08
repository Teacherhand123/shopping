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
		}
	}
	return r
}
