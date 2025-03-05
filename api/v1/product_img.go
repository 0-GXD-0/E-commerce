package v1

import (
	"E-commerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductImg(c *gin.Context) {
	var listProductImg service.ListProductImg
	if err := c.ShouldBind(&listProductImg); err == nil {
		res := listProductImg.List(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
