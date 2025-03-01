package v1

import (
	"E-commerce/pkg/util"
	"E-commerce/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	log.Printf("开始解析token")
	//获取并解析token
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	claims, err := util.ParseToken(token)
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token验证失败",
		})
		return
	}
	//绑定请求数据
	if err := c.ShouldBind(&userUpdate); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}
}
