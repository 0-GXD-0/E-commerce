package v1

import (
	"E-commerce/pkg/util"
	"E-commerce/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 注册
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

// 登录
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

// 更新信息
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

// 更新头像
func UpLoadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
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
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// 发送邮件
func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
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
	if err := c.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// 验证邮箱
func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService
	//获取token
	token := c.GetHeader("Authorization")
	// 去掉 "Bearer " 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	// //验证token
	// claims, err := util.ParseToken(token)
	// if err != nil {
	// 	log.Printf("Error parsing token: %v", err)
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Token验证失败",
	// 	})
	// 	return
	// }

	if err := c.ShouldBind(&validEmail); err == nil {
		res := validEmail.Valid(c.Request.Context(), token)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// 显示金额
func ShowMoney(c *gin.Context) {
	var showMoney service.ShowMoneyService
	log.Printf("调用Show前密钥是: %s", showMoney.Key)

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

	if err := c.ShouldBind(&showMoney); err == nil {
		res := showMoney.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
