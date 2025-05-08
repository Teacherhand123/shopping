package v1

import (
	"net/http"
	"shopping/pkg/utils"
	"shopping/service"

	"github.com/gin-gonic/gin"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createProduct := service.ProductService{}

	if err := c.ShouldBind(&createProduct); err == nil {
		res := createProduct.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
