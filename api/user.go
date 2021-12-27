package api

import (
	"j2ee/serializer"
	"j2ee/service"
	"j2ee/util/logging"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	// session := sessions.Default(c)
	// userID := session.Get("userID")
	var service service.UserRegisterService // 如果是map，需要先初始化，否则无所谓？
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserRegister()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserLogin()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}

// UserUpdate 用户修改信息
// func UserUpdate(c *gin.Context) {
// 	var service service.UserUpdateStruct
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Update()
// 		c.JSON(200, res)

// 	} else {
// 		logging.Info(err)
// 		c.JSON(200, ErrorResponse(err))
// 	}
// }

// Ping 任意发送一条消息，用jwt验证Token
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "ok",
	})
}
