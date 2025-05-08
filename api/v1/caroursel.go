package v1

import (
	"net/http"
	"shopping/service"

	"github.com/gin-gonic/gin"
)

func ListCarousel(c *gin.Context) {
	var ListCarousel service.CarouselService
	if err := c.ShouldBind(&ListCarousel); err == nil {
		res := ListCarousel.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))

	}
}
