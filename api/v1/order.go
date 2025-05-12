package v1

import (
	"net/http"
	"shopping/pkg/utils"
	"shopping/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// user_id创建订单
func CreateOrder(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createOrder := service.OrderService{}

	if err := c.ShouldBind(&createOrder); err == nil {
		res := createOrder.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 获取user_id下的所有订单
func GetOrders(c *gin.Context) {
	showOrder := service.OrderService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&showOrder); err == nil {
		res := showOrder.Get(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 获取cId的订单
// func ShowOrder(c *gin.Context) {
// 	showOrder := service.OrderService{}
// 	// _, _ := utils.ParseToken(c.GetHeader("Authorization"))
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	if err := c.ShouldBind(&showOrder); err == nil {
// 		res := showOrder.Show(c.Request.Context(), uint(id))
// 		c.JSON(http.StatusOK, res)
// 	} else {
// 		c.JSON(http.StatusBadRequest, ErrorResponse(err))
// 	}
// }

// 通过uId更新订单
func UpdateOrder(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createOrder := service.OrderService{}
	id, _ := strconv.Atoi(c.Param("id")) // 订单的id
	if err := c.ShouldBind(&createOrder); err == nil {
		res := createOrder.Update(c.Request.Context(), claim.ID, uint(id))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 删除订单
func DeleteOrder(c *gin.Context) {
	deleteOrders := service.OrderService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	cId, _ := strconv.Atoi(c.Param("id")) // 获得收藏夹的id

	if err := c.ShouldBind(&deleteOrders); err == nil {
		res := deleteOrders.Delete(c.Request.Context(), claim.ID, uint(cId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
