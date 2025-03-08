package v1

import (
	"E-commerce/pkg/util"
	"E-commerce/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	//获取token
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	//验证token
	claim, _ := util.ParseToken(token)
	//绑定请求数据
	createOrderService := service.OrderService{}
	if err := c.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func GetOrder(c *gin.Context) {
	//获取token
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	//验证token
	claim, _ := util.ParseToken(token)
	//绑定请求数据
	getOrderService := service.OrderService{}
	if err := c.ShouldBind(&getOrderService); err == nil {
		res := getOrderService.Get(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ShowOrder(c *gin.Context) {
	//获取token
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	//验证token
	claim, _ := util.ParseToken(token)
	//绑定请求数据
	showOrderService := service.OrderService{}
	if err := c.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteOrder(c *gin.Context) {
	//获取token
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	//验证token
	claim, _ := util.ParseToken(token)
	//绑定请求数据
	deleteOrderService := service.OrderService{}
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func PayOrder(c *gin.Context) {
	//获取token
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	//验证token
	claim, _ := util.ParseToken(token)
	//绑定请求数据
	payOrderService := service.OrderService{}
	if err := c.ShouldBind(&payOrderService); err == nil {
		res := payOrderService.Pay(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
