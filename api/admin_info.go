package api

import (
	"j2ee/service"
	"j2ee/util/logging"

	"github.com/gin-gonic/gin"
)

// AdminInfoGet 获取管理员其它信息
func AdminInfoGet(c *gin.Context) {
	var service service.AdminInfoGetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminInfoGet()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}
