package v1

import (
	"net/http"
	"shopping/pkg/utils"
	"shopping/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// user_id创建购物车
func CreateCart(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createCart := service.CartService{}

	if err := c.ShouldBind(&createCart); err == nil {
		res := createCart.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 获取user_id下的所有购物车
func GetCarts(c *gin.Context) {
	showCart := service.CartService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&showCart); err == nil {
		res := showCart.Get(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 获取cId的购物车
// func ShowCart(c *gin.Context) {
// 	showCart := service.CartService{}
// 	// _, _ := utils.ParseToken(c.GetHeader("Authorization"))
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	if err := c.ShouldBind(&showCart); err == nil {
// 		res := showCart.Show(c.Request.Context(), uint(id))
// 		c.JSON(http.StatusOK, res)
// 	} else {
// 		c.JSON(http.StatusBadRequest, ErrorResponse(err))
// 	}
// }

// 通过uId更新购物车
func UpdateCart(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createCart := service.CartService{}
	id, _ := strconv.Atoi(c.Param("id")) // 购物车的id
	if err := c.ShouldBind(&createCart); err == nil {
		res := createCart.Update(c.Request.Context(), claim.ID, uint(id))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 删除购物车
func DeleteCart(c *gin.Context) {
	deleteCarts := service.CartService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	cId, _ := strconv.Atoi(c.Param("id")) // 获得收藏夹的id

	if err := c.ShouldBind(&deleteCarts); err == nil {
		res := deleteCarts.Delete(c.Request.Context(), claim.ID, uint(cId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
