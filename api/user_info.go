package api

import (
	"j2ee/service"
	"j2ee/util/logging"

	"github.com/gin-gonic/gin"
)

// UserInfoGet 获取用户其它信息
func UserInfoGet(c *gin.Context) {
	var service service.UserInfoGetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserInfoGet()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}
