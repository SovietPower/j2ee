package api

import (
	"j2ee/service"
	"j2ee/util/logging"

	"github.com/gin-gonic/gin"
)

// TruckGet 获取所有货车信息
func TruckGet(c *gin.Context) {
	var service service.TruckGetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TruckGet()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}

// TruckGetByUser 用户可调用的获取货车信息。仅返回货车数量
func TruckGetByUser(c *gin.Context) {
	var service service.TruckGetByUserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TruckGetByUser()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}

// TruckCreate 添加货车
func TruckCreate(c *gin.Context) {
	var service service.TruckCreateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TruckCreate()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}

// TruckAllocate 分配货车
func TruckAllocate(c *gin.Context) {
	var service service.TruckAllocateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TruckAllocate()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}
