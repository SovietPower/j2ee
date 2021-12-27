package api

import (
	"j2ee/service"
	"j2ee/util/logging"

	"github.com/gin-gonic/gin"
)

// AdminUserGet 获取所有用户信息
func AdminUserGet(c *gin.Context) {
	var service service.AdminUserGetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminUserGet()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminUserStatus 设置用户Status
func AdminUserStatus(c *gin.Context) {
	var service service.AdminUserStatusService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminUserStatus()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}
