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
	files := form.File["file"] // 获取多张商品图片 再请求中的key是file
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createProduct := service.ProductService{}

	if err := c.ShouldBind(&createProduct); err == nil {
		res := createProduct.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ListProduct(c *gin.Context) {
	listProduct := service.ProductService{}

	if err := c.ShouldBind(&listProduct); err == nil {
		res := listProduct.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func SearchProduct(c *gin.Context) {
	searchProduct := service.ProductService{}

	if err := c.ShouldBind(&searchProduct); err == nil {
		res := searchProduct.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ShowProduct(c *gin.Context) {
	showProduct := service.ProductService{}

	if err := c.ShouldBind(&showProduct); err == nil {
		res := showProduct.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
