package api

import (
	"j2ee/service"
	"j2ee/util/logging"

	"github.com/gin-gonic/gin"
)

// AdminOrderGet 获取订单
func AdminOrderGet(c *gin.Context) {
	service := service.AdminOrderGetService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminOrderGet()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// AdminOrderDetail 获取订单详情
func AdminOrderDetail(c *gin.Context) {
	service := service.AdminOrderDetailService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminOrderDetail(c.Param("order_id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
