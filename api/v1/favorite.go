package v1

import (
	"net/http"
	"shopping/pkg/utils"
	"shopping/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateFavorite(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createProduct := service.FavoriteService{}

	if err := c.ShouldBind(&createProduct); err == nil {
		res := createProduct.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowFavorite(c *gin.Context) {
	showFavorite := service.FavoriteService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showFavorite); err == nil {
		res := showFavorite.Show(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func DeleteFavorite(c *gin.Context) {
	deleteFavorite := service.FavoriteService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	fId, _ := strconv.Atoi(c.Param("id")) // 获得收藏夹的id

	if err := c.ShouldBind(&deleteFavorite); err == nil {
		res := deleteFavorite.Delete(c.Request.Context(), claim.ID, uint(fId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
