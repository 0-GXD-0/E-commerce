package v1

import (
	"E-commerce/pkg/util"
	"E-commerce/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 创建收藏夹
func CreateFavorite(c *gin.Context) {
	//获取token
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	//验证token
	claim, _ := util.ParseToken(token)
	//绑定请求数据
	createFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&createFavoriteService); err == nil {
		res := createFavoriteService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 展示收藏夹
func ListFavorite(c *gin.Context) {
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	claim, _ := util.ParseToken(token)
	listFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&listFavoriteService); err == nil {
		res := listFavoriteService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 删除收藏夹
func DeleteFavorite(c *gin.Context) {
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	claim, _ := util.ParseToken(token)
	deleteFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&deleteFavoriteService); err == nil {
		res := deleteFavoriteService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
