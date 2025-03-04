package v1

import (
	"E-commerce/pkg/util"
	"E-commerce/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	claim, _ := util.ParseToken(token)
	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
