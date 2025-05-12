package v1

import (
	"net/http"
	"shopping/pkg/utils"
	"shopping/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// user_id创建地址
func CreateAddress(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createAddress := service.AddressService{}

	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 获取user_id下的所有地址
func GetAddresses(c *gin.Context) {
	showAddress := service.AddressService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&showAddress); err == nil {
		res := showAddress.Get(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 获取aId的地址
func ShowAddress(c *gin.Context) {
	showAddress := service.AddressService{}
	// _, _ := utils.ParseToken(c.GetHeader("Authorization"))
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&showAddress); err == nil {
		res := showAddress.Show(c.Request.Context(), uint(id))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 通过uId更新地址
func UpdateAddress(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createAddress := service.AddressService{}
	id, _ := strconv.Atoi(c.Param("id")) // 购物车的id
	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.Update(c.Request.Context(), claim.ID, uint(id))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 删除地址
func DeleteAddress(c *gin.Context) {
	deleteAddress := service.AddressService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	aId, _ := strconv.Atoi(c.Param("id")) // 获得收藏夹的id

	if err := c.ShouldBind(&deleteAddress); err == nil {
		res := deleteAddress.Delete(c.Request.Context(), claim.ID, uint(aId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
