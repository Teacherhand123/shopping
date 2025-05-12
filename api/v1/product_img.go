package v1

import (
	"net/http"
	"shopping/service"

	"github.com/gin-gonic/gin"
)

func ListProductImg(c *gin.Context) {
	var ListProductImg service.ListProductImg
	if err := c.ShouldBind(&ListProductImg); err == nil {
		res := ListProductImg.List(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))

	}
}
