package api

import (
	"j2ee/service"
	"j2ee/util/logging"

	"github.com/gin-gonic/gin"
)

// OrderCreate 新建订单
func OrderCreate(c *gin.Context) {
	service := service.OrderCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.OrderCreate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// OrderGet 获取订单
func OrderGet(c *gin.Context) {
	service := service.OrderGetService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.OrderGet()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// OrderDetail 获取订单详情
func OrderDetail(c *gin.Context) {
	service := service.OrderDetailService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.OrderDetail(c.Param("order_id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// OrderRate 评价订单
func OrderRate(c *gin.Context) {
	service := service.OrderRateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.OrderRate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
