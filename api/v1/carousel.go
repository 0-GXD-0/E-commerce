package v1

import (
	"E-commerce/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService

	//绑定请求数据
	if err := c.ShouldBind(&listCarousel); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		res := listCarousel.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
